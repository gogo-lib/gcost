package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Person cost
type x struct {
	a int16
	b int64
	c int32
}

type y struct {
	b int64
	c int32
	a int16
}

func main() {
	// typ := reflect.StructOf([]reflect.StructField{{
	// 	Name: "Height",
	// 	Type: reflect.TypeOf(int64(0)),
	// 	Tag:  `cost:"height"`,
	// }})
	t := reflect.StructField{
		Name: "Height",
		Type: reflect.TypeOf(int64(0)),
		Tag:  `cost:"height"`,
	}

	// v := reflect.New(typ).Elem()
	// fmt.Println(v.Kind(), unsafe.Sizeof(v))
	fmt.Println(unsafe.Sizeof(t), unsafe.Sizeof(y{}))
	rt := t.Type
	for i := 0; i < rt.NumField(); i++ {
		fmt.Println(rt.Field(i).Name)
	}

}
