package main

import "fmt"

func main() {
	res, err := divide(4, 2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Ergebnis: %f\n", res)
	}

	res, err = divide(4, 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	input := "Hello"
	expandString(&input)
	fmt.Println(input)
}

func divide(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division durdch 0 ist nicht erlaubt!")
	}

	return a / b, nil
}

func expandString(str *string) {
	*str = *str + " World"
}
