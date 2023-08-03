package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/TechGG1/Library/internal/model"
	"strings"
)

type PgRepo struct {
	db *sql.DB
}

func NewPgRepo(db *sql.DB) *PgRepo {
	return &PgRepo{db: db}
}

func (r *PgRepo) BooksWithPage(ctx context.Context, limit, page int) ([]model.Book, int, error) {
	var books []model.Book
	rows, err := r.db.QueryContext(ctx, `SELECT b.id, b.name, string_agg(g.name, ','),
       b.price_of_book, b.num_of_copies, b.cover_photo,b.price_per_day, b.reg_date
			FROM books as b
    			JOIN book_to_genre btg on b.id = btg.book_id
    			JOIN genres g on btg.genre_id = g.id
			GROUP BY b.id, b.name
			ORDER BY b.name LIMIT $1 OFFSET $2`, limit, (page-1)*limit)
	if err != nil {
		return nil, -1, err
	}
	for rows.Next() {
		var book model.Book
		var genres string
		if err := rows.Scan(&book.BookId, &book.Name, &genres, &book.PriceOfBook,
			&book.NumOfCopies, &book.CoverPhoto, &book.PricePerDay, &book.RegDate); err != nil {
			return nil, -1, err
		}
		arrGenres := strings.Split(genres, ",")
		for _, g := range arrGenres {
			book.Genre = append(book.Genre, model.Genre{Name: g})
		}
		books = append(books, book)
	}
	var amount int
	row := r.db.QueryRowContext(ctx, `SELECT COUNT(id) FROM books`)
	if err := row.Scan(&amount); err != nil {
		return nil, 0, err
	}
	return books, (amount / limit) + 1, nil
}

func (r *PgRepo) CreateBook(ctx context.Context, book *model.Book) (int, error) {
	var bookId int
	fmt.Println(book)
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	rowBook := tx.QueryRowContext(ctx,
		`INSERT INTO "books" (name, price_of_book, num_of_copies, cover_photo,price_per_day, reg_date)
 			VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		book.Name, book.PriceOfBook, book.NumOfCopies, book.CoverPhoto,
		book.PricePerDay, book.RegDate,
	)
	if err := rowBook.Scan(&bookId); err != nil {
		return 0, err
	}

	for _, g := range book.Genre {
		fmt.Println(bookId, g.Name)
		_, err := tx.ExecContext(ctx,
			`insert into "book_to_genre" (book_id, genre_id)
		SELECT $1, min(id) from genres as g where g.name = $2`,
			bookId, g.Name,
		)
		if err != nil {
			return -1, err
		}
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return bookId, nil
}
