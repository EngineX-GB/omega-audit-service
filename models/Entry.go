package models

type Entry struct {
	Source         string `json:"source"`
	SourceUrl      string `json:"sourceUrl"`
	ResourceUrl    string `json:"resourceUrl"`
	SystemFilePath string `json:"systemFilePath"`
	Strategy       string `json:"strategy"`
	Invoker        string `json:"invoker"`
	Channel        string `json:"channel"`
	Title          string `json:"title"`
	Timestamp      string `json:"timestamp"`
}
