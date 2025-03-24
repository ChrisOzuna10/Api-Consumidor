package entities

type Notification struct {
	ID      int    `json:"id"`
	Music   string `json:"music"`
	Message string `json:"message"`
}
