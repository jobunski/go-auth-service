package service

import (
	"authe/internal/config"
	"authe/internal/database"
	"authe/internal/model"
	"authe/pkg"
	"time"
)

func CreateUser(firstname, lastname, msisdn, email string) (res model.Response, err error) {
	var user database.User = database.User{
		UUID:         pkg.GenerateUUID(),
		FirstName:    firstname,
		LastName:     lastname,
		PhoneNumber:  msisdn,
		EmailAddress: email,
		CreatedAt:    time.Now(),
	}

	dataStoreSession := config.CreateDatabaseConnection()
	dataStoreSession.Create(&user)

	if err != nil {
		return model.Response{Status: 400, Message: "User not created"}, err
	} else {
		return model.Response{
			Status:  200,
			Message: "User created successfully",
		}, nil
	}

}
