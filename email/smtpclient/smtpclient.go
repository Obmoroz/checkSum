package smtpclient

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SMTP struct {
	host string
	port int
	to   []string
	from string
	pass string
}

func NewSMTP(host string, port int, from, pass string) (*SMTP, error) {

	return &SMTP{
		host: host,
		port: port,
		from: from,
		pass: pass,
	}, nil
}

func (s *SMTP) Send(Mail, Subject, Body string) error {
	var (
		from    = mail.Address{Name: `Citis`, Address: s.from}
		to      = mail.Address{Address: Mail}
		msg     bytes.Buffer
		headers = map[string]string{
			`From`:                      from.String(),
			`To`:                        to.String(),
			`Subject`:                   Subject,
			`Content-Type`:              `text/plain; charset="utf-8"`,
			`Content-Transfer-Encoding`: `base64`,
			`Content-Language`:          `ru`,
			`MIME-Version`:              `1.0`,
		}
	)

	var keys []string
	for k := range headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Fprintf(&msg, "%s: %s\r\n", key, headers[key])
	}
	msg.WriteString("\r\n")
	msg.WriteString(base64.StdEncoding.EncodeToString([]byte(Body)))
	msg.WriteString("\r\n")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	return s.mailProcess(ctx, from.Address, Mail, msg.Bytes())
}

func (s *SMTP) mailProcess(ctx context.Context, from string, Mail string, body []byte) error {
	host := s.host + `:` + strconv.FormatInt(int64(s.port), 10)

	// Вместо smtp.Dial()
	conn, err := net.DialTimeout("tcp", host, 7*time.Second)
	if err != nil {
		return fmt.Errorf("net.DialTimeout() failed: %w", err)
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("net.DialTimeout() failed: %w", err)
	}

	if err = c.Mail(from); err != nil {
		return fmt.Errorf("c.Mail() failed: %w", err)
	}
	if err = c.Rcpt(Mail); err != nil {
		return fmt.Errorf("smtp.Rcpt(%s) failed: %w", Mail, err)
	}
	w, err := c.Data()
	if err != nil {
		return fmt.Errorf("c.Data() failed: %w", err)
	}
	_, err = w.Write(body)
	if err != nil {
		return fmt.Errorf("w.Write() failed: %w", err)
	}
	err = w.Close()
	if err != nil {
		return fmt.Errorf("w.Close() failed: %w", err)
	}

	err = c.Quit()
	if err != nil {

		return fmt.Errorf("c.Quit() failed: %w", err)
	}

	return nil
}

func encodeStr(str []byte) string {
	if len(str) > 0 {
		return "=?UTF-8?B?" + base64.StdEncoding.EncodeToString(str) + "?="
	} else {
		return ``
	}
}

func makeCc(m []string) string {
	var out []string = make([]string, 0, len(m))
	for i := 0; i < len(m); i++ {
		s := mail.Address{Address: m[i]}
		out = append(out, s.String())
	}
	return strings.Join(out, `, `)
}
