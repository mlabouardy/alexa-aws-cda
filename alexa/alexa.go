package main

type AlexaRequest struct {
	Version string `json:"version"`
	Request struct {
		Type   string `json:"type"`
		Time   string `json:"timestamp"`
		Intent struct {
			Name  string `json:"name"`
			Slots map[string]struct {
				Value string `json:"value"`
			} `json:"slots"`
		} `json:"intent"`
	} `json:"request"`
	Session struct {
		Attributes map[string]interface{} `json:"attributes"`
	}
}

type AlexaResponse struct {
	Version  string `json:"version"`
	Response struct {
		ShouldEndSession bool `json:"shouldEndSession"`
		OutputSpeech     struct {
			Type string `json:"type"`
			Text string `json:"text"`
			SSML string `json:"ssml"`
		} `json:"outputSpeech"`
	} `json:"response"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes"`
}

func CreateResponse() *AlexaResponse {
	var resp AlexaResponse
	resp.Version = "1.0"
	return &resp
}

func (resp *AlexaResponse) Say(text string, shouldEndSession bool, speechType string) {
	resp.Response.ShouldEndSession = shouldEndSession
	resp.Response.OutputSpeech.Type = speechType

	if speechType == "SSML" {
		resp.Response.OutputSpeech.SSML = text
	} else {
		resp.Response.OutputSpeech.Text = text
	}
}
