package rabbitmq

type Message struct {
	Title      string  `json:"title"`
	Body       string  `json:"body"`
	Image      []byte  `json:"img"`
	Url        *string `json:"url"`
	FilterWord string  `json:"filter_word"`
}
