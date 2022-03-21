package skrib

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Level int

func (l Level) String() string {
	switch l {
	case INFO:
		return "info"
	case WARN:
		return "warning"
	case ERROR:
		return "error"
	case DEBUG:
		return "debug"
	case FATAL:
		return "fatal"
	}
}
