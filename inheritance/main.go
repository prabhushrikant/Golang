package main

import (
	ml "stash.experticity.com/go/TestProject_Golang/inheritance/multiLevelInheritance"
)

func main() {
	myDog := ml.MakeDog("tommy", "tominator")
	myDog.Eat()
}
