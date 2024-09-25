package models

type QuestionIn struct {
	QuestionText        string `json:"question" form:"text" binding:"required"`
	NumInCorrect int `json:"numincorrect" form:"int"`
	NumCorrect int `json:"numcorrect" form:"int"`
}

type QuestionOut struct {
	QuestionText        string `json:"question" form:"text" binding:"required"`
	Type        string `json:"type" form:"type"`
	Category    string `json:"category" form:"category"`
	Subcategory string `json:"subcategory" form:"subcategory"`
	DateAdded   string `json:"dateadded" form:"dateadded"`
	Certification string `json:"certification" form:"certification"`
	AnswerReference string `json:"answerReference" form:"answersupport"`
	Answers     []Answer `json:"answers" form:"answers" binding:""`
}

type Answer struct {
	Text string `json:"text"`
	IsCorrect bool   `json:"iscorrect"`
}
