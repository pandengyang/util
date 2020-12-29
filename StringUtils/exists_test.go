package StringUtils

import (
	"testing"
)

type TaskExistsInSlice struct {
	Slice    []string
	Value    string
	Expected bool
}

func TestExistsInSlice(t *testing.T) {
	task1 := TaskExistsInSlice{[]string{"a", "b", "c"}, "a", true}
	task2 := TaskExistsInSlice{[]string{"a", "b", "c"}, "d", false}

	t.Log("Given the need to test exists in slice.")
	{
		t.Logf("\tWhen checking exists in slice")
		{
			exists := ExistsInSlice(task1.Slice, task1.Value)
			if exists == task1.Expected {
				t.Logf("\t\tShould be able to produce %t.%s", task1.Expected, checkMark)
				t.Log("\t\t\t", exists)
			} else {
				t.Errorf("\t\tShould be able to produce %t.%s", task1.Expected, ballotX)
				t.Log("\t\t\t", exists)
			}
		}
	}

	t.Log("Given the need to test exists in slice.")
	{
		t.Logf("\tWhen checking exists in slice")
		{
			exists := ExistsInSlice(task2.Slice, task2.Value)
			if exists == task2.Expected {
				t.Logf("\t\tShould be able to produce %t.%s", task2.Expected, checkMark)
				t.Log("\t\t\t", exists)
			} else {
				t.Errorf("\t\tShould be able to produce %t.%s", task2.Expected, ballotX)
				t.Log("\t\t\t", exists)
			}
		}
	}
}
