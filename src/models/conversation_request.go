package models

type Role int

const(
	User Role = iota
	Assistant
	System
)


type ConversationRequest struct{
	Message []ConversationMessage `json:"message"`
}

type ConversationMessage struct{
	Role Role `json:"role"`
	Content string `json:"content"`
}