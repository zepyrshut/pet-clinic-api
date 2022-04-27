package dbrepo

import (
	"context"
	"time"

	"github.com/zepyrshut/pet-clinic/internal/models"
)

func (m *postgreDBRepo) NewPerson(person models.Person) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `INSERT INTO person
				(first_name, last_name, phone, email, created_at, updated_at)
			  VALUES
			    ($1, $2, $3, $4, $5, $6)
	`

	_, err := m.DB.ExecContext(ctx, query,
		person.FirstName,
		person.LastName,
		person.Phone,
		person.Email,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *postgreDBRepo) OnePerson(id int) (models.PersonDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT 
				id, first_name, last_name, phone, email, created_at, updated_at 
			  FROM 
			  	person
			  WHERE
			  	id = ($1)
	`

	row := m.DB.QueryRowContext(ctx, query, id)

	var person models.PersonDTO
	err := row.Scan(
		&person.ID,
		&person.FirstName,
		&person.LastName,
		&person.Phone,
		&person.Email,
		&person.CreatedAt,
		&person.UpdatedAt,
	)

	if err != nil {
		return models.PersonDTO{}, err
	}

	m.appendPet(&person)

	return person, nil
}

func (m *postgreDBRepo) appendPet(id *models.PersonDTO) (*models.PersonDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT 
				pet.id, pet.pet_name, pet.pet_type, pet.pet_race, pet.birth_date
			  FROM 
			  	pet
			  JOIN 
			  	pet_owner 
			  ON 
			  	pet.id = pet_owner.pet
			  JOIN 
			  	person 
			  ON 
			  	pet_owner.person = person.id
			  WHERE 
			    person.id = ($1)
	`
	rows, _ := m.DB.QueryContext(ctx, query, id.ID)
	defer rows.Close()

	petsAppend := make(map[uint]models.Pet)
	for rows.Next() {
		var pet models.Pet
		err := rows.Scan(
			&pet.ID,
			&pet.PetName,
			&pet.PetType,
			&pet.PetRace,
			&pet.BirthDate,
		)
		if err != nil {
			return nil, err
		}
		petsAppend[pet.ID] = pet

	}

	m.App.InfoLog.Println(petsAppend)

	id.Pets = petsAppend
	return id, nil

}
