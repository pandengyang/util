package FileUtils

import (
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

type Task struct {
	Path   string
	Expected bool
}

func TestFileExist(t *testing.T) {
	task1 := Task{"a", true}

	t.Log("Given the need to test producing hashed and salted password.")
	{
		t.Logf("\tWhen checking producing hashed and salted password")
		{
			exist, _ := FileExist(task1.Path)
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
