package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Person cost
type x struct {
	A int64
	B int
	C int8
	E int8
	D float64
}

func main() {
	y := createStruct()
	fmt.Println("size", y, unsafe.Sizeof(x{}))
}

func genCombination() {}

func createStruct() uintptr {
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "A",
			Type: reflect.TypeOf(int64(0)),
			Tag:  `tag:"a"`,
		},
		{
			Name: "B",
			Type: reflect.TypeOf(int(0)),
			Tag:  `tag:"b"`,
		},
		{
			Name: "C",
			Type: reflect.TypeOf(int8(0)),
			Tag:  `tag:"c"`,
		},
		{
			Name: "E",
			Type: reflect.TypeOf(int8(0)),
			Tag:  `tag:"e"`,
		},
		{
			Name: "D",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `tag:"d"`,
		},
	})

	v := reflect.New(typ).Elem()
	v.Field(0).SetInt(0)
	v.Field(1).SetInt(0)
	v.Field(2).SetInt(0)
	v.Field(3).SetInt(0)
	v.Field(4).SetFloat(0)

	return v.Type().Size()
}
