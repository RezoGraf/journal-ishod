package main

import (
	"database/sql"
	"journal-ishod-1/handlers"

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
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL
	);
    `

	_, err := db.Exec(sql)
	// выходим, если будут ошибки с SQL запросом выше
	if err != nil {
		panic(err)
	}
}

func main() {

	//-------------SQLITE--------------------------

	// os.Remove("./foo.db")

	db := initDB("./foo.db")

	// _, err = db.Exec(sqlStmt)
	// if err != nil {
	// 	log.Printf("%q: %s\n", err, sqlStmt)
	// 	return
	// }

	// Create a new instance of Echo
	e := echo.New()

	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// Start as a web server
	e.Logger.Fatal(e.Start(":8080"))
}
