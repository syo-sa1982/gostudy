package trace
import (
	"testing"
	"bytes"
)
func TestNew(t *testing.T) {
	t.Error("まだテスト作ってない")
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("まだテスト作ってない")
	} else {
		tracer.Trace("こんにちはtraceパッケージ")
		if buf.String() != "こんにちはtraceパッケージ\n" {
			t.Errorf("'%s'という誤った文字列", buf.String())
		}
	}
}
