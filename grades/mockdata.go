package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Nick",
			LastName:  "Carter",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 94,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 82,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Jimmy",
			LastName:  "Wonder",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 83,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 96,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 75,
				},
			},
		},
		{
			ID:        3,
			FirstName: "Sam",
			LastName:  "Linkin",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 88,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 78,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 92,
				},
			},
		},
	}
}
