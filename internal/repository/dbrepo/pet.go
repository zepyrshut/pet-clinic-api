package dbrepo

import (
	"context"
	"time"

	"github.com/zepyrshut/go-web-starter-gin/internal/models"
)

func (m *postgreDBRepo) OnePet(id int) (models.Pet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT 
				id, pet_name, pet_type, pet_race, birth_date, created_at, updated_at 
			  FROM 
			  	pet
			  WHERE
			  	id = ($1)
	`

	row := m.DB.QueryRowContext(ctx, query, id)

	var pet models.Pet
	err := row.Scan(
		&pet.ID,
		&pet.PetName,
		&pet.PetType,
		&pet.PetRace,
		&pet.BirthDate,
		&pet.CreatedAt,
		&pet.UpdatedAt,
	)

	if err != nil {
		return pet, err
	}

	return pet, nil

}
