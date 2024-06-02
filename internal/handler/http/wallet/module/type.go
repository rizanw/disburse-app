package module

type DisburseRequest struct {
	UserID int64   `json:"user_id"`
	Amount float64 `json:"amount"`
}

type DisburseResponse struct {
	UserID        int64   `json:"user_id"`
	RequestAmount float64 `json:"request_amount"`
	Balance       float64 `json:"balance"`
	Status        string  `json:"status"`
}
