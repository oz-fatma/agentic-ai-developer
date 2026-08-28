package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day44/db"
)

func main() {
	fmt.Println("Day 44: NULL handling and Scan types")

	conn, err := db.OpenMemory()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := db.InitEmployees(conn); err != nil {
		log.Fatal(err)
	}

	if err := db.SeedEmployees(conn); err != nil {
		log.Fatal(err)
	}

	rows, err := conn.Query(`SELECT name, department, salary FROM employees ORDER BY id`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var dept sql.NullString
		var salary sql.NullFloat64
		if err := rows.Scan(&name, &dept, &salary); err != nil {
			log.Fatal(err)
		}

		deptLabel := "none"
		if dept.Valid {
			deptLabel = dept.String
		}
		salaryLabel := "confidential"
		if salary.Valid {
			salaryLabel = fmt.Sprintf("$%.0f", salary.Float64)
		}
		fmt.Printf("  %s | dept=%s | salary=%s\n", name, deptLabel, salaryLabel)
	}
}
