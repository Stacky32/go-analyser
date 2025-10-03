package testdata

import (
	"bytes"
	"encoding/json" // want `importing forbidden package "encoding/json"`
	"net/url"       // want `importing forbidden package "net/url"`
)

func Do2() {
	_ = json.Decoder{}
	_ = url.Values{}
	_ = bytes.Buffer{}
}
