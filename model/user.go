package model

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
	Sex   string `json:"sex"`
}

func ListStudent() []*Student {
	return []*Student{
		{
			Id:    1,
			Name:  "Student01",
			Class: "Developer01",
			Sex:   "Male",
		},
		{
			Id:    2,
			Name:  "Student02",
			Class: "Tester01",
			Sex:   "Female",
		},
		{
			Id:    3,
			Name:  "Student03",
			Class: "ProjectManager01",
			Sex:   "Male",
		},
	}
}
