package main

import (
	"strconv"
	"strings"
)

type Student struct {
	name   string
	grades []int
}

func (s *Student) addGrade(grade int) {
	s.grades = append(s.grades, grade)
}
func (s *Student) setName(name string) {
	s.name = name
}
func (s Student) gradesToStr() string {
	strGrades := make([]string, len(s.grades))
	for i, grade := range s.grades {
		strGrades[i] = strconv.Itoa(grade)
	}
	return strings.Join(strGrades, ", ")
}
func findStudent(name string, students []Student) *Student {
	for i := range students {
		if students[i].name == name {
			return &students[i]
		}
	}
	return nil
}
