package declarative

import (
	"fmt"
	"testing"
)

func TestFunctional(t *testing.T) {
	fmt.Println()
	fmt.Println("--- FUNCTIONAL ---")

	job1 := setJobFunctional("this is for job 1")
	job2 := setJobFunctional("this is for job 2")

	repetitionsJob1 := 5
	repetitionsJob2 := 3

	fmt.Println("START JOB 1")
	doJobFunctional(job1, repetitionsJob1)
	fmt.Println("FINISH JOB 1")

	fmt.Println()

	fmt.Println("START JOB 2")
	doJobFunctional(job2, repetitionsJob2)
	fmt.Println("FINISH JOB 2")

	fmt.Println()
}

func setJobFunctional(value string) func() {
	return func() {
		fmt.Println(value)
	}
}

func doJobFunctional(job func(), repetitions int) {
	for i := 0; i < repetitions; i++ {
		job()
	}
}
