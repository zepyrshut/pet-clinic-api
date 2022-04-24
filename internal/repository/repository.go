package repository

import "github.com/zepyrshut/go-web-starter-gin/internal/models"

type DBRepo interface {
	// Pet
	OnePet(id int) (models.Pet, error)
	// AllPets() ([]models.Pet, error)

	// // Vetetarian
	// OneVetetarian(id int) (models.Veterinarian, error)
	// AllVeterinarians() ([]models.Veterinarian, error)

	// // PetOwner
	// OnePetOwner(id int) (models.PetOwner, error)
	// AllPetOwners() ([]models.PetOwner, error)
}
