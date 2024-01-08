package dal

import (
	"fmt"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/yasar-arafat/go-project-starter/model"
)

var DB *gorm.DB

// Logger is a simple Gorm logger that logs to Echo's logger
type Logger struct {
	logger echo.Logger
}

// Print formats log messages and sends them to Echo's logger
func (l Logger) Print(v ...interface{}) {
	l.logger.Info(fmt.Sprint(v...))
}

func InitDB(logger echo.Logger, dbPath string, maxIdleConn int) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(maxIdleConn)
	db.LogMode(true)
	db.SetLogger(Logger{logger: logger})
	DB = db
	return AutoMigrate()
}

func TestDB(dbPath string) {
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	DB = db
}

func DropTestDB(dbPath string) error {
	if err := os.Remove(dbPath); err != nil {
		return err
	}
	return nil
}

func AutoMigrate() *gorm.DB {
	return DB.AutoMigrate(
		&model.User{},
	)

}
