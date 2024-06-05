package main

import (
	"fmt"
	"math/rand"
)

type Name string
type AnimalType string
type CageNomber int

type InCage bool

var AnimalOk int
var AnimalNotOk int
var AnimalsInZoo int

type Animal struct {
	Name       Name
	AnimalType AnimalType

	InCage bool

	Cage
}

type Cage struct {
	CageNomber int
}

type ZooKeeper struct {
	Name   Name
	AtWork bool
}

var animalslist []Animal
var keeperlist []ZooKeeper

func main() {

	keeperlist = []ZooKeeper{
		{"Antony", true},
		{"Marta", true},
	}

	animalslist = []Animal{
		{"Leo", "Lion", true, (Cage{1})},
		{"Liza", "Lion", false, (Cage{1})},
		{"Wiwi", "Wolf", true, (Cage{2})},
		{"Wolly", "Wokf", true, (Cage{2})},
		{"Bery", "Bear", false, (Cage{3})},
		{"Crocos", "Crocodile", false, (Cage{4})},
		{"Piggy", "Penguin", true, (Cage{5})},
		{"Peppy", "Penguin", true, (Cage{5})},
		{"Pong", "Penguin", false, (Cage{5})},
		{"Rin", "Rabbit", true, (Cage{6})},
		{"Ron", "Rabbit", false, (Cage{6})},
		{"Ronny", "Rabbit", false, (Cage{6})},
		{"Elly", "Elephant", true, (Cage{7})},
		{"Sassy", "Snake", false, (Cage{8})},
		{"King", "Owl", true, (Cage{9})},
		{"Night", "Owl", true, (Cage{9})},
		{"Mimo", "Mouse", false, (Cage{10})},
	}

	ZooStatus()
	ZooKeeperWork()

}

func ZooStatus() {

	for _, animal := range animalslist {

		if !animal.InCage {
			fmt.Printf("%s the %s ran away from cage %d\n", animal.Name, animal.AnimalType, animal.Cage.CageNomber)
			AnimalNotOk++
		} else {
			fmt.Printf("%v the %s in cage %v.\n", animal.Name, animal.AnimalType, animal.Cage.CageNomber)
			AnimalOk++
		}
	}

	AnimalsInZoo = AnimalOk + AnimalNotOk

	fmt.Printf("\nAnimal in cage:%v\nAnimal ran away:%v\nNumbers of animals in the zoo:%v\n", AnimalOk, AnimalNotOk, AnimalsInZoo)

}

func ZooKeeperWork() {

	AnimalOk = 0
	AnimalNotOk = 0
	AnimalsInZoo = 0

	for _, animal := range animalslist {
		if !animal.InCage {

			randomKeeper := keeperlist[rand.Intn(len(keeperlist))]

			fmt.Printf("%v train to cach %v.\n", randomKeeper.Name, animal.Name)
			animal.InCage = true
			fmt.Printf("%v now in cage %v.\n", animal.Name, animal.Cage.CageNomber)

		}

		AnimalOk++

		if !animal.InCage { // it is here because a more complex scheme would need it
			AnimalNotOk++
		}
	}

	AnimalsInZoo = AnimalOk + AnimalNotOk

	fmt.Printf("\nAnimal in cage:%v\nAnimal ran away:%v\nNumbers of animals in the zoo:%v\n", AnimalOk, AnimalNotOk, AnimalsInZoo)
}
