package repository

import (
	"context"
	"github.com/TechGG1/Library/internal/model"
)

func (r *PgRepo) CreateRent(ctx context.Context, rent *model.Rent) (int, error) {
	var rentId int

	row := r.db.QueryRowContext(ctx,
		`INSERT INTO "rent" (book_id, reader_id, first_date, last_date)
 			VALUES ($1, $2, $3, $4) RETURNING id`,
		rent.BookId, rent.ReaderId, rent.FirstDate, rent.LastDate)
	if err := row.Scan(&rentId); err != nil {
		return 0, err
	}

	rowNum := r.db.QueryRowContext(ctx,
		`update books set num_of_copies = num_of_copies - 1 where id = $1 returning id`,
		rent.BookId, rent.ReaderId, rent.FirstDate, rent.LastDate)
	if err := rowNum.Scan(&rentId); err != nil {
		return 0, err
	}

	return rentId, nil
}

func (r *PgRepo) RentById(ctx context.Context, id int) (*model.Rent, error) {
	var rentFromDB model.Rent

	row := r.db.QueryRowContext(ctx,
		`select book_id, reader_id, first_date, last_date, fine
				from rent where id=$1`, id)
	if err := row.Scan(&rentFromDB.BookId, &rentFromDB.ReaderId, &rentFromDB.FirstDate, &rentFromDB.LastDate,
		&rentFromDB.Fine); err != nil {
		return nil, err
	}
	return &rentFromDB, nil
}

func (r *PgRepo) RentsWithPage(ctx context.Context, limit, page, readerId int) ([]model.Rent, int, error) {
	var rents []model.Rent
	rows, err := r.db.QueryContext(ctx, `SELECT id, book_id, reader_id, first_date, last_date, fine 
								FROM rent where reader_id=$1 ORDER BY id LIMIT $2 OFFSET $3`, readerId, limit, (page-1)*limit)
	if err != nil {
		return nil, -1, err
	}

	for rows.Next() {
		var rent model.Rent
		if err := rows.Scan(&rent.RentId, &rent.BookId, &rent.ReaderId, &rent.FirstDate, &rent.LastDate, &rent.Fine); err != nil {
			return nil, -1, err
		}
		rents = append(rents, rent)
	}
	var amount int
	row := r.db.QueryRowContext(ctx, `SELECT COUNT(id) FROM rent`)
	if err := row.Scan(&amount); err != nil {
		return nil, 0, err
	}
	return rents, (amount / limit) + 1, nil
}

func (r *PgRepo) UpdateRent(ctx context.Context, rent *model.Rent) (int, error) {
	var complete bool
	var rentId int
	row := r.db.QueryRowContext(ctx,
		`UPDATE "rent" SET complete=$1, fine=$2
               WHERE id = $3 returning id, complete`,
		rent.Complete, rent.Fine, rent.RentId)
	if err := row.Scan(&rentId, &complete); err != nil {
		return -1, err
	}
	if complete {
		rowNum := r.db.QueryRowContext(ctx,
			`update books set num_of_copies = num_of_copies - 1 where id = $1 returning id`,
			rent.BookId, rent.ReaderId, rent.FirstDate, rent.LastDate)
		if err := rowNum.Scan(&rentId); err != nil {
			return 0, err
		}

	}
	return rentId, nil
}
