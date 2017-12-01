package unit

import "testing"

func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r != 1 {
		t.Fail() // continue next testing func
	} else {
		t.Log("add test is ok")
	}
	t.Log("mark Fail")
}

func TestSub(t *testing.T) {
	r := Sub(1, 2)
	if r != 1 {
		t.FailNow() // current file testing stopped, and continue next testing func
	} else {
		t.Log("sub test is ok")
	}
	t.Log("mark FailNow")
}

func TestMul(t *testing.T) {
	r := Mul(1, 2)
	if r != 2 {
		t.Fatal() // = t.Log + t.FailNow()
	} else {
		t.Log("mul test is ok")
	}
}

func TestDiv(t *testing.T) {
	if r, err := Div(4, 2); err != nil || r != 2 { // usual case
		t.Errorf("div test failed: %v\n", err)
	} else {
		t.Log("div test is ok")
	}
}

/*
func TestDiv2(t *testing.T) {
	if _, err := Div(4, 0); err != nil { // exception test
		t.Errorf("div test failed: %v\n", err)
	}
	t.Log("div test is ok")
}*/
