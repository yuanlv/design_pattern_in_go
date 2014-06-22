package factory


//宠物工厂生产各种宠物及宠物食品

//宠物超市从宠物工厂中获取产品

type Pet interface {
	Speek() string //宠物都会叫
	String() string //宠物都有一个名称
}

type Factory interface {
	getFood() string
	getPet() Pet
}

type PetShop struct {
	petFactory Factory
} 

type DogFactory struct{}
type CatFactory struct{}

type Dog struct{}
type Cat struct{}

func (p *Dog) Speek() string{
	return "woaf"
}

func (p *Dog) String() string{
	return "Dog"
}

func (p *Cat) Speek() string{
	return "meow"
}

func (p *Cat) String() string{
	return "Cat"
}

func (p *DogFactory) getFood() string {
	return "dog food"
}

func (p *DogFactory) getPet() Pet{
	return new(Dog)
}

func (p *CatFactory) getFood() string {
	return "cat food"
}

func (p *CatFactory) getPet() Pet{
	return new(Cat)
}

