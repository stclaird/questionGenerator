package question

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/stclaird/questionGenerator/pkg/models"
	"google.golang.org/api/option"
)

func generateQuestion(ctx *gin.Context) {

    var questionDir = "questions"

    var questionIn models.QuestionIn
    if err := ctx.BindJSON(&questionIn); err != nil {
        ctx.JSON(400, gin.H{"error": err.Error()})
        return
    }

    questionOut := AskAi(questionIn)

	questionJsonBytes, _ := json.Marshal(questionOut)
    questionName := generateQuestionFileName(questionOut)
    err := os.MkdirAll(questionDir, 0755)

    if err != nil {
        panic(err)
    }

    var writeFileName = fmt.Sprintf("%s/%s", questionDir, questionName)
    os.WriteFile(writeFileName, questionJsonBytes, os.ModePerm)

    ctx.JSON(http.StatusOK, questionOut)
}

func AskAi (questionIn models.QuestionIn) models.QuestionOut {

	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	model := client.GenerativeModel("gemini-1.5-pro-latest")
	// Ask the model to respond with JSON.
	model.ResponseMIMEType = "application/json"
	prompt :=  generatePrompt(questionIn)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
	}

	var bytes []byte
	var question models.QuestionOut
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			bytes = []byte(txt)
		}
		fmt.Println(string(bytes))
	}

	json.Unmarshal(bytes, &question)

	return question

}
