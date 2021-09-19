package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type LOGIN struct {
	USER     string `json:"user" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

func initDB(filepath string) *sql.DB {
	//откроем файл или создадим его
	db, err := sql.Open("sqlite3", filepath)

	// проверяем ошибки и выходим при их наличии
	if err != nil {
		panic(err)
	}

	// если ошибок нет, но не можем подключиться к базе данных,
	// то так же выходим
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	create table IF NOT EXISTS foo (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		id INTEGER not null PRIMARY KEY,
		name text VARCHAR NOT NULL);
	delete from foo;
    `

	_, err := db.Exec(sql)
	// выходим, если будут ошибки с SQL запросом выше
	if err != nil {
		panic(err)
	}
}

func main() {

	//-------------SQLITE--------------------------

	os.Remove("./foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `

	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	// Create a new instance of Echo
	e := echo.New()

	e.GET("/tasks", func(c echo.Context) error { return c.JSON(200, "GET Tasks") })
	e.PUT("/tasks", func(c echo.Context) error { return c.JSON(200, "PUT Tasks") })
	e.DELETE("/tasks/:id", func(c echo.Context) error { return c.JSON(200, "DELETE Task "+c.Param("id")) })

	// Start as a web server
	e.Logger.Fatal(e.Start(":8080"))
}
