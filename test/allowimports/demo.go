package allowimports

import (
	"bytes" // want `importing forbidden package "bytes"`
	"io"
	"net/http" // want `importing forbidden package "net/http"`
	"strings"
)

func Do() {
	_ = io.EOF
	_ = http.Server{}
	_ = strings.Split("", "")
	_ = bytes.Buffer{}
}
