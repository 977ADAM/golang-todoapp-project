package userspostgresrepository

import (
	"context"
	"fmt"

	coreerrors "github.com/977ADAM/golang-todoapp-project/internal/core/errors"
)



func (r *UsersRepository) DeleteUser(
	ctx context.Context,
	id int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE from todoapp.users
	WHERE id=$1
	`



	cmdTag, err := r.pool.Exec(
		ctx,
		query,
		id,
	)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("user whth id='%d': %w", id, coreerrors.ErrNotFound)
	}
	return nil 
}