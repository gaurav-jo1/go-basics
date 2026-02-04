package main

import "fmt"

func main() {
	var myname string = "Gaurav"
	var myage int16 = 21

	name := "Alice"
	age := 30
	// height := 5.6

	fmt.Print("Hello", "World")
	fmt.Print("!\n")

	fmt.Println("Hello", "World")
	fmt.Println("Name:", name)

	fmt.Printf("Name: %s, Age: %d\n", name, age)

	message := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println(message)

	fmt.Printf("%v\n", name)  // %v - default format
	fmt.Printf("%T\n", age)   // %T - type
	fmt.Printf("%t\n", true)  // %t - boolean
	fmt.Printf("%d\n", 42)    // %d - integer
	fmt.Printf("%f\n", 3.14)  // %f - float
	fmt.Printf("%s\n", "hi")  // %s - string
	fmt.Printf("%p\n", &name) // %p - pointer

	new_message := fmt.Sprintf("My name is %s and I am %d years old", myname, myage)
	fmt.Println(new_message)
}
