package answers

type Answer struct {
	AnswerID   string `json:"answerID"`
	IP         string `json:"ip"`
	ATimestamp string `json:"aTimestamp"`
	AnswerB    bool   `json:"answerB"`
}
