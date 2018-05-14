package unit

import "testing"

func TestMul(t *testing.T) {
	r := Mul(1, 2)
	if r == 2 {
		t.Errorf("Errorf.TestMul") // = t.Log + t.FailNow()
	} else {
		t.Log("mul testing is ok")
	}
}

func TestDiv(t *testing.T) {
	r, err := Div(4, 2)
	if err != nil || r == 2 { // usual case
		t.Fatalf("Fatalf.TestDiv Error: %v\n", err)
	} else {
		t.Log("div testing is ok")
	}
}

/*
func TestDiv2(t *testing.T) {
	if _, err := Div(4, 0); err != nil { // exception testing
		t.Errorf("div testing failed: %v\n", err)
	}
	t.Log("div testing is ok")
}*/
