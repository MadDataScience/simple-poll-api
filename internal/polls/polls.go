package polls

import (
	database "github.com/maddatascience/simple-poll-api/internal/pkg/db/sqlite"
	"github.com/maddatascience/simple-poll-api/internal/users"
	"github.com/maddatascience/simple-poll-api/internal/questions"
	"log"
)

type Poll struct {
	PollID    string      `json:"pollID"`
	User      *users.User       `json:"user"`
	Title     string      `json:"title"`
	Questions []*questions.Question `json:"questions"`
}

func (poll Poll) Save() int64 {
	stmt, err := database.Db.Prepare("INSERT INTO Polls(Title, UserID) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(poll.Title, poll.User.UserID)
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

func GetAll() []Poll {
	stmt, err := database.Db.Prepare("select PollID, P.Title, UserID, U.Email from Polls P join Users U using (UserID)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var polls []Poll
	var email string
	var userID string
	for rows.Next() {
		var poll Poll
		err := rows.Scan(&poll.PollID, &poll.Title, &userID, &email)
		if err != nil{
			log.Fatal(err)
		}
		poll.User = &users.User{
			UserID: userID,
			Email: email,
		}
		polls = append(polls, poll)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return polls
}

func GetOne(pollID string) Poll {
	stmt, err := database.Db.Prepare("select PollID, P.Title, QuestionID, Question, QType from Polls P join Questions Q using (PollID) where PollID = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(pollID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var poll Poll
	for rows.Next() {
		var q questions.Question
		err := rows.Scan(&poll.PollID, &poll.Title, &q.QuestionID, &q.QuestionString, &q.QType)
		if err != nil {
			log.Fatal(err)
		}
		poll.Questions = append(poll.Questions, &q)
	}
	return poll
}