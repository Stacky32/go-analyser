package testdata

import (
	"net/http" // want `importing forbidden package "net/http"`
)

func Do1() {
	http.ListenAndServe(":8080", nil)
}
