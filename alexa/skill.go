package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

const welcomeMessage = "Welcome to AWS Certified Developer Quiz. Say 'Start Quiz' to begin"
const exitMessage = "See you later"

func HandleRequest(ctx context.Context, r AlexaRequest) (AlexaResponse, error) {
	resp := CreateResponse()

	switch r.Request.Intent.Name {
	case "Begin":
		resp.Say("OK", true)
	case "AMAZON.HelpIntent":
		resp.Say(welcomeMessage, false)
	case "AMAZON.StopIntent":
		resp.Say(exitMessage, true)
	case "AMAZON.CancelIntent":
		resp.Say(exitMessage, true)
	default:
		resp.Say(welcomeMessage, false)
	}
	return *resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
