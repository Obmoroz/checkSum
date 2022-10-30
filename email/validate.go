package email

import "regexp"

const (
	minEmailLen = 3
	maxEmailLen = 255
)

var emailMask = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func IsValid(email []string) string {
	for i := range email {
		l := len(email[i])
		if l < minEmailLen || l > maxEmailLen {
			return email[i]
		}
		if !emailMask.MatchString(email[i]) {
			return email[i]
		}
	}
	return ``
}
