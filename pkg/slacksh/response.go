package slacksh

type response struct {
	Type string `json:"response_type"`
	Text string `json:"text"`
}
