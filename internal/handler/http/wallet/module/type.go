package module

type DisburseRequest struct {
	UserID int64   `json:"user_id"`
	Amount float64 `json:"amount"`
}
