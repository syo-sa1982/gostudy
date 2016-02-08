package trace
import (
	"testing"
	"bytes"
)
func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("返り値がnil")
	} else {
		tracer.Trace("こんにちはtraceパッケージ")
		if buf.String() != "こんにちはtraceパッケージ\n" {
			t.Errorf("'%s'という誤った文字列", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("でーた")
}
