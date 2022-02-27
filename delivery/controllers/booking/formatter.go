package booking

import (
	"be/entities"
)

type CreateBookingRequesFormat struct {
	User_uid      string
	Room_uid      string `json:"room_uid"`
	Start_date    string `json:"start_date"`
	End_date      string `json:"end_date"`
	PaymentMethod string `json:"paymentmethod"`
	Status        string `json:"status"`
}

type GetBookingResponseFormat struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    entities.Booking `json:"data"`
}

//

type PaymentResponse struct {
	OrderID     string   `json:"order_id"`
	GrossAmount string   `json:"gross_amount"`
	PaymentType string   `json:"payment_type"`
	Url         string `json:"url"`
}

type PaymentTypeRequest struct {
	Payment_method string `json:"payment_method" validate:"required"`
}

type RequestCallBackMidtrans struct {
	Transaction_time   string `json:"transaction_time"`
	Transaction_status string `json:"transaction_status"`
	Transaction_id     string `json:"transaction_id"`
	Status_message     string `json:"status_message"`
	Status_code        string `json:"status_code"`
	Signature_key      string `json:"signature_key"`
	Settlement_time    string `json:"settlement_time"`
	Payment_type       string `json:"payment_type"`
	Order_id           string `json:"order_id"`
	Merchant_id        string `json:"merchant_id"`
	Gross_amount       string `json:"gross_amount"`
	Fraud_status       string `json:"fraud_status"`
	Currency           string `json:"currency"`
}
