package pb

import (
	fmt "fmt"
	"regexp"
)

var (
	phone = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	email = regexp.MustCompile(`^.*@.*$`)
)

// Validate ensures that the contact identifier is a valid E.164 phone number or an
// email address and returns an error if it is not. Note that this is a fairly
// lightweight approach to validation and is prone to false positives. Error handling
// on send or receive messages is essential.
//
// See also: https://www.twilio.com/docs/glossary/what-e164
// See also: https://tools.ietf.org/rfc/rfc5322.txt
func (c *Contact) Validate() (err error) {
	switch c.Type {
	case ContactType_PHONE:
		if !phone.MatchString(c.Contact) {
			return fmt.Errorf("could validate phone number '%s'", c.Contact)
		}
	case ContactType_EMAIL:
		if !email.MatchString(c.Contact) {
			return fmt.Errorf("could not validate email '%s'", c.Contact)
		}
	default:
		return fmt.Errorf("unknown contact type %s", c.Type)
	}
	return nil
}
