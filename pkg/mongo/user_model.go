package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"backupBro/pkg"
	"gopkg.in/mgo.v2"
)

type userModel struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Username string
	Email string
	Password string
}

func userModelIndex() mgo.Index {
	return mgo.Index{
		Key: []string{"email"},
		Unique: true,
		DropDups: true,
		Background: true,
		Sparse: true,
	}
}

func newUserModel(u *root.User) *userModel {
	return &userModel{
		Username: u.Username,
		Email: u.Email,
		Password: u.Password,
	}
}

func(u *userModel) toRootUser() *root.User {
	return &root.User{
		Id: u.Id.Hex(),
		Username: u.Username,
		Email: u.Email,
		Password: u.Password,
	}
}