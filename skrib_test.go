package skrib

import (
	"encoding/json"
	"io"
	"reflect"
	"testing"
	"time"
)

var (
	message = "test_message"
	fields  = map[string]interface{}{
		"field_one": "field one as string",
		"field_two": 0.10,
	}
)

func TestBaseCase(t *testing.T) {
	r, w := io.Pipe()
	s := NewSkrib(w, 0)

	p := payload{
		Timestamp: time.Now().UTC(),
		Msg:       message,
		Fields:    fields,
	}
	go func(s Skrib) {
		for {
			s.Log(INFO, message, fields)
			time.Sleep(1 * time.Second)
		}
	}(s)

	var p2 payload
	err := json.NewDecoder(r).Decode(&p2)

	if err != nil {
		t.Error(err)
	}

	if p.Msg != p2.Msg && reflect.DeepEqual(p.Fields, p2.Fields) {
		t.Errorf("expect %s but got %s", p.Msg, p2.Msg)
	}
}
