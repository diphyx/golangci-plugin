# golangci-plugin

A [golangci-lint](https://golangci-lint.run) module plugin that enforces
**diphyx** Go code-style rules. It can also run standalone as a `go vet`-style tool.

The plugin registers under the name **`diphyx`**.

## Rules

| Analyzer              | Enforces                                                                                                                        |
| --------------------- | ------------------------------------------------------------------------------------------------------------------------------- |
| `boolnaming`          | Boolean variables must start with `is`, `has`, `can`, `should`, `enable`, or `in` (exceptions: `ok`, `found`, `done`, `valid`). |
| `constnaming`         | Constants must be `ALL_CAPS_WITH_UNDERSCORES`.                                                                                  |
| `constructornaming`   | Exported functions returning `*Type` must be named `New{Type}` (action-prefixed names like `Parse`, `Get`, `Build` are exempt). |
| `errorcreation`       | Use `errors.New` for static messages and `fmt.Errorf` for formatted ones.                                                       |
| `errornaming`         | Error variables must follow the `{functionName}Error` pattern; bare `err`/`e` and the `…Err` suffix are rejected.               |
| `exporteddoc`         | Exported functions and methods must have a documentation comment.                                                               |
| `functionaloptions`   | Exported functions returning `func(*T)` must be named `With{Option}`.                                                           |
| `importgroups`        | Standard library imports must precede third-party imports; blank imports go last.                                               |
| `logginglevel`        | Only `Debug` and `Error` log levels are allowed (no `Info`, `Warn`, `Notice`, `Trace`, `Fatal`, `Panic`).                       |
| `newlineafterdefer`   | A blank line is required after a `defer` statement.                                                                             |
| `newlinebeforereturn` | A blank line is required before a `return` statement.                                                                           |
| `noabbreviation`      | Forbids `ctx`, `cmd`, `cfg`, `msg`, `req`, `resp`, `res` — use the full word.                                                   |
| `receivernaming`      | Method receivers must be the camelCase of the full type name.                                                                   |
| `structnewline`       | A blank line is required after embedded fields and before the first private field.                                              |
| `switchcasebraces`    | Each `switch`/type-switch case body must be wrapped in a block `{ }` (empty and `fallthrough` cases are exempt).                |

## Use as a golangci-lint plugin

Build a custom `golangci-lint` binary that bundles this plugin. From a project
that wants to use it, add `.custom-gcl.yml`:

```yaml
version: v2.10.1
plugins:
    - module: github.com/diphyx/golangci-plugin
      version: v0.1.0
```

Build the custom binary and enable the linter in `.golangci.yml`:

```yaml
version: "2"
linters:
    settings:
        custom:
            diphyx:
                type: module
                description: diphyx code-style rules
    enable:
        - diphyx
```

```sh
golangci-lint custom        # produces ./custom-gcl
./custom-gcl run ./...
```

## Use as a standalone tool

```sh
go install github.com/diphyx/golangci-plugin/cmd/golangci-plugin@latest
golangci-plugin ./...
```

## Development

```sh
go build ./...
go test ./...               # runs analysistest fixtures in analyzer/testdata
```

Each analyzer has a fixture under `analyzer/testdata/src/<analyzer>/` using the
standard `// want "regexp"` convention.
