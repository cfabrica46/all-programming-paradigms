package imperative

import (
	"fmt"
	"testing"
)

func TestFunctional(t *testing.T) {
	fmt.Println()
	fmt.Println("--- STRUCTURED ---")
	valueJob1 := "this is for job 1"
	valueJob2 := "this is for job 2"

	repetitionsJob1 := 5
	repetitionsJob2 := 3

	fmt.Println("START JOB 1")
	for i := 0; i < repetitionsJob1; i++ {
		fmt.Println(valueJob1)
	}
	fmt.Println("FINISH JOB 1")

	fmt.Println()

	fmt.Println("START JOB 2")
	for i := 0; i < repetitionsJob2; i++ {
		fmt.Println(valueJob2)
	}
	fmt.Println("FINISH JOB 2")

	fmt.Println()
	fmt.Println()
}
