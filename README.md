### example queries

```graphql
query {
	polls{
    title
    questions{
      question
      qType
      answers{
        answerB
      }
    }
    user{
      email
    }
  }
}
```

```graphql
mutation {
  createPoll(input: {title: "What the?"}){
    title,
    pollID
  }
}
```

```graphql
mutation {
  createUser(input: {email: "paul@flourishtech.us", password: "123"})
}
```

```
{
  "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDczMDA4MzYsInVzZXJuYW1lIjoidXNlcjEifQ.dKR_Qc5hnP7fhjZNbT8zT8lyTJW5T-lx7S0wEDkgFGc"
}
```

```graphql
query {
  polls{
    pollID
    title
  }
}
```

```graphql
mutation {
  createQuestion(input: {pollID: "1", question: "Are you not?", qType: "Boolean"}){
    title,
    pollID,
    questions{
     questionID,
     question
    }
  }
}
```

---

Not working:

`migrate -database sqlite3:///poll.db -path internal/pkg/db/migrations/sqlite up`

> error: database driver: unknown driver sqlite3 (forgotten import?)
