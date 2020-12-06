package questions

import (
	"github.com/maddatascience/simple-poll-api/internal/answers"
	database "github.com/maddatascience/simple-poll-api/internal/pkg/db/sqlite"
	"log"
)
type Question struct {
	QuestionID string    `json:"questionID"`
	PollID string `json:"pollID"`
	QuestionString string    `json:"question"`
	QType    string    `json:"qType"`
	Answers  []*answers.Answer `json:"answers"`
}


func (question Question) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Questions(PollID, Question, QType) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(question.PollID, question.QuestionString, question.QType)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")
	return id
}