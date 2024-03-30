# Gobin

A simple pastebin clone written in go and templ.

## Installing templ cli

```sh
go install github.com/a-h/templ/cmd/templ@latest
```

## Running

```sh
templ generate
go run .
```

### Database information

If it displays `CGO_ENABLED=0` related error, do:

```sh
go env -w CGO_ENABLED=1
```

You will have to install gcc for it to work. On windows I reccomend using scoop:

```powershell
scoop install gcc
```
