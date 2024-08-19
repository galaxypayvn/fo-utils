package queryencoder

import (
	"fmt"
	"net/url"
	"testing"
)

// Filter is a struct of filter with a custom URL encoding
type Filter struct {
	Name   string `json:"name,omitempty"`
	Status int    `json:"status,omitempty"`
}

// EncodeValues ...
func (m Filter) EncodeValues(key string, v *url.Values) error {
	v.Set(fmt.Sprintf("%v.%v", key, "name"), m.Name)
	v.Set(fmt.Sprintf("%v.%v", key, "status"), fmt.Sprintf("%v", m.Status))
	return nil
}

func TestValues(t *testing.T) {
	var in = struct {
		V Filter `query:"filter"`
	}{
		V: Filter{
			Name:   "example.ping",
			Status: 1,
		},
	}

	q, err := Values(in)
	if err != nil {
		t.Error(err)
	}
	t.Log(q.Encode())
}
