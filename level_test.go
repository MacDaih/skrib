package skrib

import "testing"

func TestLogLevel(t *testing.T) {
	lvls := map[Level]string{
		0: "debug",
		1: "info",
		2: "warning",
		3: "error",
		4: "fatal",
	}

	for k, v := range lvls {
		if k.String() != v {
			t.Errorf("mismatch level value : expects %s, got %s", k.String(), v)
		}
	}
}
