package main

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/levigross/grequests"
	"log"
)

var (
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

func Handler(request events.APIGatewayProxyRequest) (string, error) {
	log.Printf("Process Lambda request %s\n", request.RequestContext.RequestID)
	tokenJSON := token{}
	getToken(&tokenJSON)
	return tokenJSON.Update.NA.Formatted.Buy, nil
}

type token struct {
	Update struct {
		NA struct {
			Formatted struct {
				Buy string `json:"buy"`
			} `json:"formatted"`
		} `json:"NA"`
	} `json:"update"`
}

func getToken(token *token) *token {
	resp, err := grequests.Get("https://data.wowtoken.info/wowtoken.json", nil)
	if err != nil {
		log.Fatalln("Unable to make request", err)
	}

	jsonErr := json.Unmarshal(resp.Bytes(), &token)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return token
}

func main() {
	lambda.Start(Handler)
}
