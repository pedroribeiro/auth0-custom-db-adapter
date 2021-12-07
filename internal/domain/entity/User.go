package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string `gorm:"column:email;NOT NULL;index;unique" json:"email"`
	Password   string `gorm:"column:password;NOT NULL" json:"password"`
	ClientId   string `gorm:"column:client_id;NOT NULL;index" json:"client_id"`
	Tenant     string `gorm:"column:tenant;NOT NULL;" json:"tenant"`
	Connection string `gorm:"column:connection;NOT NULL;" json:"connection"`
}

// * email: the user's email
// * password: the password entered by the user, in plain text
// * tenant: the name of this Auth0 account
// * client_id: the client ID of the application where the user signed up, or
//              API key if created through the API or Auth0 dashboard
// * connection: the name of this database connection
