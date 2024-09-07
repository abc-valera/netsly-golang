package service

type IEmailer interface {
	SendEmail(e EmailSendRequest) error
}

type EmailSendRequest struct {
	Subject     string
	Content     string
	To          []string
	AttachFiles []string
}
