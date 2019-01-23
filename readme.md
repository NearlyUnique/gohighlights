# Overview

Highlights of go talk

Slides use [reveal.js](https://revealjs.com)

You can host them locally easily using [Caddy](https://caddyserver.com/), run `./caddy` from the slides directory and visit http://localhost:2015

## Demos

1-6 have srcX directories (X = 1-6)

## Additional Info

### Installing packages

All src folders have `go.mod` files so go build_should just work_

### Testing

```
# basic run tests
go test

# more verbose (all passing is quite terse)
go test -v

# run tests with names in regex
go test -run someregex

# with coverage, cover.out is the coverage data file
go test -coverprofile cover.out
go tool cover -html cover.out

```
