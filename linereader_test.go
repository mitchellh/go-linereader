package linereader

import (
	"bytes"
	"reflect"
	"testing"
)

func TestReader(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("foo\nbar\n")

	var result []string
	r := New(&buf)
	for line := range r.Ch {
		result = append(result, line)
	}

	expected := []string{"foo", "bar"}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("bad: %#v", result)
	}
}
