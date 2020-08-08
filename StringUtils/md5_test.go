package StringUtils

import (
	"testing"
)

type TaskMd5 struct {
	Content  string
	Expected string
}

func TestMd5(t *testing.T) {
	task1 := TaskMd5{"pandengyang", "4FC27A297282A239DE5DA231EBEF0131"}

	t.Log("Given the need to test producing md5.")
	{
		t.Logf("\tWhen checking producing md5")
		{
			checksum := Md5(task1.Content)
			if checksum == task1.Expected {
				t.Logf("\t\tShould be able to produce %s.%s", task1.Expected, checkMark)
				t.Log("\t\t\t", checksum)
			} else {
				t.Errorf("\t\tShould be able to produce %s.%s", task1.Expected, ballotX)
				t.Log("\t\t\t", checksum)
			}
		}
	}
}
