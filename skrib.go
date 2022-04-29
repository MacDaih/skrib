package skrib

import (
	"encoding/json"
	"io"
	"os"
	"time"
)

type payload struct {
	Lvl       string                 `json:"level"`
	Timestamp time.Time              `json:"timestamp"`
	Msg       string                 `json:"message"`
	Fields    map[string]interface{} `json:"payload"`
}

type logdata map[string]interface{}

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

func (s *Skrib) SetLevel(lvl Level) {
	s.logLevel = lvl
}

func (s *Skrib) Log(level Level, msg string, args logdata) {
	if level < s.logLevel {
		return
	}
	m := payload{
		Lvl:       level.String(),
		Timestamp: time.Now().UTC(),
		Msg:       msg,
		Fields:    args,
	}

	json.NewEncoder(s.out).Encode(m)
	if level == FATAL {
		os.Exit(1)
	}
}
