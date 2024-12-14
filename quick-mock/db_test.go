package main

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 使用 ctrl.Finish() 断言 DB.Get() 被是否被调用，如果没有被调用，后续的 mock 就失去了意义

	m := NewMockDB(ctrl) // NewMockDB() 的定义在 db_mock.go 中，由 mockgen 自动生成
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist")) // 当 Get() 的参数为 Tom，则返回 error

	// stubs 打桩

	// 参数 Eq, Any, Not, Nil
	//	Eq(value) 表示与 value 等价的值
	//	Any() 可以用来表示任意的入参
	//	Not(value) 用来表示非 value 以外的值
	//	Nil() 表示 None 值
	//m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	m.EXPECT().Get(gomock.Any()).Return(630, nil)
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
	m.EXPECT().Get(gomock.Nil()).Return(0, errors.New("nil"))

	// 返回值 Return, DoAndReturn
	//	Return 返回确定的值
	//	Do Mock 方法被调用时，要执行的操作吗，忽略返回值
	//	DoAndReturn 可以动态地控制返回值
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil)
	m.EXPECT().Get(gomock.Any()).Do(func(key string) {
		t.Log(key)
	})
	m.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (int, error) {
		if key == "Sam" {
			return 630, nil
		}
		return 0, errors.New("not exist")
	})

	// 调用次数 Times
	//	Times() 断言 Mock 方法被调用的次数。
	//	MaxTimes() 最大次数。
	//	MinTimes() 最小次数。
	//	AnyTimes() 任意次数（包括 0 次）。
	m.EXPECT().Get(gomock.Not("Sam")).Return(0, nil).Times(2)

	// 调用顺序 InOrder
	o1 := m.EXPECT().Get(gomock.Eq("Tom")).Return(0, errors.New("not exist"))
	o2 := m.EXPECT().Get(gomock.Eq("Sam")).Return(630, nil)
	gomock.InOrder(o1, o2)

	if v := GetFromDB(m, "Tom"); v != -1 { // 测试方法 GetFromDB() 的逻辑是否正确(如果 DB.Get() 返回 error，那么 GetFromDB() 返回 -1)
		t.Fatal("expected -1, but got", v)
	}
}

// $ go test . -cover -v

// 如何编写可 mock 的代码
// - mock 作用的是接口，因此将依赖抽象为接口，而不是直接依赖具体的类。
// - 不直接依赖的实例，而是使用依赖注入降低耦合性。

// 如果 GetFromDB() 方法长这个样子：
//func GetFromDB(key string) int {
//	db := NewDB()
//	if value, err := db.Get(key); err == nil {
//		return value
//	}
//
//	return -1
//}
// 对 DB 接口的 mock 并不能作用于 GetFromDB() 内部，这样写是没办法进行测试的。那如果将接口 db DB 通过参数传递到 GetFromDB()，那么就可以轻而易举地传入 Mock 对象了