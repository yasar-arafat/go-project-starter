package server

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/natefinch/lumberjack"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/yasar-arafat/go-project-starter/api"
	"github.com/yasar-arafat/go-project-starter/config"
	"github.com/yasar-arafat/go-project-starter/constants"
	"github.com/yasar-arafat/go-project-starter/dal"
	"github.com/yasar-arafat/go-project-starter/utils"
)

// Run sets up and starts the server using the configuration defined in the Setup function.
// It also handles graceful shutdown of the server when an interrupt signal is received.
func Run() {
	// Set up the Echo server with configuration
	e := Setup()

	// Wait for an interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals, as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Create a context with a timeout of 10 seconds for the server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := e.Shutdown(ctx); err != nil {
		// Log a fatal error if the shutdown process fails
		e.Logger.Fatal(err)
	}
}

// Setup initializes and configures the Echo web framework, creating a fully
// configured instance ready for use. This includes setting up middleware,
// loading configuration settings, creating necessary directories, and
// initializing the database. The function returns an instance of the Echo
// framework.
func Setup() *echo.Echo {
	// Create a new Echo instance
	e := echo.New()

	// Request ID middleware generates a unique id for a request.
	e.Use(middleware.RequestID())

	// Load configuration settings from the config package
	cfg, err := config.Read()
	if err != nil {
		// Log a fatal error and terminate the application if config loading fails
		e.Logger.Fatalf("Error while loading config : %v\n", err)
	}

	// Parse the log level from the configuration file using the ParseLogLevel function from the utils package.
	logLevel, err := utils.ParseLogLevel(cfg.Log.Level)

	// Check if an error occurred during log level parsing.
	if err != nil {
		// Log a warning message indicating that an invalid log level was provided.
		// Since there was an error, the default log level INFO will be used.
		e.Logger.Warn("Invalid log level provided, using default INFO log level")
	}

	// Set the logging level for detailed logging
	e.Logger.SetLevel(logLevel)

	// Remove trailing slashes from the request path
	e.Pre(middleware.RemoveTrailingSlash())

	// Initialize Lumberjack logger for log rotation
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(constants.LogFolderPath, cfg.Log.FileName), // Set the log file path
		MaxSize:    cfg.Log.MaxSize,                                          // Max size in megabytes before log rotation
		MaxBackups: cfg.Log.MaxBackups,                                       // Max number of old log files to retain
		MaxAge:     cfg.Log.MaxAge,                                           // Max age of log files before deletion
		Compress:   cfg.Log.Compress,                                         // Whether to compress old log files
	}

	// Set up the Echo logger middleware with Lumberjack
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: io.MultiWriter(os.Stdout, lumberjackLogger),
	}))

	e.Logger.SetOutput(io.MultiWriter(os.Stdout, lumberjackLogger))

	// Enable Cross-Origin Resource Sharing (CORS) with specific configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Set a custom validator for request and response validation
	e.Validator = utils.NewValidator()

	// Create data and log directories if they don't exist
	if err := utils.CreateFolder(constants.DBFolderPath); err != nil {
		// Log a fatal error and terminate the application if folder creation fails
		e.Logger.Fatal("failed to create data directory")
	}
	if err := utils.CreateFolder(constants.LogFolderPath); err != nil {
		// Log a fatal error and terminate the application if folder creation fails
		e.Logger.Fatal("failed to create log directory")
	}

	// Initialize the database with the specified file name
	gormDBErr := dal.InitDB(e.Logger, filepath.Join(constants.DBFolderPath, cfg.Database.FileName), cfg.Database.MaxIdleConnections)
	if gormDBErr.Error != nil {
		// Log a fatal error and terminate the application if migration fails
		e.Logger.Fatalf("failed to run migration error: %v\n", gormDBErr.Error)
	}
	// Register HTTP request handlers defined in the application
	RegisterHandlers(e)

	// Start the server in a goroutine
	go func() {
		// Start the server on the specified port from the configuration
		if err := e.Start(":" + cfg.Server.Port); err != nil && err != http.ErrServerClosed {
			// Log a fatal error and terminate the application if server start fails
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Return the fully configured Echo instance
	return e
}

// RegisterHandlers sets up the HTTP request handlers for the specified Echo instance.
// It defines routes and associates them with corresponding handler functions.
func RegisterHandlers(e *echo.Echo) {
	// Serve Swagger UI at /swagger/* using echoSwagger middleware
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Create a versioned API group "/api"
	v1 := e.Group("/api")

	// Define routes for guest users under the "/api/users" endpoint
	guestUsers := v1.Group("/users")

	// Define a route to handle user sign-up (POST /api/users)
	guestUsers.POST("", api.SignUp)

	// Define a route to retrieve all users (GET /api/users)
	guestUsers.GET("", api.GetAllUser)
}
