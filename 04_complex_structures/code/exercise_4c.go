package main

import "fmt"

func avgScore(scores [5]float64) float64 {

	total := 0.0
	for score := range scores {
		total += float64(score)
	}
	return total / float64(len(scores))

}

var petNames map[string]string = map[string]string{
	"fido":     "dog",
	"spot":     "dog",
	"penelope": "horse",
}

func mapTest(name string) bool {

	_, ok := petNames[name]
	return ok
}

var gList = []string{"choco", "apple", "banana"}

func sliceTest(items ...string) []string {
	foods := gList
	for _, grocery := range items {
		foods = append(foods, grocery)
	}
	return foods
	// gslice := make([]string, 3, 10)
	// newThing := append(gslice, items)
	// return newThing

}

func main() {
	scores := [5]float64{1, 2, 3, 4, 5}
	// avgScore()
	fmt.Println(avgScore(scores))

	// fmt.Println(mapTest("fido"))
	pet := "fido"
	petExists := mapTest(pet)
	fmt.Println(petExists)
	//
	groceryList := sliceTest("yeast", "wine", "tuna")
	fmt.Println(groceryList)
}
