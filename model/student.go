package model

type Student struct {
	ID string `json:"id"`
	StudentCode string `json:"student_code"`
	StudentName string `json:"student_name"`
	StudentAddress string `json:"student_address"`
	StudentOccupation string `json:"student_occupation"`
	JoiningReason string `json:"joining_reason"`
}

type Students struct {
	Students []Student `json:"participants"`
}