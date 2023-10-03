package service

type EmailSender interface {
	SendEmail(e Email) error
}

type Email struct {
	Subject     string
	Content     string
	To          []string
	AttachFiles []string
}
