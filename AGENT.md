# AGENT.md — DiPhyx Go Code Conventions (`diphyx` golangci-lint plugin)

This project lints Go with the **`diphyx`** plugin — 15 opinionated code-style rules run as a
[golangci-lint module plugin](https://golangci-lint.run/plugins/module-plugins/) (module
`github.com/diphyx/golangci-plugin`). This file is the agent's guide to those rules: **write
Go to match these conventions from the start**, and when you see a `diphyx` finding, look up
its rule prefix here for the fix.

**How to read a finding**

golangci-lint prints each finding as:

```
path/file.go:5:2: <rule>: <message> (diphyx)
	ready := false
	^
```

- The `(diphyx)` tag means it came from this plugin.
- The message begins with the **rule name** (e.g. `boolnaming:`) — that prefix is the id you
  Ctrl-F in §0 below.

**How to use this file**

1. Read the `<rule>:` prefix of a finding. Ctrl-F it — it's in the lookup table (§0) and a
   detail section.
2. **None of these rules auto-fix.** `golangci-lint run --fix` will NOT touch them — edit by
   hand using the before→after shown here.
3. After editing, re-run the linter to confirm the finding is gone and you didn't trip a
   neighbouring rule (e.g. adding braces can shift where a blank-line rule looks).

**Scope:** every rule runs on **all non-test `.go` files** in the target module (`./...`).
`_test.go` files are skipped by design, so these conventions do not apply to test files.

---

## 0. Fast lookup — every rule, trigger → fix

| Rule (message prefix) | Fires when…                                                            | Fix                                                      |
| --------------------- | ---------------------------------------------------------------------- | -------------------------------------------------------- |
| `boolnaming`          | `x := true` / `x := false` where `x` lacks an approved prefix          | Prefix with `is`/`has`/`can`/`should`/`enable`/`in`      |
| `constnaming`         | a `const` name is not `ALL_CAPS_WITH_UNDERSCORES`                      | Rename to `ALL_CAPS` (upper / digits / `_` only)         |
| `constructornaming`   | exported func returning `*Type` not named `New{Type}`                  | Rename to `New{Type}` (or an action prefix)              |
| `errorcreation`       | `fmt.Errorf("static")` / `errors.New("has %verb")`                     | `errors.New(...)` for static, `fmt.Errorf(...)` for fmt  |
| `errornaming`         | error var named `err`/`e`, or ending in `…Err`                         | Name it `{functionName}Error`                            |
| `exporteddoc`         | exported func/method without a doc comment                             | Add a `// Name …` doc comment                            |
| `functionaloptions`   | exported func returning `func(...)` not named `With…`                  | Rename to `With{Option}`                                 |
| `ifinitstatement`     | an `if` carries an init statement (`if x := …; cond {`)                | Declare the variable on its own line before the `if`     |
| `importgroups`        | a stdlib import after a third-party one, or a blank import not last    | stdlib group first, third-party next, blank imports last |
| `logginglevel`        | `log.Info/Warn/Notice/Trace/Fatal/Panic(...)`                          | Use only `log.Debug` or `log.Error`                      |
| `newlineafterdefer`   | a statement immediately after a `defer` (no blank line)                | Insert a blank line after the `defer`                    |
| `newlinebeforereturn` | a `return` with a non-trivial statement directly above it              | Insert a blank line before the `return`                  |
| `noabbreviation`      | an identifier named `ctx`/`cmd`/`cfg`/`msg`/`req`/`resp`/`res`         | Spell it out (`context`, `command`, …)                   |
| `receivernaming`      | a short (≤2 char) receiver that isn't the camelCase type name          | Name it the camelCase of the full type                   |
| `structnewline`       | no blank line after an embedded field / before the first private field | Insert the blank line                                    |
| `switchcasebraces`    | a `switch`/type-switch case body not wrapped in `{ }`                  | Wrap the case body in a block `{ }`                      |

---

## 1. Naming

**Boolean literals need an intent prefix.** `boolnaming`
Fires only on short declarations whose value is the literal `true`/`false`
(`ready := someFunc()` is not checked). The prefix must be followed by an uppercase letter.
Approved prefixes: `is`, `has`, `can`, `should`, `enable`, `in`. Exempt names: `ok`, `found`,
`done`, `valid`.

