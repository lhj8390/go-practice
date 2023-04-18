package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Age:%d", s.Age)
}

// PrintAge 구체화된 타입 (Student) 로 타입 변환하는 예제
func PrintAge(stringer Stringer) {
	s := stringer.(*Student)
	fmt.Printf("Age: %d\n", s.Age)
}

type Actor struct {
}

func (a *Actor) String() string {
	return "Actor"
}

// ConvertType 타입 변환 시 발생하는 예외에 대한 예시
func ConvertType(stringer Stringer) {
	//student := stringer.(*Student) // 오류 발생 (*Student 타입이 아니기 때문에)
	student := stringer.(*Actor)
	fmt.Println(student)
}

func main() {
	s := &Student{15}
	PrintAge(s)

	actor := &Actor{}
	ConvertType(actor)
}
