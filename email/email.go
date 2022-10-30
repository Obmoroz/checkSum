package email

var MainMailer *Mailer

type Transport interface {
	Send(Subject, Body string) error
}

type Mailer struct {
	Transport
	subject string
}

func NewMailer(transport Transport, subject string) *Mailer {
	return &Mailer{Transport: transport, subject: subject}
}

func (m Mailer) Send(Body string) error {
	err := m.Transport.Send(m.subject, Body)
	if err != nil {
		return err
	}
	return nil
}
