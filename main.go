package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("students.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	students := []Student{}
	for sc.Scan() {
		line := strings.Split(sc.Text(), " ")
		name := line[0]
		grade, _ := strconv.Atoi(line[1])

		student := findStudent(name, students)
		if student == nil {
			newStudent := Student{name: name, grades: []int{grade}}
			students = append(students, newStudent)
		} else {
			student.addGrade(grade)
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	for _, student := range students {
		sum := 0
		for _, grade := range student.grades {
			sum += grade
		}

		fmt.Println(student.name+"\n"+"Scores: ", student.gradesToStr(), "\nAverage score:", float64(sum)/float64(len(student.grades)), "\n")
	}
}
