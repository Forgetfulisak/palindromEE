package storage

import (
	"palindromee/pkg/userservice"
)

type Storage struct {
}

type dbUser struct {
}

func New() *Storage {
	return nil
}

func (s *Storage) GetUser(userID string) (userservice.User, error) {
	return userservice.User{
		UserID:    userID,
		FirstName: "swag",
		LastName:  "geag",
	}, nil
}

func (s *Storage) InsertUser(user userservice.User) (userservice.User, error) {
	return userservice.User{
		UserID:    "potet",
		FirstName: "swag",
		LastName:  "geag",
	}, nil
}

func (s *Storage) UpdateUser(user userservice.User) (userservice.User, error) {
	return userservice.User{
		UserID:    "potet",
		FirstName: "swag",
		LastName:  "geag",
	}, nil
}

func (s *Storage) DeleteUser(userID string) (userservice.User, error) {
	return userservice.User{
		UserID:    "potet",
		FirstName: "swag",
		LastName:  "geag",
	}, nil
}
