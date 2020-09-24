package temp

type QuizUpdate struct {
	NumberOfQuestion int
}

type HostUpdate struct {
	Name       			string
	NumberOfParticipant int
}

type QuestionUpdate struct {
	Question string
	Choice1 string
	Choice2 string
	Choice3 string
	Choice4 string
	answer int
}
