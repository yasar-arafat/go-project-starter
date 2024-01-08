package api

import (
	"github.com/yasar-arafat/go-project-starter/model"
)

type userResponse struct {
	User struct {
		Username string  `json:"username"`
		Email    string  `json:"email"`
		Bio      *string `json:"bio"`
		Image    *string `json:"image"`
	} `json:"user"`
}

type allUsersResponse struct {
	Users []User `json:"user"`
}

type User struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      *string `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

func newUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Bio = u.Bio
	r.User.Image = u.Image
	return r
}

func getAllUsersResponse(u []model.User) allUsersResponse {
	allUsersResponse := allUsersResponse{}
	for _, user := range u {
		r := User{}
		r.Username = user.Username
		r.Email = user.Email
		r.Bio = user.Bio
		r.Image = user.Image
		allUsersResponse.Users = append(allUsersResponse.Users, r)
	}
	return allUsersResponse
}
