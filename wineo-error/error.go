package wineoerr

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type Error struct {
	Code   int
	Params map[string]interface{}
}

func (err *Error) Error() string {
	switch {
	case err.Code == InvalidJSONInRequestBody:
		return fmt.Sprintf("invalid json in request body: %s", err.Params["explanation"])
	default:
		return fmt.Sprintf("unknown error with code '%d' and params:\n%s", err.Code, spew.Sdump(err.Params))
	}
}
