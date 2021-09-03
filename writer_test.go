package skrib

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

const (
	EXP = `\w\S\d{4}/\d{2}/\d{2}\s\d{2}:\d{2}:\d{2}\s\w+`
)

func TestWrite(t *testing.T) {
	var batch []Log = []Log{
		{LogType: ERROR, Message: fmt.Sprintf("%v", errors.New("Some kind of ERROR"))},
		{LogType: SUCCESS, Message: fmt.Sprintf("%v", errors.New("Some kind of SUCCESS"))},
		{LogType: WARNING, Message: fmt.Sprintf("%v", errors.New("Some kind of WARNING"))},
		{LogType: INFO, Message: fmt.Sprintf("%v", errors.New("Some kind of INFO"))},
	}

	for _, j := range batch {
		s := strings.Split(j.Message, " ")
		if !reflect.DeepEqual(j.Coloring().Prefix, s[len(s)-1]+"\t") {
			t.Fail()
		}
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stdout)
		}()
		j.Write(false)
		re := regexp.MustCompile(EXP)
		if m := re.MatchString(buf.String()); !m {
			t.Fail()
		}
	}
}
