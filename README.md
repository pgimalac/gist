# gist
Repository to share pieces of code

## branch go-runtime-caller-filename-covervars-go
This branch is a code example where `runtime.Caller` will return `covervar.go` as caller filename
instead of the actual file of the caller.

Requires Go 1.21 to trigger (cannot reproduce with 1.20.x), running `go test` with coverage enabled,
ie. the following command:
```
go test ./... -cover
```
