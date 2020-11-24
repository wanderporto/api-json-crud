package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//abrir a conexao
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/api-json-go")

	if err != nil {
		panic(err)
	}

	sqlResult, err := db.Exec(
		`CREATE TABLE users(
				id INT NOT NULL AUTO_INCREMENT,
				name VARCHAR(50) NOT NULL,
				email VARCHAR(100) NOT NULL,
				password VARCHAR(100) NOT NULL,
				createdat VARCHAR(100) ,
				PRIMARY KEY(id)
				)`)

	if err != nil {
		log.Fatal(err.Error())
	}

	//_, errorInsert := db.Exec(`INSERT INTO livros(autor,titulo) values('Jose de alencar','O Guarani'),('insert mult','multinserts')`)

	/*
		if (errorInsert) != nil {
			log.Fatal(`Erro ao inserir registro ` + errorInsert.Error())
		}*/

	fmt.Println(sqlResult.RowsAffected())

}
