# Tabletop Droll

A CLI dice roller built with Go.

## Set up

- Currently requires Go 1.12 or above
- Currently needs to be built by running `go build` within the directory

## Examples

```sh
# Single die roll
./droll --roll=1d8

or

./droll --roll=d20

# Stacked dice rolls
./droll --roll=4d12

# Multiple dice rolls, flag parses as csv
./droll --roll=4d12,8d6,2d20
```
