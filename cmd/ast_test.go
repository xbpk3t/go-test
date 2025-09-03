package cmd

import (
	"fmt"
	"testing"
)

func TestAST(t *testing.T) {
	configuration := "zzz"
	fmt.Println(configuration)

	fmt.Println(ModConfiguration())
}

func ModConfiguration() bool {
	return true
}
