package FileUtils

import (
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

type Task struct {
	Path     string
	Expected bool
}

func TestFileExists(t *testing.T) {
	task1 := Task{"", true}

	t.Log("Given the need to test whether file exists.")
	{
		t.Logf("\tWhen checking whether file exists")
		{
			exist, _ := FileExists(task1.Path)
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
