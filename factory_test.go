package factory

import "testing"



func TestSpeek(t *testing.T) {
	factory := new(CatFactory)
	shop := PetShop{}

	shop.petFactory = factory
	pet := shop.petFactory.getPet()

	if speek := pet.Speek(); speek == "woaf" {
		t.Errorf("should get moaf, get %s", speek)
	}


	dogFactory := new(DogFactory)
	shop = PetShop{}

	shop.petFactory = dogFactory
	pet = shop.petFactory.getPet()

	if speek := pet.Speek(); speek != "woaf" {
		t.Errorf("should get woaf, get %s", speek)
	}
}

