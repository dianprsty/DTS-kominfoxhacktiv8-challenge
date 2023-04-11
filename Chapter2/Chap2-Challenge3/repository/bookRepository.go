package repository

import (
	"Chap2-Challenge3/config"
	"Chap2-Challenge3/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetAll() ([]models.Book, error) {
	var Books []models.Book

	db, err := config.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	queryText := "SELECT * FROM book"
	rowQuery, err := db.Query(queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var book models.Book
		if err = rowQuery.Scan(
			&book.Id, &book.Title, &book.Author, &book.Desc,
		); err != nil {
			return nil, err
		}

		Books = append(Books, book)

	}
	return Books, nil

}

func GetById(id string) (interface{}, error) {

	db, err := config.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var book models.Book
	queryText := fmt.Sprintf("SELECT * FROM book WHERE id=%v", id)
	err = db.QueryRow(queryText).Scan(&book.Id,
		&book.Title,
		&book.Author,
		&book.Desc)
	if err != nil {
		return nil, err
	}

	if book.Id == "" && book.Title == "" {
		return nil, errors.New("id tidak ada")
	}
	return book, nil

}

func InsertBook(book models.Book) (interface{}, error) {
	var b = models.Book{}
	db, err := config.ConnectDB()
	if err != nil {
		return nil, errors.New("can't connect db")
	}
	defer db.Close()

	fmt.Println(book.Desc)
	queryText := `
	INSERT INTO book (title, author, "desc")
	VALUES ($1, $2, $3)
	Returning *
	`
	err = db.QueryRow(queryText, book.Title, book.Author, book.Desc).Scan(&b.Id, &b.Title, &b.Author, &b.Desc)

	if err != nil {
		return nil, err
	}

	return b, nil

}

func Update(book models.Book, id string) error {
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	queryText := `UPDATE book SET title=$1, author=$2, "desc"=$3 WHERE id=$4`
	s, err := db.Exec(queryText, book.Title, book.Author, book.Desc, id)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		return err
	}

	return nil
}

func Delete(id string) error {
	db, err := config.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	queryText := fmt.Sprintf("DELETE FROM %v WHERE id=%v", "book", id)

	s, err := db.Exec(queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
