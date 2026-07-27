# bh-piscine

A flat collection of small, standalone exercises from the Reboot01/01-edu "piscine" (the intensive onboarding bootcamp), covering foundational Go programming and basic Unix shell/CLI usage. Most exercises are a single Go file implementing one function or program; a handful are one-line shell answers.

## About the exercise

This is coursework from the piscine's early Go and shell modules. The Go exercises train writing small, self-contained functions without relying on the standard library's higher-level helpers (e.g. implementing `IsPrime`, `Atoi`, string reversal, factorial, and similar primitives by hand, and printing output through the school's own `z01` package instead of `fmt`). The shell exercises train basic `find`/`ls`/`curl`/`jq` usage, including querying a superhero JSON API for specific records.

## Tech stack

- Go 1.21 (module `piscine`, see `go.mod`)
- `github.com/01-edu/z01` — the school's low-level print helper library, used by the `main.go` programs to print output character-by-character instead of via `fmt`
- POSIX shell (`.sh` files and a few extension-less answer files) using `find`, `ls`, `curl`, `jq`

## Project structure

The repo is a flat directory of exercises rather than one application. Representative layout:

```
go.mod, go.sum            # Go module "piscine", depends on github.com/01-edu/z01

*.go (top level, ~50 files)  # each implements one function in `package piscine`,
                              # e.g. isprime.go, fibonacci.go, strrev.go,
                              # basicatoi.go, iterativefactorial.go, sqrt.go,
                              # rot14.go, convertbase.go, map.go, foreach.go

<name>/main.go             # standalone runnable programs (package main),
                              # e.g. printalphabet/, printdigits/, sortparams/,
                              # revparams/, displayfile/, point/, boolean/, ztail/
displayfile/quest8.txt      # sample input file used by displayfile/main.go

hello.sh, myfamily.sh,      # one-line shell exercises: greeting script,
who-are-you.sh, ...          # curl+jq lookups against a superhero JSON API

look, mastertheLS, r,       # extension-less files holding a single shell
to-git-or-not-to-git.sh       # command/answer each (find, ls, curl+jq, git)
```

## Build & run

Each top-level `*.go` file declares `package piscine` and is a library function rather than a runnable program — these are meant to be built/tested as part of the `piscine` module (`go build ./...` from the repo root, or `go vet ./...`).

The subfolders (`printalphabet/`, `sortparams/`, `displayfile/`, `point/`, `boolean/`, `ztail/`, etc.) each contain a `package main` with their own `main.go`, and can be run directly, for example:

```bash
go run ./printalphabet
go run ./displayfile
```

The `.sh` files and extension-less answer files are single shell commands; run them with `sh <file>` (or `bash <file>`) after making them executable, e.g.:

```bash
bash hello.sh
bash myfamily.sh
```
