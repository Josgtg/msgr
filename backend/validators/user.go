package validators

import (
	"errors"
	"regexp"
)

func Name(name string) bool {
	return len(name) <= 64
}

func Password(password string) error {
	warnings := "password has the next flaws:\n"
	if len(password) < 8 {
		warnings += "is too short, must have a minimum of 8 characters\n"
	}
	if match, _ := regexp.MatchString(`[0-9]`, password); !match {
		warnings += "must contain a digit\n"
	}
	if match, _ := regexp.MatchString(`\s`, password); match {
		warnings += "must not contain whitespaces\n"
	}
	if match, _ := regexp.MatchString(`(?!\s)\W`, password); !match {
		warnings += "must contain at least one symbol\n"
	}
	if warnings == "password has the next flaws:\n" {
		return errors.New(warnings)
	}
	return nil
}
