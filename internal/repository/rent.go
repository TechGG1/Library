package repository

import (
	"context"
	"github.com/TechGG1/Library/internal/model"
)

func (r *PgRepo) CreateRent(ctx context.Context, reader *model.Reader) (int, error) {
	var rentId int

	row := r.db.QueryRowContext(ctx,
		`INSERT INTO "rent" ()
 			VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		reader.Name, reader.Surname, reader.DateOfBirth, reader.Address, reader.Email)
	if err := row.Scan(&rentId); err != nil {
		return 0, err
	}

	return rentId, nil
}
