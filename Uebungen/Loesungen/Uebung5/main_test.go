package main

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	result, err := divide(4, 2)
	if err != nil {
		fmt.Errorf("Unerwarteter Fehler")
	}

	if result != 2 {
		fmt.Errorf("Erwartet %f, Ergebnis %f", 2.0, result)
	}
}

func TestDivideWithError(t *testing.T) {
	_, err := divide(4, 0)
	if err == nil {
		fmt.Errorf("Fehler bei Division durch 0 erwartet")
	}
}
