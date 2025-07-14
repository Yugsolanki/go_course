package main

import "fmt"

type Product struct {
	id    int
	title string
	price float64
}

func main() {
	// Task 1:
	hobbies := [3]string{"Tennis", "Art", "Baking"}
	fmt.Println(hobbies)

	// Task 2:
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// Task 3:
	task3 := hobbies[0:2]
	// task3 := hobbies[:2]
	fmt.Println(task3)

	// Task 4:
	fmt.Println(cap(task3))
	task4 := task3[1:3] // cap=3, so we can go to the right (end) always
	fmt.Println(task4)

	// Task 5:
	courseGoals := []string{"Learning GO!", "Learning all the basics"}
	fmt.Println(courseGoals)

	// Task 6:
	courseGoals[1] = "Learning all the details"
	courseGoals = append(courseGoals, "Learning Interfaces")
	fmt.Println(courseGoals)

	// Task 7:
	product1 := Product{
		id:    1,
		title: "Shampoo",
		price: 3.4,
	}
	product2 := Product{
		id:    2,
		title: "Conditioner",
		price: 2.6,
	}
	arr := []Product{product1, product2}
	fmt.Println(arr)

	product3 := Product{
		id:    3,
		title: "Cleanser",
		price: 3.6,
	}
	arr = append(arr, product3)
	fmt.Println(arr)
}
