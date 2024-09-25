package question

import (
	b64 "encoding/base64"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/stclaird/questionGenerator/pkg/models"
)

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func generateQuestionFileName(q models.QuestionOut) (name string ){
    //Generate File name for Question
    slug := generateQuestionID(q)[:7]

    cat := strings.ToLower(q.Category)
    catsafe := safeFileName(cat)

    subcat := strings.ToLower(q.Subcategory)
    subcatsafe := safeFileName(subcat)

    name = fmt.Sprintf("%s-%s-%s.json", catsafe, subcatsafe, slug )

    return name
}

func generateQuestionID(q models.QuestionOut) string {
    s := q.QuestionText
    se := b64.StdEncoding.EncodeToString([]byte(s))
    return se
}

func safeFileName(s string) string {
    //Remove any undesirable characters from a file name
    re := regexp.MustCompile(`[\\/:*?"<>|]`)
    cleanString := re.ReplaceAllString(s, "")

    return cleanString
}

func generatePrompt(questionIn models.QuestionIn) string {
    // create the prompt to make use of AI's JSONs RESPONSE
    var answersStr []string

    //How many answers do we want in our json template
    numAns := questionIn.NumCorrect + questionIn.NumInCorrect

    //We can't have 0 so if it is 0 set it to be 2 as this gives us 1 correct 1 incorrect.
    if numAns == 0 {
        numAns = 2
    }

    for i := 0; i < numAns; i++ {
        answersStr = append(answersStr, "{'text': string, 'iscorrect':bool}" )
    }

    promptPrefix := "Ask me a question regarding"
    promptAnswers := fmt.Sprintf("give me %v correct answers and %v incorrect answers using this JSON schema:", questionIn.NumCorrect, questionIn.NumInCorrect)
    promptJson := fmt.Sprintf("Question =  {'question': string, 'answerReference': string, 'answers':[ %s ]}", answersStr)
    promptSuffix := "Return: <Question>"


    fullQuestion := fmt.Sprintf("%s %s %s %s %s", promptPrefix, questionIn.QuestionText, promptAnswers, promptJson, promptSuffix )

    return fullQuestion
}
