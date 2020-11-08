package temp

type QuizUpdate struct {
	NumberOfQuestion int
}


type QuizReceived struct {
	Creator string
	Name string
	NumberOfQuestion string
}

type QuestionUpdate struct {
	Question string
	Choice1  string
	Choice2  string
	Choice3  string
	Choice4  string
	Answer   int
}

type mapQuestion struct {
	num map[string]QuestionUpdate
}
