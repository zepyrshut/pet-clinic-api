package repository

import "github.com/zepyrshut/pet-clinic/internal/models"

type DBRepo interface {
	// Pet
	NewPet(pet models.Pet) error       
	OnePet(id int) (models.PetDTO, error) 
	//UpdatePet(pet models.Pet) error    
	//DeletePet(id int) error            

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
