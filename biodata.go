package main

import (
	"assignment-01/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func cleanedText(text string) string {
	return strings.ToLower(strings.Replace(text, "-", "", -1))
}

func isStudentExist(userInput, studentCode string) bool {
	userinputCleaned := cleanedText(userInput)
	studentCodeCleaned := cleanedText(studentCode)

	return userinputCleaned == studentCodeCleaned
}

func retrieveAndShowStudentInfo(userInput string) {
	jsonPath := filepath.Join("data", "data.json")

	data, err := ioutil.ReadFile(jsonPath)

	if err != nil {
		log.Fatal(err)
	}

	var students model.Students

	if err := json.Unmarshal(data, &students); err != nil {
		log.Fatal(err)
	}

	if userInput == "" {
		fmt.Println("Please input the student code to retrieve student information!")
		return
	}

	for _, student := range students.Students {
		if isStudentExist(userInput, student.StudentCode) {
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Student Code: %s\n", student.StudentCode)
			fmt.Printf("Name: %s\n", student.StudentName)
			fmt.Printf("Address: %s\n", student.StudentAddress)
			fmt.Printf("Occupation: %s\n", student.StudentOccupation)
			fmt.Printf("Joining Reason: %s\n", student.JoiningReason)
			return
		}
	}

	fmt.Printf("The student with student code '%s' was not found. Please enter a valid code! \n", userInput)
}

func getUserInput() string {
	args := os.Args

	if len(args) < 2 {
		return ""
	}

	return args[1]
}

func main() {
	userInput := getUserInput()

	retrieveAndShowStudentInfo(userInput)
}