package main

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"unsafe"
)

// Person cost
type x struct {
	a int8
	b int64
	c int16
}

func main() {
	rt := reflect.TypeOf(x{})

	start := make([]int, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		start[i] = i
	}
	calcSizeFunc := calcSizeFunc(rt)
	var min uintptr = math.MaxUint
	rs := map[uintptr][]string{}
	for start != nil {
		size, rawStructTmp := calcSizeFunc(start)
		if size <= min {
			rs[size] = append(rs[size], rawStructTmp)
			min = size
		}
		start = nextPermutation(start)
	}
	fmt.Println("current size: ", unsafe.Sizeof(x{}))
	fmt.Println("min size: ", min)
	fmt.Println("variations:")
	for _, v := range rs[min] {
		fmt.Println(v)
	}

}

func nextPermutation(nums []int) []int {
	numsLen := len(nums)
	first := -1
	second := -1

	for i, j := numsLen-2, numsLen-1; i >= 0; {
		if nums[i] < nums[j] {
			first = i
			second = j
			break
		} else {
			i--
			j--
		}
	}

	if !(first == -1) {
		smallestGreaterIndex := second
		for i := second + 1; i < numsLen; i++ {
			if nums[i] > nums[first] && nums[i] < nums[smallestGreaterIndex] {
				smallestGreaterIndex = i
			}
		}
		nums[first], nums[smallestGreaterIndex] = nums[smallestGreaterIndex], nums[first]

		sort.Slice(nums[second:numsLen], func(i, j int) bool {
			return nums[second+i] < nums[second+j]
		})
	} else {
		return nil
	}

	return nums
}

func calcSizeFunc(rt reflect.Type) func(order []int) (uintptr, string) {
	var listStructField []reflect.StructField
	for i := 0; i < rt.NumField(); i++ {
		listStructField = append(listStructField, rt.Field(i))
	}

	return func(order []int) (uintptr, string) {
		var tmpStructField []reflect.StructField
		for _, v := range order {
			tmpStructField = append(tmpStructField, reflect.StructField{
				Name:    listStructField[v].Name,
				Type:    listStructField[v].Type,
				PkgPath: listStructField[v].PkgPath,
			})
		}
		structType := reflect.StructOf(tmpStructField)

		return structType.Size(), structType.String()
	}

}
