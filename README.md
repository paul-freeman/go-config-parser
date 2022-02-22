# go-config-parser

This package provides the `Config` type.

This type is a validated string, where the validation rules are applied when the
type is created. The rules for the `Config` type are that it must contain valid
matching and nested brackets (eg. `"[(){}[]]"`, `"{}[]()"`, etc.) Any other characters are ignored.

This package has no external dependencies.

## Usage

To use this type in you applications, first run
```
go get github.com/paul-freeman/go-config-parser
```

You can then use the package in your application like this:
```go
import github.com/paul-freeman/go-config-parser

config, err := parse.New("(){}[]")
if err != nil {
    panic(err)
}
```

## Testing

Tests are provided at 100% coverage. These can be run using the standard `go test` command.