```go
isReady := true   // not: ready := true  (or "isready" — prefix needs an uppercase next letter)
hasItems := false
```

**Constants are `ALL_CAPS_WITH_UNDERSCORES`.** `constnaming`
Applies to **every** constant (exported or not). Only uppercase letters, digits, and `_` are
allowed — a single word like `Pi` or `maxSize` fails.

```go
const MAX_SIZE = 10   // not: const maxSize = 10
const RETRY_COUNT = 3
```

**Constructors returning `*Type` are named `New{Type}`.** `constructornaming`
Only exported functions (no receiver) whose first result is `*ExportedType`. Names starting
with `New` — or an action prefix — are accepted: `Parse`, `Collect`, `Read`, `Load`,
`Ensure`, `Get`, `Find`, `Create`, `Build`, `Make`, `Open`, `Connect`, `Start`, `Init`,
`Fetch`, `Extract`, `Resolve`, `Decode`, `Unmarshal`, `From`, `Pointer`, `Clone`, `Copy`,
`Wrap`.

```go
func NewServer() *Server { return &Server{} }   // not: func ProvideServer() *Server
```

**Error variables use the `{functionName}Error` pattern.** `errornaming`
Fires on short declarations (`:=`) **outside** an `if` initializer. Bare `err`/`e` and any
name ending in `…Err` (but not `…Error`) are rejected.

```go
parseError := parse()      // not: err := parse()  /  parseErr := parse()

if parseError := parse(); parseError != nil { … }   // an if-init `err` is exempt from the rule,
                                                     // but prefer the descriptive name anyway
```

**Functional options are named `With{Option}`.** `functionaloptions`
Exported function (no receiver) with exactly one result that is a function type.

```go
func WithName(name string) func(*Command) { … }   // not: func SetName(...) func(*Command)
```

**No abbreviations.** `noabbreviation`
Exact-name match on declarations, parameters, results, and range vars. Spell out:

| Forbidden | Use       | Forbidden | Use        |
| --------- | --------- | --------- | ---------- |
| `ctx`     | `context` | `msg`     | `message`  |
| `cmd`     | `command` | `req`     | `request`  |
| `cfg`     | `config`  | `resp`    | `response` |
|           |           | `res`     | `response` |

**Method receivers are the camelCase of the full type name.** `receivernaming`
Fires only when the receiver is **≤2 characters** and differs from the expected name — it
targets abbreviated one/two-letter receivers.

```go
func (server *Server) Start() {}   // not: func (s *Server) Start()
```

---

## 2. Documentation comments

**Exported functions and methods have a doc comment** (starting with the name). Types, vars,
and consts are not checked by this rule. `exporteddoc`

```go
// Good does a thing.
func Good() {}

func (thing *Thing) Run() {}   // ← missing doc → exporteddoc fires
```

---

## 3. Errors

**Match the constructor to the message.** `errorcreation`

- `fmt.Errorf` with a single argument (no format verbs) → use `errors.New`.
- `errors.New` with a string literal containing `%` → use `fmt.Errorf`.

```go
return errors.New("service not found")              // not: fmt.Errorf("service not found")
return fmt.Errorf("service '%s' not found", name)   // not: errors.New("service %s not found")
```

(Error **variable** naming is `errornaming`, in §1.)

---

## 4. Imports

**Standard library first, third-party next, blank imports last.** `importgroups`
An import is "third-party" when the first path segment contains a dot (`example.com/...`,
`github.com/...`). Ordering — not blank-line grouping — is what's checked.

```go
import (
	"fmt"
	"os"

	"github.com/diphyx/x"

	_ "github.com/lib/pq"   // blank imports go last
)
```

---

## 5. Logging

**Only `Debug` and `Error` levels.** `logginglevel`
Fires on `log.<Level>(…)` where the identifier is literally `log` and the level is one of
`Info`, `Warn`, `Notice`, `Trace`, `Fatal`, `Panic`.

```go
log.Debug().Msg("ok")
log.Error().Msg("failed")   // not: log.Info().Msg(...)
```

---

## 6. Whitespace & layout

**Blank line after a `defer`.** `newlineafterdefer`
A statement directly after a `defer` needs a blank line between them. Exempt: a `defer` of a
function literal (`defer func() { … }()`), and back-to-back `defer`s.

```go
defer file.Close()

process()   // not: process() on the very next line
```

