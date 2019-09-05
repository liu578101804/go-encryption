package email


type MailContentType string
const (
	MailContentTypeHtml = "Content-Type:text/plain; charset=UTF-8"
	MailContentTypeTxt = "Content-Type:text/html; charset=UTF-8"
)

type IEmail interface {
	SendEmail(to []string, subject, body string, contentType MailContentType) error
}