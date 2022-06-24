package main

import "fmt"

const (
	firstName = iota + 6
	secondName
	Thirdname
	fifth
)

func main() {

	fmt.Println(firstName, secondName, Thirdname, fifth)

	//----------
	var name string = "Pablo"
	name1 := "Jose"
	var name2 int
	name2 = 50
	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(name2)

	const pi = 3.14
	var newNumb int = 5
	fmt.Println(pi + float32(newNumb))
	//---------------------------------------

	var number1 *int = new(int)
	*number1 = 52
	fmt.Println(*number1)
	fmt.Println(&number1)
	//----------------------------------------

	teach1 := "Bobby"
	fmt.Println(teach1)
	fmt.Println(&teach1)

	ptr := &teach1

	teach1 = "Maria"
	fmt.Println(ptr, *ptr)

	//----------------------------------------

}
