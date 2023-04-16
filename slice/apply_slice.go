package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

// Sort() 함수 사용 위해 Len(), Less(), Swap() 세 메서드 필요

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func SortStructSlice() {
	s := []Student{
		{"화랑", 31}, {"백두산", 52}, {"류", 42},
		{"켄", 38}, {"송하나", 18},
	}

	sort.Sort(Students(s))
	fmt.Println("========= 구조체 슬라이스 정렬 =========")
	fmt.Println(s)
	fmt.Println()
}

func AddElement() {
	slice := []int{1, 2, 3, 4, 5, 6}

	slice = append(slice, 0)
	idx := 2 // 인덱스 2에 값 추가

	copy(slice[idx+1:], slice[idx:]) // 한 칸씩 뒤로 밀기
	slice[idx] = 100

	fmt.Println("============ 요소에 값 추가 ===========")
	checkValues("기존 요소", []int{1, 2, 3, 4, 5, 6})
	checkValues("slice", slice)
	fmt.Println()
}
