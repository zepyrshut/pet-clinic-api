package dbrepo

import (
	"context"
	"time"
)

func (m *postgreDBRepo) BindPetWithOwner(pet int, person int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO pet_owner
				(pet, person, created_at, updated_at)
			  VALUES
			    ($1, $2, $3, $4)
	`

	_, err := m.DB.ExecContext(ctx, query,
		pet,
		person,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
