package repository

import (
	"database/sql"
	"github.com/hexhoc/go-mall-api/internal/entity"
	"time"
)

//BookPostgres mysql repo
type BookPostgres struct {
	db *sql.DB
}

//NewBookPostgres create new repository
func NewBookPostgres(db *sql.DB) *BookPostgres {
	return &BookPostgres{
		db: db,
	}
}

//Create a book
func (r *BookPostgres) Create(e *entity.Book) (entity.ID, error) {
	stmt, err := r.db.Prepare(`
		insert into book (id, title, author, pages, quantity, created_at) 
		values(?,?,?,?,?,?)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.Title,
		e.Author,
		e.Pages,
		e.Quantity,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

//Get a book
func (r *BookPostgres) Get(id entity.ID) (*entity.Book, error) {
	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book where id = ?`)
	if err != nil {
		return nil, err
	}
	var b entity.Book
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.Quantity, &b.CreatedAt)
	}
	return &b, nil
}

//Update a book
func (r *BookPostgres) Update(e *entity.Book) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec("update book set title = ?, author = ?, pages = ?, quantity = ?, updated_at = ? where id = ?", e.Title, e.Author, e.Pages, e.Quantity, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

//Search books
func (r *BookPostgres) Search(query string) ([]*entity.Book, error) {
	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book where title like ?`)
	if err != nil {
		return nil, err
	}
	var books []*entity.Book
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b entity.Book
		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.Quantity, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &b)
	}

	return books, nil
}

//List books
func (r *BookPostgres) List() ([]*entity.Book, error) {
	stmt, err := r.db.Prepare(`select id, title, author, pages, quantity, created_at from book`)
	if err != nil {
		return nil, err
	}
	var books []*entity.Book
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var b entity.Book
		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Pages, &b.Quantity, &b.CreatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, &b)
	}
	return books, nil
}

//Delete a book
func (r *BookPostgres) Delete(id entity.ID) error {
	_, err := r.db.Exec("delete from book where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
