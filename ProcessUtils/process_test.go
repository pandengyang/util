package ProcessUtils

import (
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

type Task struct {
	Pid      int
	Expected bool
}

func TestProcessExists(t *testing.T) {
	task1 := Task{49850, false}

	t.Log("Given the need to test whether process exists.")
	{
		t.Logf("\tWhen checking whether process exists")
		{
			exist := ProcessExists(task1.Pid)
			if exist == task1.Expected {
				t.Logf("\t\tShould be able to produce %v.%s", task1.Expected, checkMark)
				t.Logf("\t\t\t%v", exist)
			} else {
				t.Errorf("\t\tShould be able to produce %v.%s", task1.Expected, ballotX)
				t.Logf("\t\t\t%v", exist)
			}
		}
	}
}
