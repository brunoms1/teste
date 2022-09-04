package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres dbname=alura_loja password=b10c02 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// insert := `insert into produtos("nome", "descricao", "preco", "quantidade") values('Jeans', 'Jeans', 21, 1)`
	// re, e := db.Exec(insert)

	// if e != nil {
	// 	panic(e.Error())
	// }

	// fmt.Println(re)

	pa := ""
	row2 := db.QueryRow("Select nome FROM produtos WHERE id=1").Scan(&pa)
	fmt.Println(pa)
	if row2 != nil {
		panic(row2.Error())
	}
}
