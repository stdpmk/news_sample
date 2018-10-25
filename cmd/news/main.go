package main

import (
	"log"

	"github.com/stdpmk/news_sample/app"
	"github.com/stdpmk/news_sample/db"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	a := app.NewApp(&db.ConnOpts{
		User:         "postgres",
		Password:     "postgres",
		Host:         "localhost",
		Port:         "5432",
		DatabaseName: "postgres",
	})

	a.Run("8081")
	a.Close()

}
