package declarative

import (
	"fmt"
	"testing"
)

type myJob struct {
	ID          int
	Job         func()
	Repetitions int
}

func newJob(id, repetitions int, job func()) myJob {
	return myJob{ID: id, Job: job, Repetitions: repetitions}
}

func (j myJob) do() {
	fmt.Printf("START JOB %d\n", j.ID)
	for i := 0; i < j.Repetitions; i++ {
		j.Job()
	}
	fmt.Printf("FINISH JOB %d\n", j.ID)
	fmt.Println()
}

func TestObjectOriented(t *testing.T) {
	fmt.Println()
	fmt.Println("--- OBJECT-ORIENTED ---")

	job1 := newJob(1, 5, func() { fmt.Println("this is for job 1") })
	job2 := newJob(2, 3, func() { fmt.Println("this is for job 2") })

	job1.do()
	job2.do()
}
