package dto

type TransferReqBody struct {
	ReceiverId string  `json:"receiver_id"`
	Amount     float64 `json:"amount"`
}
