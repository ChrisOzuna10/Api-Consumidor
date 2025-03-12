package entities

type Notification struct {
	ID      int    `json:"id"`
	Music   string `json:"music"`
	Message string `json:"message"`
}

func NewNotification(id int, music string) *Notification {
	return &Notification{
		ID:      id,
		Music:   music,
		Message: "La música se ha registrado: " + music,
	}
}
