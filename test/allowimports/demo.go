package allowimports

import (
	"io"
	"net/http" // want `importing forbidden package "net/http"`
)

func Do() {
	_ = io.EOF
	_ = http.Server{}
}
