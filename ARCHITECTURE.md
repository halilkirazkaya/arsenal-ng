# arsenal-ng Architecture

Modern pentest command launcher written in Go. Loads ~2900 cheat commands
from embedded YAML files, lets the user fuzzy-search and fill arguments in a
Bubble Tea TUI, then injects the final command into the shell's input buffer
via TIOCSTI so it can be reviewed before execution.

Stack: Go 1.26, Bubble Tea v2 / Bubbles v2 / Lip Gloss v2 (`charm.land`),
`go.yaml.in/yaml/v4`, `golang.org/x/sys`.

## Data Flow

```
cheat-files/*.yaml ──► loader ──► search ──► ui (TUI) ──► prefill ──► shell
                                          ▲
                                    variables.json
```

1. **Load**: `loader` walks the embedded `cheat-files` tree, parses YAML
   into `model.Cheat` values.
2. **Search**: `search` filters and ranks cheats per keystroke.
3. **Interact**: `ui` runs the Bubble Tea program (search → args → execute).
4. **Prefill**: `prefill` writes the chosen command to `/dev/tty` with the
   TIOCSTI ioctl. Global variables persist in `variables.json`.

## Project Tree

```
cmd/arsenal-ng/          Entry point: main + run + injectable deps
internal/
  config/                App identity: name, version, logo, UI defaults
  frame/                 Shared TUI frame: header, footer, usable width
  highlight/             Fuzzy-match and command-syntax highlighting
  keys/                  Key binding strings and search-box command literals
  loader/                Embedded YAML loading (cheat-files/ lives here)
  model/                 Cheat/Argument types, parsing, building, validation
  paths/                 OS config directory resolution
  prefill/               TIOCSTI terminal injection (darwin/linux)
  search/                Field-weighted cheat ranking
  textutil/              Truncate and word-wrap helpers
  theme/                 Color palette and all Lip Gloss styles
  variables/             Persistent global variable store (variables.json)
  ui/                    Bubble Tea parent: routes messages, owns transitions
    args/                Argument input form view
    finder/              Main search view (filter, navigate, hints)
    help/                Help screen view
    tablex/              Shared styled-table constructors
    tools/               Paginated tools table view
    variables/           Variables table view
```

## Package Details

### Entry Point

- `cmd/arsenal-ng/main.go` — `main` exits non-zero on error; `run` wires
  logging, loading, the TUI session, and the prefill fallback message.
- `cmd/arsenal-ng/deps.go` — `deps` interface abstracts side effects
  (log path, cheat load, TUI run, prefill, warnings); `production`
  implements it with the real backends.

### Domain

- `internal/model/cheat.go` — `CheatFile`/`Action` (YAML shape) and the
  flattened runtime `Cheat` plus `Argument`.
- `internal/model/parse.go` — `{{arg}}` / `{{arg|default}}` extraction
  with first-appearance order and dedup.
- `internal/model/build.go` — Single-pass placeholder substitution.
- `internal/model/validate.go` — `HasEmptyArgs` / `GetEmptyArgs` for
  required arguments.
- `internal/model/tools.go` — `ToolSummary`: unique tools with command
  counts, alphabetically sorted.

### Loading and Search

- `internal/loader/loader.go` — `Load` reads the embedded tree;
  `loadFromFS` takes any `fs.FS` so error paths are testable.
- `internal/loader/parse.go` — YAML bytes → `Cheat` values, `isYAMLFile`.
- `internal/search/search.go` — `Query`: all-terms-must-match filter.
- `internal/search/score.go` — Per-term field weights
  (tool 100 > tag 60 > title 40 > command 20 > desc 10) with
  exact ×4 / prefix ×2 / substring ×1 multipliers, summed across fields.
- `internal/search/rank.go` — Stable sort by score, ties keep load order.

### State and Paths

- `internal/variables/store.go` — Thread-safe `Store` with JSON
  persistence (`New`, `Set`, `Get`, `Unset`, `All`, `Count`).
- `internal/variables/file.go` — Atomic save (temp file + rename, 0600)
  and tolerant load (missing/empty/null).
- `internal/paths/paths.go` — `ConfigDir`, `Variables`, `Log`;
  XDG on Linux, Application Support on macOS.
- `internal/config/` — `AppName`, `Version` (ldflags-overridable),
  ASCII `Logo`, terminal/input defaults.

### Terminal Output

- `internal/prefill/prefill.go` — `Command` opens `/dev/tty`, disables
  echo/canonical mode, injects bytes, always restores termios.
- `internal/prefill/ioctl.go` — Platform TIOCSTI/TCGETS/TCSETS constants.

### UI Shell

- `internal/ui/app.go` — `New` builds the app with isolated globals;
  `Init` returns the cursor blink.
- `internal/ui/model.go` — `App` composes one sub-model per view plus
  shared cheats, globals, dimensions, and the result.
- `internal/ui/update.go` — Global resize handling, message routing.
- `internal/ui/view.go` — `View() tea.View` wraps the active string view
  and enables the alternate screen.
- `internal/ui/search_flow.go` — Search outcomes: pick, show-vars/tools/
  help, set/unset, cancel.
- `internal/ui/flows.go` — Args/tools outcomes and prefill counting.
- `internal/ui/variables_flow.go` — Variables outcome with lazy rebuild.
- `internal/ui/transitions.go` — `showVariables`, `showTools`,
  `showHelp`, `backToSearch`.
- `internal/ui/variables.go` — `set`/`unset` command execution against
  the store with status feedback.
- `internal/ui/tablex/` — `New`/`NewWide` styled tables, `Fit` stretches
  the last column, `Width`/`Height` size math.

### UI Views (one MVU triple each)

- `internal/ui/finder/` — Main view: input, ranked list, cursor/page
  navigation, `set`/`unset`/`variables`/`tools`/`help` hints, footer.
  Reports `Outcome` (`Stay`, `Picked`, `Show*`, `SetVar`, `UnsetVar`,
  `Cancel`); the parent applies cross-view effects.
- `internal/ui/args/` — Per-argument text inputs with global/default
  prefill; `Submit` carries the built command (empty = blocked).
- `internal/ui/tools/` — Paginated tool table; `Pick` returns the tool
  name to search for.
- `internal/ui/variables/` — Variables table; any exit key goes back.
- `internal/ui/help/` — Stateless shortcut/workflow reference.

### Shared UI Support

- `internal/frame/` — `Header` (logo + version), `Footer` (position,
  variable count, key hints), `Width` clamp.
- `internal/theme/` — `palette.go` hex constants; `colors`,
  `components`, `containers`, `syntax` Lip Gloss styles; `tags.go`
  hash-colored tag rendering.
- `internal/highlight/` — `Fuzzy` rune-based match highlighting;
  `Syntax` command coloring via `tokenize` + `operators` + `spans`.
- `internal/keys/` — `tea.KeyPressMsg.String()` literals
  (`"enter"`, `"ctrl+c"`, …) and search-box commands.
- `internal/textutil/` — `Truncate` with ellipsis, `Wrap` at word
  boundaries.

## Conventions

- Files stay near ~100 lines; packages own one concern.
- Every file uses 80-char `// ===` section banners and English doc
  comments starting with the identifier name.
- Tests live next to sources, use isolated HOME dirs, and pin the logo
  color; suite runs with `go test ./... -count=1`.
