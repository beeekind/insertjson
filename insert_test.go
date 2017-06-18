package insertjson_test

import (
	"bytes"
	"fmt"
	"github.com/b3ntly/insertjson"
	"testing"
)

type testCase struct {
	key    string
	value  string
	input  []byte
	output []byte
}

func TestInsert(t *testing.T) {
	tests := []*testCase{
		&testCase{
			key:    "id",
			value:  "1",
			input:  []byte(`{"name":"fred", "thing":"foo"}`),
			output: []byte(`{"id":"1","name":"fred", "thing":"foo"}`),
		},

		&testCase{
			key:    "id",
			value:  "1",
			input:  []byte(`{"name":"fred", "thing":{"foo":"bar"}`),
			output: []byte(`{"id":"1","name":"fred", "thing":{"foo":"bar"}`),
		},
	}

	for _, test := range tests {
		t.Run("will insert a KV pair as expected", func(t *testing.T) {
			result := insertjson.Property(test.key, test.value, test.input)
			if bytes.Compare(test.output, result) != 0 {
				t.Log(fmt.Sprintf("Failed to insert Key: %v into Input: %v", test.key, result))
				t.FailNow()
			}
		})
	}
}
