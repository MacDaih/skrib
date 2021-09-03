package skrib

import (
	"fmt"
)

func HandleError(err error, exit bool) {
	if err != nil {
		var l Log = Log{
			LogType: ERROR,
			Message: fmt.Sprintf("%v", err),
		}
		l.Write(exit)
	}
}

func HandleLog(msg string, t int) {
	l := Log{
		LogType: t,
		Message: msg,
	}
	l.Write(false)
}
