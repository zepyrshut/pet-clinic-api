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

// Entities
type Pet struct {
	ID        uint      `json:"id"`
	PetName   string    `json:"pet_name" form:"pet_name" binding:"required"`
	PetType   string    `json:"pet_type" form:"pet_type" binding:"required"`
	PetRace   string    `json:"pet_race" form:"pet_race"`
	BirthDate time.Time `json:"birth_date" form:"birth_date" binding:"required" time_format:"2006-01-02"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Person struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name" form:"first_name" binding:"required"`
	LastName  string    `json:"last_name" form:"last_name"`
	Phone     string    `json:"phone" form:"phone" binding:"required,e164"`
	Email     string    `json:"email" form:"email" binding:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PetOwner struct {
	ID        uint      `json:"id"`
	PetID     uint      `json:"pet_id" form:"pet_id" binding:"required"`
	PersonID  uint      `json:"person_id" form:"person_id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Veterinarian struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name" form:"first_name" binding:"required"`
	LastName  string    `json:"last_name" form:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Visit struct {
	ID           uint         `json:"id"`
	Date         time.Time    `json:"date" form:"date" binding:"required" time_format:"2006-01-02"`
	Pet          Pet          `json:"pet_id" form:"pet_id" binding:"required"`
	Veterinarian Veterinarian `json:"veterinarian_id" form:"veterinarian_id" binding:"required"`
	Reason       string       `json:"reason" binding:"required"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

// DTO
type PersonDTO struct {
	ID        uint         `json:"id"`
	FirstName string       `json:"first_name" form:"first_name" binding:"required"`
	LastName  string       `json:"last_name" form:"last_name"`
	Phone     string       `json:"phone" form:"phone" binding:"required,e164"`
	Email     string       `json:"email" form:"email" binding:"email"`
	Pets      map[uint]Pet `json:"pets"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
