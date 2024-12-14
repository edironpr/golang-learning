package main

import (
	"testing"
)

// 单元测试
func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("Add(1, 2) expected be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("Add(10, -20) expected be -30, but %d got", ans)
	}
}

// 子测试
func TestMul(t *testing.T) {

	// .Error 遇错不停，还会继续执行其他的测试用例
	// .Fatal 遇错即停
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Fatal("fail")
		}
	})
	t.Run("neg", func(t *testing.T) {
		if Mul(2, -3) != -6 {
			t.Fatal("fail")
		}
	})

	// table-driven tests
	// 好处：
	// 	新增用例非常简单，只需给 cases 新增一条测试数据即可。
	// 	测试代码可读性好，直观地能够看到每个子测试的参数和期待的返回值。
	// 	用例失败时，报错信息的格式比较统一，测试报告易于阅读。
	// 如果数据量较大，或是一些二进制数据，推荐使用相对路径从文件中读取。
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 1, 2, 2},
		{"neg", -1, 2, -2},
		{"zero", 0, 2, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, ans)
			}
		})
	}

}

type CalcCase struct {
	A, B, Expected int
}

// 帮助函数
func createDivTestCase(t *testing.T, c *CalcCase) {
	// helper
	// 用于标注该函数是帮助函数，报错时将输出帮助函数调用者的信息，而不是帮助函数的内部信息。帮助函数在多处调用，如果在帮助函数内部报错，报错信息都在同一处，不方便问题定位。
	t.Helper()

	if ans := Div(c.A, c.B); ans != c.Expected {
		t.Fatalf("%d * %d expected %d, but %d got", c.A, c.B, c.Expected, ans)
	}
}

func TestDiv(t *testing.T) {
	// 帮助函数使用建议：
	// 1. 不要返回错误，帮助函数内部直接使用 t.Error 或 t.Fatal 即可，在用例主逻辑中不会因为太多的错误处理代码，影响可读性。
	// 2. 调用 t.Helper() 让报错信息更准确，有助于定位。
	createDivTestCase(t, &CalcCase{4, 2, 2})
	createDivTestCase(t, &CalcCase{-4, 2, -2})
	createDivTestCase(t, &CalcCase{0, 2, 2}) // wrong case
}