**Blank line before a `return`.** `newlinebeforereturn`
A `return` needs a blank line above it, unless it's the first statement in the block, or the
statement above is a single-line `if … { return }`, a `defer`, or a `case`/`select` clause.

```go
value := compute()

return value   // not: return value directly under the assignment
```

**Struct field spacing.** `structnewline`
A blank line is required **after an embedded field** and **before the first private field**
that follows an exported one.

```go
type Config struct {
	embedded

	Name string

	count int
}
```

---

## 7. Control flow

**No init statement in an `if`.** `ifinitstatement`
An `if` (and `else if`) must not combine a declaration with its condition — this includes the
common error-check form. Move the statement to its own line above the `if` so each line does
one thing.

```go
_, isPresent := byName[step]
if !isPresent {
	doSomething()
}

callError := doWork()
if callError != nil {
	return callError
}

// not:
if _, isPresent := byName[step]; !isPresent {   // ifinitstatement: declare it before the if
	doSomething()
}
if callError := doWork(); callError != nil {    // ifinitstatement: error checks too
	return callError
}
```

**Wrap each switch case body in braces.** `switchcasebraces`
Every `switch` and type-switch case (and `default`) body must be a single block `{ }`.
Exempt: empty cases (`case 1:` with no body) and cases whose last statement is `fallthrough`
(which is illegal inside a nested block).

```go
switch value {
case 1:
	{
		doOne()
	}
default:
	{
		doOther()
	}
}

// not:
switch value {
case 1:
	doOne()      // switchcasebraces: wrap in { }
}
```

---

## Reference

- **Linter name** in config and findings: `diphyx`. Findings read
  `file:line:col: <rule>: <message> (diphyx)`.
- **Module:** `github.com/diphyx/golangci-plugin`. The rule set is the single source of truth
  in `analyzer.All()` (`analyzer/analyzers.go`), shared by the plugin and a standalone
  `go vet`-style command (`cmd/golangci-plugin`).
- **All findings are reports, none auto-fix** — the goal is guidance; fix by hand.
- **Standalone alternative:** `go run github.com/diphyx/golangci-plugin/cmd/golangci-plugin@latest ./...`
  runs the same analyzers as a `go vet`-style tool, without golangci-lint (the `@latest`
  suffix lets Go fetch the command without adding it as a project dependency).

---

## Verify — install, configure, and run the linter on the target project

Because `diphyx` is a **module plugin**, you don't install it as a normal linter — you build a
custom `golangci-lint` binary that embeds it.

**1. Prerequisites:** Go 1.25+ and golangci-lint v2 (`golangci-lint --version`).

**2. Declare the plugin build** — create `.custom-gcl.yml` at the project root:

```yaml
version: v2.10.1 # match your installed golangci-lint
name: custom-gcl
destination: .
plugins:
    - module: github.com/diphyx/golangci-plugin
      version: v0.1.0 # a released tag; or use `path: .` to build from a local checkout
```

**3. Build the custom binary:**

```bash
golangci-lint custom   # produces ./custom-gcl with the diphyx linter compiled in
```

**4. Enable the linter** — create/extend `.golangci.yml` at the project root:

```yaml
version: "2"
linters:
    enable:
        - diphyx
    settings:
        custom:
            diphyx:
                type: module
                description: diphyx code style rules
```

Use `default: none` under `linters` if you want to run **only** the diphyx rules.

**5. Run** the custom binary (not the system `golangci-lint`):

```bash
./custom-gcl run ./...
```

Each finding prints as `file:line:col: <rule>: <message> (diphyx)`. Look up the `<rule>:`
prefix in §0 and fix by hand — `--fix` will not touch these rules.

**6. Verify a single rule / package** while iterating:

```bash
./custom-gcl run ./path/to/pkg/...                       # one package tree
./custom-gcl run ./... 2>&1 | grep 'switchcasebraces:'   # only one rule's findings
```

**7. Sanity-check the setup itself** — confirm the plugin loaded:

```bash
./custom-gcl linters | grep diphyx   # a line means the linter is registered
```

**Expected outcome:** `custom-gcl` runs without config errors, `diphyx` appears in
`custom-gcl linters`, and any convention violations surface as `(diphyx)` findings you can
resolve using the sections above. Re-run step 5 after edits until the `diphyx` findings are
gone.
