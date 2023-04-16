package main

import "fmt"

func checkValues(name string, slice []int) {
	fmt.Println(name+":", slice, len(slice), cap(slice))
}

func EmptySpace() {
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)

	fmt.Println("====== 슬라이스에 빈 공간이 있을 경우 ======")
	checkValues("slice1", slice1)
	checkValues("slice2", slice2)

	slice1[1] = 100

	fmt.Println("After change second element")
	checkValues("slice1", slice1)
	checkValues("slice2", slice2)

	slice1 = append(slice1, 500)
	fmt.Println("After append 500")
	checkValues("slice1", slice1)
	checkValues("slice2", slice2)
	fmt.Println()
}

func NoEmptySpace() {
	slice3 := []int{1, 2, 3}
	slice4 := append(slice3, 4, 5)

	fmt.Println("====== 슬라이스에 빈 공간이 없을 경우 ======")
	checkValues("slice3", slice3)
	checkValues("slice4", slice4)

	slice3[1] = 100
	fmt.Println("After change second element")
	checkValues("slice3", slice3)
	checkValues("slice4", slice4)

	slice3 = append(slice3, 500)
	fmt.Println("After append 500")
	checkValues("slice3", slice3)
	checkValues("slice4", slice4)
	fmt.Println()
}

func UseSpacing() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]

	fmt.Println("============ 슬라이싱 사용 ============")
	fmt.Println("array: ", array)
	checkValues("slice", slice)

	array[1] = 100
	fmt.Println("After change second element")
	fmt.Println("array: ", array)
	checkValues("slice", slice)

	slice = append(slice, 500)
	fmt.Println("After append 500")
	fmt.Println("array: ", array)
	checkValues("slice", slice)
	fmt.Println()
}
