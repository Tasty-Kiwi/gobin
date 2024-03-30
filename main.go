package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func prepareDb(db *sql.DB) {
	schema, err := os.ReadFile("./schema.sql")
	checkErr(err)

	stmt, err := db.Prepare(string(schema))
	checkErr(err)

	_, err = stmt.Exec()
	checkErr(err)
}

func main() {
	// Create tables for the database
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)
	prepareDb(db)

	http.Handle("GET /", templ.Handler(Index()))
	http.Handle("GET /new", templ.Handler(New()))
	http.HandleFunc("POST /new", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if len(r.Form.Get("contents")) > 0 {
			stmt, err := db.Prepare("INSERT INTO bins(uuid, content, creation_date) values(?,?,?)")
			checkErr(err)

			uid := uuid.New().String()
			_, err = stmt.Exec(uid, r.Form.Get("contents"), time.Now().Unix())
			checkErr(err)

			NewSuccess(uid).Render(context.Background(), w)
		} else {
			Error("Empty operation.").Render(context.Background(), w)
		}
	})
	http.HandleFunc("GET /view/{id}", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM bins WHERE uuid = ?", r.PathValue("id"))

		checkErr(err)
		var uid string
		var content string
		var created int

		if rows.Next() {
			err = rows.Scan(&uid, &content, &created)
			checkErr(err)
			View(uid, content, time.Unix(int64(created), 0).Format(time.RFC822)).Render(context.Background(), w)
		} else {
			Error("404 not found").Render(context.Background(), w)
		}
		rows.Close()
	})

	fmt.Println("Listening on :3000")
	http.ListenAndServe("127.0.0.1:3000", nil)
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
