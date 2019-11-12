package CollectionJSON

import (
	"testing"
	"io/ioutil"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

var pTaskTemplateStr *string

type Task struct {
	Id int
	Description string
	DateDue string
	Completed bool
}

func init() {
	content, err := ioutil.ReadFile("template/v1/task.json")
	if err != nil {
		panic("read template error!")
	}

	contentStr := string(content)
	pTaskTemplateStr = &contentStr
}

func TestItems(t *testing.T) {
	task1 := Task{1, "买书", "2018-09-23", false}
	task2 := Task{2, "吃月饼", "2018-10-23", false}
	tasks := [] interface{} {
		task1, task2,
	}

	t.Log("Given the need to test producing items response.")
	{
		t.Logf("\tWhen checking producing items response")
		{
			jsonStr, err := Items(tasks, pTaskTemplateStr)

			if err == nil {
				t.Log("\t\tShould be able to produce items response.", checkMark)
				t.Log("\t\t\t", jsonStr)
			} else {
				t.Error("\t\tShould be able to produce items response.", ballotX)
			}
		}
	}
}

func TestItem(t *testing.T) {
	task1 := Task{1, "买书", "2018-09-23", false}

	t.Log("Given the need to test producing item response.")
	{
		t.Logf("\tWhen checking producing item response")
		{
			jsonStr, err := Item(task1, pTaskTemplateStr)

			if err == nil {
				t.Log("\t\tShould be able to produce item response.", checkMark)
				t.Log("\t\t\t", jsonStr)
			} else {
				t.Log(err.Error())
				t.Error("\t\tShould be able to produce item response.", ballotX)
			}
		}
	}
}

func TestQueries(t *testing.T) {
	t.Log("Given the need to test producing queries response.")
	{
		t.Logf("\tWhen checking producing queries response")
		{
			jsonStr, err := Queries(pTaskTemplateStr)

			if err == nil {
				t.Log("\t\tShould be able to produce queries response.", checkMark)
				t.Log("\t\t\t", jsonStr)
			} else {
				t.Log(err.Error())
				t.Error("\t\tShould be able to produce queries response.", ballotX)
			}
		}
	}
}

func TestTemplate(t *testing.T) {
	t.Log("Given the need to test producing template response.")
	{
		t.Logf("\tWhen checking producing template response")
		{
			jsonStr, err := Template(pTaskTemplateStr)

			if err == nil {
				t.Log("\t\tShould be able to produce template response.", checkMark)
				t.Log("\t\t\t", jsonStr)
			} else {
				t.Log(err.Error())
				t.Error("\t\tShould be able to produce template response.", ballotX)
			}
		}
	}
}
