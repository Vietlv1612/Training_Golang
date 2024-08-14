package main

func main() {
	// Define string variable
	var name string = "John Doe"
	// Print string variable
	println(name)
	// Function to get string length
	println(len(name))
	// Function to get string character by index
	println(string(name[0]))

	// Connect string with +
	var firstName string = "John"
	var lastName string = "Doe"
	var fullName string = firstName + " " + lastName
	println(fullName)
}
