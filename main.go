package main

import (
	"fmt"
	"os"

	"github.com/codepnw/hexagonal/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := sqlx.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		panic(err)
	}

	custRepo := repository.NewCustomerRepositoryDB(db)

	customers, err := custRepo.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(customers)
}
