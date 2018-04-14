package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
)

const welcomeMessage = "Welcome to AWS Certified Developer Quiz. Say 'Start Quiz' to begin"
const exitMessage = "See you later"

var services = map[string]string{
	"A": "EC2",
	"B": "VPC",
	"C": "DynamoDB",
	"D": "S3",
	"E": "SQS",
}

func HandleRequest(ctx context.Context, r AlexaRequest) (AlexaResponse, error) {
	resp := CreateResponse()

	switch r.Request.Intent.Name {
	case "Begin":
		resp.Say(`<speak>
			Choose the AWS service you want to be tested on <break time="1s"/>
			A <break time="1s"/> EC2 <break time="1s"/>
			B <break time="1s"/> VPC <break time="1s"/>
			C <break time="1s"/> DynamoDB <break time="1s"/>
			D <break time="1s"/> S3 <break time="1s"/>
			E <break time="1s"/> SQS
		</speak>`, false, "SSML")
	case "ServiceChoice":
		number := strings.TrimSuffix(r.Request.Intent.Slots["choice"].Value, ".")

		questions, _ := getQuestions(services[number])

		resp.SessionAttributes = make(map[string]interface{}, 1)
		resp.SessionAttributes["service"] = services[number]
		resp.SessionAttributes["current"] = 0
		resp.SessionAttributes["score"] = 0
		resp.SessionAttributes["correct"] = questions[0].Correct

		log.Println(questions)

		text := fmt.Sprintf(`<speak>You chose %s. First question, %s ? <break time="1s"/>`, questions[0].Category, questions[0].Question)
		for choice, answer := range questions[0].Answers {
			text += fmt.Sprintf(`%s <break time="1s"/> %s <break time="1s"/>`, choice, answer)
		}
		text += "</speak>"

		resp.Say(text, false, "SSML")
	case "AnswerChoice":
		resp.SessionAttributes = make(map[string]interface{}, 1)
		score := int(r.Session.Attributes["score"].(float64))
		correctChoice := r.Session.Attributes["correct"]
		userChoice := strings.TrimSuffix(r.Request.Intent.Slots["choice"].Value, ".")
		currentIndex := int(r.Session.Attributes["current"].(float64))
		service := r.Session.Attributes["service"].(string)
		questions, _ := getQuestions(service)

		log.Println(userChoice)
		log.Println(correctChoice)
		log.Println(currentIndex)
		log.Println(questions)
		log.Println(score)

		text := "<speak>"
		if correctChoice == userChoice {
			text += "Correct,"
			score++
		} else {
			text += "Incorrect,"
		}

		if currentIndex == 4 {
			text += fmt.Sprintf("You answered %d of 5 questions</speak>", score)
			resp.Say(text, true, "SSML")
		} else {
			text += fmt.Sprintf(`Next question, %s ? <break time="1s"/>`, questions[currentIndex+1].Question)
			for choice, answer := range questions[currentIndex+1].Answers {
				text += fmt.Sprintf(`%s <break time="1s"/> %s <break time="1s"/>`, choice, answer)
			}
			text += "</speak>"

			resp.SessionAttributes["score"] = score
			resp.SessionAttributes["service"] = service
			resp.SessionAttributes["current"] = currentIndex + 1
			resp.SessionAttributes["correct"] = questions[currentIndex+1].Correct

			resp.Say(text, false, "SSML")
		}

	case "AMAZON.HelpIntent":
		resp.Say(welcomeMessage, false, "PlainText")
	case "AMAZON.StopIntent":
		resp.Say(exitMessage, true, "PlainText")
	case "AMAZON.CancelIntent":
		resp.Say(exitMessage, true, "PlainText")
	default:
		resp.Say(welcomeMessage, false, "PlainText")
	}
	return *resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
