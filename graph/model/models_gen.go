// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Answer struct {
	AnswerID   string `json:"answerID"`
	IP         string `json:"ip"`
	ATimestamp string `json:"aTimestamp"`
	AnswerB    bool   `json:"answerB"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewAnswer struct {
	AnswerB bool `json:"answerB"`
}

type NewAnswers struct {
	Answers []*NewAnswer `json:"answers"`
}

type NewPoll struct {
	Title string `json:"title"`
}

type NewQuestion struct {
	PollID   string `json:"pollID"`
	Question string `json:"question"`
	QType    string `json:"qType"`
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Poll struct {
	PollID    string      `json:"pollID"`
	User      *User       `json:"user"`
	Title     string      `json:"title"`
	Questions []*Question `json:"questions"`
}

type Question struct {
	QuestionID string    `json:"questionID"`
	Question   string    `json:"question"`
	QType      string    `json:"qType"`
	Answers    []*Answer `json:"answers"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type User struct {
	UserID string `json:"userID"`
	Email  string `json:"email"`
}
