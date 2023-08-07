package repository

import (
	"context"
	"github.com/TechGG1/Library/internal/model"
)

func (r *PgRepo) CreateReader(ctx context.Context, reader *model.Reader) (int, error) {
	var readerId int

	row := r.db.QueryRowContext(ctx,
		`INSERT INTO "readers" (name, surname, date_of_birth, address, email)
 			VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		reader.Name, reader.Surname, reader.DateOfBirth, reader.Address, reader.Email)
	if err := row.Scan(&readerId); err != nil {
		return 0, err
	}

	return readerId, nil
}

func (r *PgRepo) ReadersWithPage(ctx context.Context, limit, page int) ([]model.Reader, int, error) {
	var readers []model.Reader
	rows, err := r.db.QueryContext(ctx, `SELECT id, name, surname, date_of_birth, address, email 
								FROM readers ORDER BY name LIMIT $1 OFFSET $2`, limit, (page-1)*limit)
	if err != nil {
		return nil, -1, err
	}
	for rows.Next() {
		var reader model.Reader
		if err := rows.Scan(&reader.Id, &reader.Name, &reader.Surname,
			&reader.DateOfBirth, &reader.Address, &reader.Email); err != nil {
			return nil, -1, err
		}
		readers = append(readers, reader)
	}
	var amount int
	row := r.db.QueryRowContext(ctx, `SELECT COUNT(id) FROM readers`)
	if err := row.Scan(&amount); err != nil {
		return nil, 0, err
	}
	return readers, (amount / limit) + 1, nil
}

func (r *PgRepo) UpdateReader(ctx context.Context, reader *model.Reader) (int, error) {
	var readerId int
	row := r.db.QueryRowContext(ctx,
		`UPDATE "readers" SET name=$1, surname=$2, address=$3
               WHERE id = $4 RETURNING id`,
		reader.Name, reader.Surname, reader.Address, reader.Id)
	if err := row.Scan(&readerId); err != nil {
		return -1, err
	}
	return readerId, nil
}
