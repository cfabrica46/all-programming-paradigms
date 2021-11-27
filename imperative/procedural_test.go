package imperative

import (
	"fmt"
	"testing"
)

func TestStructured(t *testing.T) {
	fmt.Println()
	fmt.Println("--- PROCEDURAL ---")

	jobID1 := "1"
	jobID2 := "2"

	valueJob1 := "this is for job " + jobID1
	valueJob2 := "this is for job " + jobID2

	repetitionsJob1 := 5
	repetitionsJob2 := 3

	doJobStructured(jobID1, valueJob1, repetitionsJob1)

	doJobStructured(jobID2, valueJob2, repetitionsJob2)
}

func doJobStructured(jobID, valueJob string, repetitions int) {
	fmt.Println("START JOB " + jobID)
	for i := 0; i < repetitions; i++ {
		fmt.Println(valueJob)
	}
	fmt.Println("FINISH JOB " + jobID)
	fmt.Println()
}
