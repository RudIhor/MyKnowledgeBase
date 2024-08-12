package enum

type QuestionStatus bool

var (
	QuestionStatusAnswered QuestionStatus = true
	QuestionStatusOpen     QuestionStatus = false
)

func (s QuestionStatus) Bool() bool {
	return bool(s)
}

func (s QuestionStatus) DisplayName() string {
	var name string

	switch s {
	case QuestionStatusAnswered:
		name = "Answered"
	case QuestionStatusOpen:
		name = "Open"
	}

	return name
}
