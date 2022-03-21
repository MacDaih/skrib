package skrib

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

type payload struct {
	Timestamp time.Time              `json:"timestamp"`
	Msg       string                 `json:"message"`
	Fields    map[string]interface{} `json:"payload"`
}

type errorfield struct {
	head string
	err  error
}

type infofield struct {
	head string
	info string
}

func InfoField(head string, info string) infofield {
	return infofield{
		head: head,
		info: info,
	}
}

func Err(head string, err error) errorfield {
	return errorfield{
		head: head,
		err:  err,
	}
}

type Skrib struct {
	logLevel Level
	out      io.Writer
}

func NewSkrib(w io.Writer, lvl Level) Skrib {
	return Skrib{
		logLevel: lvl,
		out:      w,
	}
}

func (s *Skrib) Log(level Level, msg string, args ...interface{}) {
	if level < s.logLevel {
		return
	}
	m := payload{
		Timestamp: time.Now().UTC(),
		Msg:       msg,
		Fields:    make(map[string]interface{}),
	}

	for _, v := range args {
		switch t := v.(type) {
		case infofield:
			m.Fields[t.head] = t.info
		case errorfield:
			m.Fields[t.head] = t.err.Error()
		default:
			m.Fields[level.String()] = v
		}
	}

	json.NewEncoder(s.out).Encode(m)
	if level == FATAL {
		os.Exit(1)
	}
}
