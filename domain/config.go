package domain

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "evolve"
)

func Init() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	//dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	errs := db.AutoMigrate(Book{})
	if errs != nil {
		return nil
	}

	a := [20]Book{}
	err3 := faker.FakeData(&a)
	if err3 != nil {
		fmt.Println(err3)
	}
	for _, v := range a {
		db.Create(&v)
	}
	fmt.Printf("%+v", a)

	return db
}
