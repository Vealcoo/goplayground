package reflect

import (
	"fmt"
	"reflect"
)

type T struct{}

func (t *T) Geeks(test string, testInt int) []string {
	fmt.Println("GeekforGeeks", test, testInt)
	return []string{"1", "2", "3", "4"}
}

func Run() {
	var t T

	// use of Call() method
	f := reflect.ValueOf(&t).MethodByName("Geeks")
	if !f.IsValid() {
		panic("invalid func")
	}

	for i, v := range f.Call([]reflect.Value{reflect.ValueOf("test"), reflect.ValueOf(1234)}) {
		fmt.Println(i, v)
		fmt.Println(v.Type())
	}
}
