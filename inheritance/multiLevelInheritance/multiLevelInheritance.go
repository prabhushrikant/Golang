package multiLevelInheritance

import "fmt"

type animal interface{
	Eat()
}

//type wildAnimal interface {
//	animal
//	Hunt()
//}

type animaImpl struct {

}

func (a animaImpl) Eat() {
	fmt.Println("I eat anything")
}

type domesticAnimal struct {
	animal
	familyName string
}

type dog struct{
	domesticAnimal
	name string
}

func MakeDog(s string, fs string) dog {
	return dog{
		domesticAnimal{
			animal: animaImpl{},
			familyName: fs,
		},
		s,
	}
}

func main() {
	myDog := MakeDog("tommy", "tominator")
	myDog.Eat()
}