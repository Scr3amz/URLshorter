package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"
	
)

type URLs struct {
	ID       string `json:"id"`
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

var q1 string = `
	INSERT INTO urls
 	(longurl, shorturl) 
 	VALUES 
 	($1, $2)
	`

var q2 string = `
INSERT INTO urls
(longurl, shorturl)
VALUES
(longurl, shorturl)
`


func main() {
	db, err := sqlx.Connect("pgx", "host=localhost port=5433 user=postgres password=M3088 dbname=urlDB sslmode=disable")
	if err!= nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
        os.Exit(1)
    }
	defer db.Close()

	u:= URLs{LongURL: "https://habr.com/ru/companies/inDrive/articles/6900", ShortURL: ""}
	ctx := context.Background()

	// if _, err = db.Exec(q1, urls.Longurl, urls.Shorturl) ; err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to create row: %v\n", err)
    //     os.Exit(1)
	// }

	// if _,err := db.NamedExec(q2, urls) ; err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to create row: %v\n", err)
    //     os.Exit(1)
	// }

	q := `
	SELECT shorturl 
	FROM urls
 	WHERE longurl =  $1
	`
	var	url URLs
	if err := db.QueryRowxContext(ctx, q, u.LongURL).Scan(&url.ShortURL); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to find row: %v\n", err)
        os.Exit(1)
	}

	fmt.Printf("%v", url)


	fmt.Println("success")

}