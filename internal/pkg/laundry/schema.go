package laundry

type TransactionRequest struct {
	ServiceID     uint   `json:"service_id"`
	DurationID    uint   `json:"duration_id"`
	CustomerName  string `json:"customer_name"`
	CustomerPhone string `json:"customer_phone"`
}
