# `cobra` CLI testing demo

Demo repo showing 2 strategies for testing `golang` `cobra` command line apps
inspired by
[[1]](https://gianarb.it/blog/golang-mockmania-cli-command-with-cobra),
[[2]](https://github.com/spf13/cobra/issues/770)
and references therein.

The 2 strategies are

- var_example
    - keeps the original/default `var &cobra.Command` top level definition of the app
- func_example
    - changes the fundamental object from a `var` to a `func <name>() *cobra.Command` that returns a `cobra.Command` ref

## usage

- running either app `go run main.go --toggle` will run OK
    - errors if fail to supply `--toggle`/`-t` arg
- testing either app `go test ./...`
