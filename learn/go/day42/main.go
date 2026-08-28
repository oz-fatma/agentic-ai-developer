package main

import (
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day42/db"
)

func main() {
	fmt.Println("Day 42: prepared statements")

	conn, err := db.OpenMemory()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := db.InitProducts(conn); err != nil {
		log.Fatal(err)
	}

	stmt, err := conn.Prepare(`INSERT INTO products (name, price) VALUES (?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	items := [][2]any{
		{"keyboard", 79.99},
		{"mouse", 29.99},
		{"monitor", 249.00},
	}
	for _, item := range items {
		if _, err := stmt.Exec(item[0], item[1]); err != nil {
			log.Fatal(err)
		}
	}

	rows, err := conn.Query(`SELECT name, price FROM products ORDER BY price DESC`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Products by price:")
	for rows.Next() {
		var name string
		var price float64
		if err := rows.Scan(&name, &price); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  %s -> $%.2f\n", name, price)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
