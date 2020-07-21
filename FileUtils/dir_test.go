package FileUtils

import (
	"testing"
)

type TaskDir struct {
	ParentDir string
	err       error
}

func TestMkdirAllByDate(t *testing.T) {
	task1 := TaskDir{"./uploads", nil}

	t.Log("Given the need to test whether mkdir by date.")
	{
		t.Logf("\tWhen checking whether mkdir by date")
		{
			err := MkdirAllByDate(task1.ParentDir)
			if err == task1.err {
				t.Logf("\t\tShould be able to produce %v.%s", task1.err, checkMark)
				t.Logf("\t\t\t%v", err)
			} else {
				t.Errorf("\t\tShould be able to produce %v.%s", task1.err, ballotX)
				t.Logf("\t\t\t%v", err)
			}
		}
	}
}
