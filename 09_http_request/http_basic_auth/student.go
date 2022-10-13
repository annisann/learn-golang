package main

var students = []*Student{}

type Student struct {
	ID    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.ID == id {
			return each
		}
	}
	return nil
}

func init() {
	students = append(students, &Student{ID: "S001", Name: "Haley Dunphy", Grade: 6})
	students = append(students, &Student{ID: "S002", Name: "Manny Delgado", Grade: 3})
	students = append(students, &Student{ID: "S003", Name: "Alexander Dunphy", Grade: 4})
	students = append(students, &Student{ID: "S004", Name: "Lily Tucker-Pritchett", Grade: 1})
	students = append(students, &Student{ID: "S005", Name: "Luke Dunphy", Grade: 3})
}
