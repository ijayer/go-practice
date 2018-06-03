/*
 * 说明：
 * 作者：zhe
 * 时间：2018-05-14 15:06
 * 更新：
 */

package unit

import "testing"

func TestAdd(t *testing.T) {
	r := Add(1, 2)
	if r == 3 {
		t.Fail() // t.Fail 标记测试函数为失败，然后继续执行当前函数测试代码以及剩余所有的测试函数。
		t.Log("_______________________________TestAdd_Fail")
	} else {
		t.Log("add testing is ok")
	}
}

func TestAdd1(t *testing.T) {
	r := Add(1, 2)
	if r == 3 {
		t.FailNow() // t.FailNow 标记测试失败并且立即停止执行当前函数测试代码，继续执行下一个（默认按书写顺序）测试函数(文件)。
		t.Log("_______________________________TestAdd1_FailNow")
	} else {
		t.Log("add testing is ok")
	}
}

func TestAdd2(t *testing.T) {
	r := Add(1, 2)
	if r == 3 {
		t.Errorf("Errorf.TestAdd") // t.ErrorF = t.Logf + t. Fail
		t.Log("_______________________________TestAdd2_Errorf")
	} else {
		t.Log("add testing is ok")
	}
}

func TestAdd3(t *testing.T) {
	r := Add(1, 2)
	if r == 3 {
		t.Fatalf("Fatalf.TestAdd") // t.Fatalf = t.Logf + t.FailNow
		t.Log("_______________________________TestAdd3_Fatalf")
	} else {
		t.Log("add testing is ok")
	}
}

func TestSub(t *testing.T) {
	r := Sub(2, 1)
	if r == 1 {
		t.FailNow() // current file testing stopped, and continue next testing func
		t.Log("_______________________________TestSub_FailNow")
	} else {
		t.Log("sub testing is ok")
	}
}
