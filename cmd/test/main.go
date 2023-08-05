package main

import (
	"fmt"

	"fyne.io/fyne/v2/data/binding"
	"golang.org/x/exp/slices"
)

type person struct {
	Name string
	Age  int
}
type personList []person

func printList(pl personList) {
	fmt.Println("**********")
	for i := range pl {
		fmt.Printf("%v\n", pl[i])
	}
	fmt.Println("**********")
}

func main() {
	// a := app.New()
	// w := a.NewWindow("Hello")

	// w.Resize(fyne.Size{Width: 400, Height: 400})
	// w.ShowAndRun()
	left := []person{{Name: "John Doe", Age: 21}, {Name: "Fred Smith", Age: 44}}
	numPersons := len(left)
	leftBindings := make([]binding.Struct, numPersons)
	for i := 0; i < len(left); i++ {
		leftBindings[i] = binding.BindStruct(&left[i])
		leftBindings[i].AddListener(binding.NewDataListener(func() {
			fmt.Println("Left item changed")
		}))
	}
	right := []person{}
	rightBindings := make([]binding.Struct, numPersons)
	for i := 0; i < len(right); i++ {
		rightBindings[i] = binding.BindStruct(&right[i])
		rightBindings[i].AddListener(binding.NewDataListener(func() {
			fmt.Println("Right item changed")
		}))
	}

	fmt.Println("Before")
	printList(left)
	printList(right)

	leftBindings[0].SetValue("Name", "hjhhgjg")
	rightBindings[0] = leftBindings[0]
	leftBindings = slices.Delete(leftBindings, 0, 0)

	fmt.Println("After")
	printList(left)
	printList(right)

	// src := person{Name: "John Doe", Age: 21}
	// values := binding.BindStruct(&src)
	// fmt.Println("Map size:", len(values.Keys()))
	// name, _ := values.GetValue("Name")
	// fmt.Println("Value for Name:", name)
	// age, _ := values.GetValue("Age")
	// fmt.Println("Value for Age:", age)
}
