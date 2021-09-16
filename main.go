// Go program to illustrate
// the type assertion
package main

import "fmt"

func myfun(a interface{}) {

	//fmt.Println( a.(type), a)
	// Using type switch
	switch a.(type) {

	case int:
		fmt.Println("Type: int, Value:", a.(int))
	case string:
		fmt.Println("\nType: string, Value: ", a.(string))
	case float64:
		fmt.Println("\nType: float64, Value: ", a.(float64))
	default:
		fmt.Println("\nType not found")
	}
}

type class struct {
	name string
	age  int
}

type animal struct {
	legs  int
}



type person interface {
	 getAgePlusTen() int

	getLegs() int
}

func (cs class) getAgePlusTen() int {
	return cs.age * 10
}

func (an animal)  getAgePlusTen() int {
	return an.legs
}

func (an animal) getLegs() int {
	return an.legs
}

func main() {

	val := 626.22
	var name string
	name = "asd"
	fmt.Println(name)
	myfun(val)

	var ahmed person

 	ahmed = animal{legs: 4}

	fmt.Println(ahmed.getLegs())

}
