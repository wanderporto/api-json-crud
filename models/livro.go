package models

type Book struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

func NewBook(book Book) (bool, error) {
	con := Connect()

	sql := "INSERT INTO livros(titulo,autor) VALUES (?, ?)"

	stmt, err := con.Prepare(sql)

	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(book.Titulo, book.Autor)

	if err != nil {
		return false, err
	}

	defer stmt.Close()
	defer con.Close()

	return true, nil
}

func GetBooks() ([]Book, error) {

	con := Connect()

	sql := "SELECT id,titulo,autor from livros"

	rs, err := con.Query(sql)

	if err != nil {
		return nil, err
	}

	var books []Book

	for rs.Next() {

		var book Book

		err := rs.Scan(&book.Id, &book.Titulo, &book.Autor)

		if err != nil {
			return nil, err
		}

		books = append(books, book)

	}

	defer rs.Close()
	defer con.Close()

	return books, nil

}

func GetBook(id int) (Book, error) {

	con := Connect()

	sql := "SELECT id,titulo,autor FROM livros where id = ?"

	rs := con.QueryRow(sql, id)

	var book Book

	errorScan := rs.Scan(&book.Id, &book.Titulo, &book.Autor)

	if errorScan != nil {
		return book, errorScan
	}

	defer con.Close()

	return book, nil

}

func UpdateBook(book Book) (int64, error) {
	con := Connect()

	sql := `UPDATE livros 
			   SET titulo = ?,
				    autor = ?
			 WHERE id = ?`

	stmt, err := con.Prepare(sql)

	if err != nil {
		return 0, err
	}

	rs, errRs := stmt.Exec(book.Titulo, book.Autor, book.Id)

	if errRs != nil {
		return 0, errRs
	}

	rows, errRows := rs.RowsAffected()

	if errRows != nil {
		return 0, errRows
	}

	defer stmt.Close()
	defer con.Close()

	return rows, nil

}

func DeleteBook(id int) (int64, error) {
	con := Connect()

	sql := `DELETE FROM livros WHERE id = ?`

	stmt, err := con.Prepare(sql)

	if err != nil {
		return 0, err
	}

	rs, err := stmt.Exec(id)

	if err != nil {
		return 0, err
	}

	rows, err := rs.RowsAffected()

	defer stmt.Close()
	defer con.Close()

	return rows, nil

}
