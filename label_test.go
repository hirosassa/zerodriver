package zerodriver

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestLabel(t *testing.T) {
	t.Parallel()

	l := Label("key", "value")
	assert.Equal(t, &label{key: "labels.key", value: "value"}, l)
}

func TestLabels(t *testing.T) {
	t.Parallel()

	var e Event
	tests := map[string]struct {
		expect *zerolog.Event
		labels []*label
	}{
		"success": {
			expect: e.Dict("logging.googleapis.com/labels", zerolog.Dict().Fields(map[string]interface{}{"foo": "bar", "baz": "qux"})),
			labels: []*label{Label("foo", "bar"), Label("baz", "qux")},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.expect, e.Labels(tt.labels...))
		})
	}
}

