package zerodriver

import (
	"strings"
	"sync"

	"github.com/rs/zerolog"
)

type LabelSource struct {
	Key   string
	Value string
}

type labels struct {
	store map[string]interface{}
	mutex *sync.RWMutex
}

func newLabels() *labels {
	return &labels{store: map[string]interface{}{}, mutex: &sync.RWMutex{}}
}

// Label adds an optional label to the payload in the format of `key: value`.
// Key of label must start with `labels.`
func Label(key, value string) *LabelSource {
	return &LabelSource{Key: "labels." + key, Value: value}
}

// Labels takes label structs, filters the ones that have their key start with the
// string `labels.` and their value type set to string type. It then wraps those
// key/value pairs in a top-level `labels` namespace.
func (e *Event) Labels(labels ...*LabelSource) *zerolog.Event {
	lbls := newLabels()

	lbls.mutex.Lock()
	for i := range labels {
		if isLabelEvent(labels[i]) {
			lbls.store[strings.Replace(labels[i].Key, "labels.", "", 1)] = labels[i].Value
		}
	}
	lbls.mutex.Unlock()

	return e.Event.Dict("logging.googleapis.com/labels", zerolog.Dict().Fields(lbls.store))
}

func isLabelEvent(label *LabelSource) bool {
	return strings.HasPrefix(label.Key, "labels.")
}
