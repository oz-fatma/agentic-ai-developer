package main

import (
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day43/db"
)

func main() {
	fmt.Println("Day 43: transactions")

	conn, err := db.OpenMemory()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := db.InitAccounts(conn); err != nil {
		log.Fatal(err)
	}

	if err := db.SeedAccounts(conn); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Before transfer:")
	printBalances(conn)

	if err := transfer(conn, 1, 2, 150); err != nil {
		log.Fatal("transfer failed:", err)
	}

	fmt.Println("\nAfter transfer:")
	printBalances(conn)

	if err := transfer(conn, 2, 1, 9999); err != nil {
		fmt.Println("\nExpected failure:", err)
	}

	fmt.Println("\nBalances unchanged after failed transfer:")
	printBalances(conn)
}

func transfer(dbConn *db.DB, from, to int64, amount int) error {
	tx, err := dbConn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var balance int
	if err := tx.QueryRow(`SELECT balance FROM accounts WHERE id = ?`, from).Scan(&balance); err != nil {
		return err
	}
	if balance < amount {
		return fmt.Errorf("insufficient funds: have %d need %d", balance, amount)
	}

	if _, err := tx.Exec(`UPDATE accounts SET balance = balance - ? WHERE id = ?`, amount, from); err != nil {
		return err
	}
	if _, err := tx.Exec(`UPDATE accounts SET balance = balance + ? WHERE id = ?`, amount, to); err != nil {
		return err
	}
	return tx.Commit()
}

func printBalances(dbConn *db.DB) {
	rows, err := dbConn.Query(`SELECT id, owner, balance FROM accounts ORDER BY id`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		var owner string
		var balance int
		if err := rows.Scan(&id, &owner, &balance); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  account %d (%s): $%d\n", id, owner, balance)
	}
}
