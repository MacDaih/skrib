package skrib

import (
	"bytes"
	"errors"
	"log"
	"os"
	"testing"
)

type exiter func(code int)

func TestHandleError(t *testing.T) {
	e := errors.New("Some error occured")
	HandleError(e, false)
}

func TestHandleLog(t *testing.T) {
	msg := map[int]string{
		0: "Some kind of SUCCESS",
		1: "Some kind of WARNING",
		3: "Some kind of INFO",
	}
	for k, v := range msg {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stdout)
		}()
		HandleLog(v, k)
		if buf.Bytes() == nil {
			t.Fail()
		}
	}
}
