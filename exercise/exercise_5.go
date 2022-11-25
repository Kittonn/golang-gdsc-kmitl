package exercise

import "fmt"

type Student struct {
	Student_ID string
	Address    *string
}

func (student *Student) ChangeAddress(address string) {
	student.Address = &address
}

func (student *Student) ChangeStudentID(studentID string) {
	student.Student_ID = studentID
}

func Exercise_5() {
	address := "MVK"
	student := Student{
		Student_ID: "4638",
		Address:    &address,
	}
	student.ChangeAddress("KMITL")
	student.ChangeStudentID("65010077")

	fmt.Println(student.Student_ID)
	fmt.Println(*student.Address)
	fmt.Println("-----")
}
