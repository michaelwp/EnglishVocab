package errors

import (
	"errors"
	"log"
	"strings"
)

func New(message ...string) error {
	messages := strings.Join(message, " ")
	log.Println(messages)
	return errors.New(messages)
}
