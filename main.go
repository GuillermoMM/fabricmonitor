package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nraboy/alexa-slick-dealer/alexa"
)

// type FeedResponse struct {
// 	Channel struct {
// 		Item []struct {
// 			Title string `xml:"title"`
// 			Link  string `xml:"link"`
// 		} `xml:"item"`
// 	} `xml:"channel"`
// }

// func RequestFeed(mode string) (FeedResponse, error) {
// 	endpoint, _ := url.Parse("https://slickdeals.net/newsearch.php")
// 	queryParams := endpoint.Query()
// 	queryParams.Set("mode", mode)
// 	queryParams.Set("searcharea", "deals")
// 	queryParams.Set("searchin", "first")
// 	queryParams.Set("rss", "1")
// 	endpoint.RawQuery = queryParams.Encode()
// 	response, err := http.Get(endpoint.String())
// 	if err != nil {
// 		return FeedResponse{}, err
// 	} else {
// 		data, _ := ioutil.ReadAll(response.Body)
// 		var feedResponse FeedResponse
// 		xml.Unmarshal(data, &feedResponse)
// 		return feedResponse, nil
// 	}
// }

func IntentDispatcher(request alexa.Request) alexa.Response {
	var response alexa.Response
	if request.Body.Type == "LaunchRequest" {
		response = alexa.NewRepromptResponse("Hello, Welcome to the unnofficial app for provisioning the Nokia Fabric Service System. How can I help you today?", "For instructions on what you can ask, please say help me.")
	}
	switch request.Body.Intent.Name {
	case "CreateFabricIntent":
		response = HandleCreateFabricIntent(request)
	// case "PopularDealIntent":
	// 	response = HandlePopularDealIntent(request)
	case alexa.HelpIntent:
		response = HandleHelpIntent(request)
	case alexa.StopIntent:
		response = HandleStopIntent(request)
	case alexa.CancelIntent:
		response = HandleStopIntent(request)
	case "AboutIntent":
		response = HandleAboutIntent(request)
	case alexa.FallbackIntent:
		response = HandleHelpIntent(request)
	}
	return response
}

func HandleCreateFabricIntent(request alexa.Request) alexa.Response {
	var response alexa.Response
	if request.Body.Intent.Slots["fab_type"].Value == "sandbox" {
		response = alexa.NewRepromptResponse("I will create sandbox fabric with "+request.Body.Intent.Slots["settings_type"].Value+"settings", "a sandbox fabric I said")
	} else if request.Body.Intent.Slots["fab_type"].Value == "real" {
		response = alexa.NewRepromptResponse("I will create real fabric with"+request.Body.Intent.Slots["settings_type"].Value+"settings", "a real fabric I said")
	} else if request.Body.Intent.Slots["fab_type"].Value == "demo" {
		response = alexa.NewRepromptResponse("I will create a demo sandbox fabric with default settings", "a sandbox fabric I said")
	} else {
		response = alexa.NewRepromptResponse("Sorry, I don't know how to configure that yet", "I said that I don't know")
	}
	return response
}

// func HandlePopularDealIntent(request alexa.Request) alexa.Response {
// 	feedResponse, _ := RequestFeed("popdeals")
// 	var builder alexa.SSMLBuilder
// 	builder.Say("Here are the current popular deals:")
// 	builder.Pause("1000")
// 	for _, item := range feedResponse.Channel.Item {
// 		builder.Say(item.Title)
// 		builder.Pause("1000")
// 	}
// 	return alexa.NewSSMLResponse("Popular Deals", builder.Build())
// }

func HandleHelpIntent(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder
	builder.Say("Here are some of the things you can ask:")
	builder.Pause("1000")
	builder.Say("Create a sandbox fabric with default settings.")
	builder.Pause("1000")
	builder.Say("Create a real fabric with default settings.")
	builder.Pause("1000")
	builder.Say("Create a sandobox fabric with customize settings.")
	builder.Pause("1000")
	builder.Say("Create a real fabric with customize settings.")
	return alexa.NewSSMLResponse("Slick Dealer Help", builder.Build())
}

func HandleAboutIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponseWithCard("About", "Fabric Builder was created by Guillermo Murguia in Antwerp, Belgium as an unofficial fabric builder application.")
}

func HandleStopIntent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Cancel and Stop", "Goodbye")
}

func Handler(request alexa.Request) (alexa.Response, error) {
	return IntentDispatcher(request), nil
}

func main() {
	lambda.Start(Handler)
}
