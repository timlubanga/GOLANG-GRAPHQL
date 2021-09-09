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
	log.Println("database connected succesfully")
	autoMigration()
	// usid := fmt.Sprintf("%v", uuid.NewV4())
	// user := model.User{Name: "timothy", ID: usid}
	// DBCon.Create(&user)
	// todo := model.Todo{Text: "finish today's tickets", User: &user}
	// DBCon.Create(&todo)
	// time.Sleep(time.Millisecond * 5)
	// var todofirst model.Todo
	// DBCon.First(&todofirst)
	// fmt.Println("this is the first", todofirst)
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
	DBCon.AutoMigrate(&model.User{}, &model.Todo{})

}
