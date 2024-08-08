package arrayslice

import "fmt"

/*
	Time to practice what you learned!
	1) Create a new array (!) that contains three hobbies you have
		Output (print) that array in the command line
	2) Also output more data about that array:
		- The first element (standalone)
		- The second and third element combined as a new list
	3) Create a slice based on the first element that contains the first and second elements
		Create that slice in two different ways (i.e. Create two slice in the end)
	4) Re-slice the slice from (3) and change it to contain the second and last element of original array
	5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	7) Bonus: Create a "Product" struct with title, id, price and create a dynamic list of products (at least 2 products).
		Then add a third product to existing list of products
*/

type Product struct {
	id    string
	title string
	price float64
}

func Array_slice() {
	//1)
	var Hobbies = [3]string{"Flutte", "Draw", "Music"}
	fmt.Println(Hobbies)

	//2)
	fmt.Println(Hobbies[0])
	fmt.Println(Hobbies[1:3])

	//3)
	var NewHobbies = Hobbies[0:2]

	fmt.Println(NewHobbies)

	//4)
	fmt.Println(cap(NewHobbies))
	NewHobbies = NewHobbies[1:3]
	fmt.Println(NewHobbies)

	//5)
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	//6)
	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics!")
	fmt.Println(courseGoals)

	//7)
	products := []Product{
		Product{
			"first-product",
			"A first product",
			12.99,
		},
		Product{
			"second-product",
			"A second product",
			129.99,
		},
	}

	fmt.Println(products)
	newProduct := Product{
		"third-product",
		"A third product",
		30,
	}

	products = append(products, newProduct)
	fmt.Println(products)
}
