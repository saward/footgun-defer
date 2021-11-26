package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var loops = 200
var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")

	// Get second argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go [func|loop]")
		os.Exit(1)
	}

	mode := os.Args[1]

	if err != nil {
		panic(err)
	}

	switch mode {
	case "loop":
		deferInLoop()
	case "func":
		deferInFunc()
	case "explicit_loop":
		rollbackInLoop()
	default:
		panic("unknown mode")
	}
}

func deferInLoop() {
	for i := 0; i < loops; i++ {
		var result bool
		tx, err := db.Begin()
		defer tx.Rollback()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		err = tx.QueryRow("SELECT true").Scan(&result)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		log.Printf("loop count %d.  Result: %t", i, result)
	}
}

func rollbackInLoop() {
	for i := 0; i < loops; i++ {
		var result bool
		tx, err := db.Begin()

		if err != nil {
			fmt.Println(err.Error())
			tx.Rollback()
			continue
		}
		err = tx.QueryRow("SELECT true").Scan(&result)
		if err != nil {
			fmt.Println(err.Error())
			tx.Rollback()
			continue
		}
		log.Printf("loop count %d.  Result: %t", i, result)
		tx.Rollback()
	}
}

func deferInFunc() {
	for i := 0; i < loops; i++ {
		err := deferInFuncFetch(db, i)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}
}

func deferInFuncFetch(db *sql.DB, i int) error {
	var result bool
	tx, err := db.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}
	err = tx.QueryRow("SELECT true").Scan(&result)
	if err != nil {
		return err
	}
	log.Printf("loop count %d", i)
	err = tx.Commit()

	if err != nil {
		return err
	}
	return nil
}
