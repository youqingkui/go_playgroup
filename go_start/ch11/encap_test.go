package ch11

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

// 实例的成员会进行值复制
// func (e Employee) String() string {
// 	fmt.Printf("Address is %x \n", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s Name:%s Age:%d", e.Id, e.Name, e.Age)
// }

// 实例使用指针避免内存拷贝-通常使用该方式
func (e *Employee) String() string {
	fmt.Printf("Address is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s Name:%s Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x \n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 29}
	e2 := new(Employee) // 返回指针
	e2.Id = "2"
	e2.Age = 28
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e2)

	t.Logf("e is %T", e)
	t.Logf("e2 is %T", e2)
}
