package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

//import "os"

func main() {
	a := 1
	b := 5
	fmt.Println("hello world")
	fmt.Println(cal(a, b))

	lowestNumber := findLowest(3, 4, 7, 1, 2, 9)

	fmt.Println(lowestNumber)

	addRange([]int{9, 9, 9, 92, 222})
	//map
	test := make(map[string]int)
	test["test"] = 4
	test["new"] = 100
	delete(test, "new")

	fmt.Printf("\nThe result is %v\n", test["test"])
	//struct

	type person struct {
		firstName string
		lastName  string
	}
	personx := person{"guo", "huang"}
	fmt.Println(personx)

	//go routine
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("Sleeping")
	}()

	go func() {
		defer waitGroup.Done()
		fmt.Println("done")
	}()
	fmt.Println("end")
	waitGroup.Wait()

	folder, _ := os.Open()
	defer folder.Close()

	fmt.Println(folder)
}

func cal(a, b int) int {
	return a + b
}

func findLowest(num ...int) int {
	best := num[0]
	for _, i := range num {
		if i < best {
			best = i
		}
	}

	return best
}

func addRange(num []int) {
	addNums := []int{1, 2, 3, 4}
	newNum := append(num, addNums...)
	fmt.Println(newNum)
}
