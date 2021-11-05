package database

import (
	"math/big"
	"time"
)

type User struct {
	ID big.Int  `db:"id" gorm:"primaryKey"`
	UUID string `db:"uuid" gorm:"uniq"`
	FirstName string `db:"first_name"`
	LastName    string `db:"last_name"`
	PhoneNumber string `db:"phone_number" required:"true"`
	EmailAddress string `db:"email_address"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`


}
