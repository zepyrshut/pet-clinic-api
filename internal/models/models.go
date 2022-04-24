package models

import (
	"database/sql"
	"time"
)

// Database connection values.
type DBModel struct {
	DB *sql.DB
}

// Wrapper for all models.
type Models struct {
	DB DBModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// ========================================================
// Entity
// ========================================================

// In the database the Table struct is called "mesa".
type Pet struct {
	ID        int       `json:"id"`
	PetName   string    `json:"pet_name"`
	PetType   string    `json:"pet_type"`
	PetRace   string    `json:"pet_race"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type People struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PetOwner struct {
	ID        int       `json:"id"`
	Pet       Pet       `json:"pet_id" db:"pet_id"`
	People    People    `json:"people_id" db:"people_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Veterinarian struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Visit struct {
	ID           int          `json:"id"`
	Date         time.Time    `json:"date"`
	Pet          Pet          `json:"pet_id" db:"pet_id"`
	Veterinarian Veterinarian `json:"veterinarian_id" db:"veterinarian_id"`
	Reason       string       `json:"reason"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}
