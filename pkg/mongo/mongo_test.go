package mongo_test

import (
	"testing"
	"backupBro/pkg"
	"backupBro/pkg/mongo"
	"log"
	"backupBro/pkg/mock"
)

const (
	mongoUrl = "localhost:27017"
	dbName = "test_db"
	userCollectionName = "name"
)

func Test_UserService(t *testing.T) {
	t.Run("CreateUser", createUser_should_insert_user_into_mongo)
}

func createUser_should_insert_user_into_mongo(t *testing.T) {
	session, err := mongo.NewSession(mongoUrl)
	if err != nil {
		log.Fatalf("Unable to connect to Mongo: %s", err)
	}

	defer func() {
		session.DropDatabase(dbName)
		session.Close()
	}()

	mockHash := mock.Hash{}
	userService := mongo.NewUserService(session.Copy(), dbName, userCollectionName, &mockHash)
	testUsername := "kiirabrightstar"
	testEmail := "kiirabrightstar@gmail.com"
	testPassword := "k11rabr1ght5tar"
	user := root.User{
		Username: testUsername,
		Email: testEmail,
		Password: testPassword,
	}

	err = userService.CreateUser(&user)

	if err != nil {
		t.Errorf("Unable to create user: %s", err)
	}

	var results []root.User
	session.GetCollection(dbName, userCollectionName).Find(nil).All(&results)

	count := len(results)
	if count != 1 {
		t.Error("Too many results. Expected `1`, got `%i`", count)
	}
	if results[0].Email != user.Email {
		t.Errorf("Incorrect email. Expected `%s`, got `%s`", testEmail, results[0].Email)
	}


}