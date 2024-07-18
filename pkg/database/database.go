package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init() {
    connectDatabase()
    createTables()
}

func connectDatabase() {
    var err error
    DB, err = sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable password=123456 host=localhost")
    if err != nil {
        log.Fatalln(err)
    }
}

func createTables() {
    schema := `
    CREATE TABLE IF NOT EXISTS todos (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        completed BOOLEAN NOT NULL
    );`

    DB.MustExec(schema)
	fmt.Println("Create table todos successfully")

    insertInitialData()
}

func insertInitialData() {
    var count int
    err := DB.Get(&count, "SELECT COUNT(*) FROM todos")
    if err != nil {
        log.Fatalln(err)
    }

    if count == 0 {
        todos := []struct {
            Title     string
            Completed bool
        }{
            {"Create a new API", true},
            {"Update the API", false},
            {"Delete the API", false},
			{"Edit the API", true},
        }

        for _, todo := range todos {
            _, err := DB.Exec("INSERT INTO todos (title, completed) VALUES ($1, $2)", todo.Title, todo.Completed)
            if err != nil {
                log.Fatalln(err)
            }
        }
    }
}
