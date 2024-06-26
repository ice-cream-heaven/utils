//go:build !(linux && amd64)
// +build !linux !amd64

package json

import (
	"encoding/json"
	"github.com/ice-cream-heaven/utils/stringx"
	"io"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func MarshalString(v any) (string, error) {
	buf, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return stringx.ToString(buf), nil
}

func UnmarshalString(data string, v any) error {
	return json.Unmarshal(stringx.ToBytes(data), v)
}

func NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

func NewDecoder(r io.Reader) *json.Decoder {
	return json.NewDecoder(r)
}
