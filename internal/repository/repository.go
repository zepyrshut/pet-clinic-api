package repository

import "github.com/zepyrshut/pet-clinic/internal/models"

type DBRepo interface {
	// Pet
	NewPet(pet models.Pet) error       // C
	OnePet(id int) (models.Pet, error) // R
	//UpdatePet(pet models.Pet) error    // U
	//DeletePet(id int) error            // D

	//AllPets() ([]models.Pet, error)

	// // Vetetarian
	// OneVetetarian(id int) (models.Veterinarian, error)
	// AllVeterinarians() ([]models.Veterinarian, error)

	// // Person
	NewPerson(petOwner models.Person) error
	OnePerson(id int) (models.PersonDTO, error)
	// AllPeople() ([]models.person, error)

	// PetOwner
	BindPetWithOwner(petID int, ownerID int) error
}
