package errors

import (
	"errors"
	"github.com/rs/zerolog/log"
)

var ErrGenericError = errors.New("unexpected error")
var ErrNotFound = errors.New("not found")

func ManageErrorPanic(err error) {
	if err != nil {
		log.Fatal().Msg(err.Error())
		panic(err)
	}
}
