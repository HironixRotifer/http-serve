package models

type Intents struct {
	Service        string `json:"service"`
	ServiceItem    string `json:"service_item"`
	Country        string `json:"country"`
	Operator       string `json:"operator"`
	ProjectID      string `json:"project_id"`
	GateID         string `json:"gate_id"`
	GateSenderID   string `json:"gate_sender_id"`
	ClientSenderID string `json:"client_sender_id"`
	MessageID      string `json:"message_id"`
	PhoneNumber    string `json:"phone_number"`
	DeliveryTTL    int32  `json:"delivery_ttl"`
}
