package postgres

import (
	"errors"
	"fmt"
	"gqlgen-todos/graph/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBCon *gorm.DB
)

func InitDB() (*gorm.DB, error) {
	var err error
	url := os.Getenv("POSTGRES_URL")
	if url == "" {
		return nil, errors.New("could not get the url")

	}
	dns := url
	//connect to postgres database
	DBCon, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	//where myhost is port is the port postgres is running on
	//user is your postgres use name
	//password is your postgres password
	if err != nil {
		return nil, fmt.Errorf("failed to connect database %v", err)

	}
	user := &model.User{ID: "1", Name: "timothy"}
	todo := &model.Todo{ID: "1", Text: "this is the first todo", Done: false, User: user}
	DBCon.Create(user)
	DBCon.Create(todo)

	log.Println("database connected succesfully")
	autoMigration()

	return DBCon, nil

}

// Close DB
func Close() {
	dbSQL, err := DBCon.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}

func autoMigration() {
	log.Println("automigration")
	DBCon.AutoMigrate(&model.User{}, &model.Animal{}, &model.Todo{})

}
