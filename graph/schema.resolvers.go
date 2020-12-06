package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/maddatascience/simple-poll-api/graph/generated"
	"github.com/maddatascience/simple-poll-api/graph/model"
	"github.com/maddatascience/simple-poll-api/internal/auth"
	"github.com/maddatascience/simple-poll-api/internal/polls"
	"github.com/maddatascience/simple-poll-api/internal/questions"
	"github.com/maddatascience/simple-poll-api/internal/users"
	"github.com/maddatascience/simple-poll-api/pkg/jwt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Email = input.Email
	user.Password = input.Password
	user.Create()
	token, err := jwt.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Email = input.Email
	user.Password = input.Password
	correct := user.Authenticate()
	if !correct {
		// 1
		return "", &users.WrongEmailOrPasswordError{}
	}
	token, err := jwt.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	email, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	token, err := jwt.GenerateToken(email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) CreatePoll(ctx context.Context, input model.NewPoll) (*model.Poll, error) {
	var poll polls.Poll
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Poll{}, fmt.Errorf("access denied")
	}
	poll.Title = input.Title
	poll.User = user
	pollID := poll.Save()
	grahpqlUser := &model.User{
		UserID: user.UserID,
		Email:  user.Email,
	}
	return &model.Poll{PollID: strconv.FormatInt(pollID, 10), Title: poll.Title, User: grahpqlUser}, nil
}

func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.NewQuestion) (*model.Poll, error) {
	var q questions.Question
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Poll{}, fmt.Errorf("access denied")
	}
	q.PollID = input.PollID
	q.QType = input.QType
	q.QuestionString = input.Question
	qID := q.Save()
	dbPoll := polls.GetOne(q.PollID)
	resultPoll := model.Poll{
		PollID: dbPoll.PollID,
		Title: dbPoll.Title,
	}
	okay := false
	for _, pollQuestion := range dbPoll.Questions {
		resultQuestion := model.Question{
			QuestionID: pollQuestion.QuestionID,
			QType: pollQuestion.QType,
			Question: pollQuestion.QuestionString,
		}
		if resultQuestion.QuestionID == strconv.FormatInt(qID, 10) {
			okay = true
		}
		resultPoll.Questions = append(resultPoll.Questions, &resultQuestion)
	}
	var err error
	if okay != true {
		err = fmt.Errorf("Question Failed to insert")
	}
	return &resultPoll, err
}

func (r *mutationResolver) CreateAnswer(ctx context.Context, input model.NewAnswers) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Polls(ctx context.Context) ([]*model.Poll, error) {
	var resultPolls []*model.Poll
	var dbPolls []polls.Poll
	dbPolls = polls.GetAll()
	for _, poll := range dbPolls {
		grahpqlUser := &model.User{
			UserID: poll.User.UserID,
			Email:  poll.User.Email,
		}
		resultPolls = append(resultPolls, &model.Poll{PollID: poll.PollID, Title: poll.Title, User: grahpqlUser})
	}
	return resultPolls, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
