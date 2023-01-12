package main

import "fmt"

func main() {

}

func divide(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("Division durdch 0 ist nicht erlaubt!")
	}

	return a / b, nil
}
