# Go Analyzer

Provides an analyzer to validate package dependencies against a whitelist.

To use the analyzer, create a YAML configuration file within your repository including a whitelist of package path prefixes:

``` yaml
allow:
  - net/http
  - strings
```

Build the tool, then verify:
``` bash
go build -o allowimports ./cmd/allowimports/main.go
cd test/allowimports
go vet -vettool=$(pwd)/allowimports /
  -config=$(pwd)/test/allowimports/allow.yaml ./test/allowimports/
```