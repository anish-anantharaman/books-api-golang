package sql_queries

/* this file will contain all the SQL queries that are used by different APIs */
const (
	CreateBook = `INSERT INTO books (book_name, author, category, price, isbn) VALUES (?, ?, ?, ?, ?);`

	GetBookByID = `SELECT id, book_name, author, category, price, isbn FROM books WHERE id = ?;`

	GetBooks = `SELECT* FROM books;`

	UpdateBook = `UPDATE books SET book_name = ?, author = ?, category = ?, price = ?, isbn = ? WHERE id = ?;`

	DeleteBook = `DELETE FROM books WHERE id = ?;`
)