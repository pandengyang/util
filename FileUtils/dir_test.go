package FileUtils

import (
	"testing"
)

type TaskDir struct {
	ParentPath string
	Expected   string
}

func TestMkdirAllByDate(t *testing.T) {
	task1 := TaskDir{"./uploads", "uploads/2020-07-21"}

	t.Log("Given the need to test whether mkdir by date.")
	{
		t.Logf("\tWhen checking whether mkdir by date")
		{
			path, _ := MkdirAllByDate(task1.ParentPath)
			if path == task1.Expected {
				t.Logf("\t\tShould be able to produce %v.%s", task1.Expected, checkMark)
				t.Logf("\t\t\t%v", path)
			} else {
				t.Errorf("\t\tShould be able to produce %v.%s", task1.Expected, ballotX)
				t.Logf("\t\t\t%v", path)
			}
		}
	}
}
