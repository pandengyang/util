package StringUtils

import (
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

type Task struct {
	Passwd string
	Salt string
	Expected string
}

func TestSha256PasswdSalt(t *testing.T) {
	task1 := Task{"123456", "salt", "73616C748D969EEF6ECAD3C29A3A629280E686CF0C3F5D5A86AFF3CA12020C923ADC6C92"}

	t.Log("Given the need to test producing hashed and salted password.")
	{
		t.Logf("\tWhen checking producing hashed and salted password")
		{
			hashedSaltedPasswd := Sha256PasswdSalt(task1.Passwd, task1.Salt)
			if hashedSaltedPasswd == task1.Expected {
				t.Logf("\t\tShould be able to produce %s.%s", task1.Expected, checkMark)
				t.Log("\t\t\t", hashedSaltedPasswd)
			} else {
				t.Errorf("\t\tShould be able to produce %s.%s", task1.Expected, ballotX)
				t.Log("\t\t\t", hashedSaltedPasswd)
			}
		}
	}
}
