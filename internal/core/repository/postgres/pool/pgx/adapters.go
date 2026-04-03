package corepgxpool

import (
	"errors"
	"fmt"

	corepostgrespool "github.com/977ADAM/golang-todoapp-project/internal/core/repository/postgres/pool"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRows struct {
	pgx.Rows
}

type pgxRow struct {
	pgx.Row
}

func (r pgxRow) Scan(dest ...any) error {
	err := r.Row.Scan(dest...)
	if err != nil {
		return mapErrors(err)
	}
	return nil
}

type pgxCommandTag struct {
	pgconn.CommandTag
}

func mapErrors(err error) error {
	const (
		pgxViolatesForeignKeyCode = "23503"
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return corepostgrespool.ErrNoRows
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgxViolatesForeignKeyCode {
			return fmt.Errorf(
				"%v: %w",
				err,
				corepostgrespool.ErrViolatesForeignKey,
			)
		}
	}

	return fmt.Errorf(
		"%v: %w",
		err,
		corepostgrespool.ErrUnknown,
	)
}
