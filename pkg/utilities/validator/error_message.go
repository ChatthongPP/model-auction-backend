package validator

import (
	"net/http"
)

const (
	ERR_VALIDATION_FAILED = `VALIDATION_FAILED`
)

var ErrorMessage = map[string]string{
	`VALIDATION_FAILED`: `Failed for validate the input`,
}

var ErrorStatusCode = map[string]int{
	`VALIDATION_FAILED`: http.StatusNotAcceptable,
}
