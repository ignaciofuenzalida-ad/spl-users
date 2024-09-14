package model

type QueueUsersStatistics struct {
	General  Details `json:"general"`
	LastHour Details `json:"lastHour"`
	Waiting  int     `json:"waiting"`
	Total    int     `json:"total"`
}

type Details struct {
	Found     int `json:"found"`
	NotFound  int `json:"notFound"`
	WithError int `json:"withError"`
}
