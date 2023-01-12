package main

import "fmt"

func main() {
	intSlice := []int{}
	intSlice = append(intSlice, 1)
	intSlice = append(intSlice, 2)
	intSlice = append(intSlice, 3)

	for _, element := range intSlice {
		fmt.Println(element)
	}

	strSlice := make([]string, 5)
	strSlice = append(strSlice, "S")
	strSlice = append(strSlice, "L")
	strSlice = append(strSlice, "I")
	strSlice = append(strSlice, "C")
	strSlice = append(strSlice, "E")

	for i, _ := range strSlice {
		fmt.Print(strSlice[i])
	}

}
