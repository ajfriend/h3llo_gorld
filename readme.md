# "Hello World" for `h3-go`

(If you've never used `go` before.)


## Get `go`ing

First we'll get a basic Go example running (without `h3-go`).

We'll try this example: https://gobyexample.com/values

On my Mac:

```
brew install go
```

Create the `values.go` file as in the example (also included here),
and then run:

```
go run values.go
```

## Get `h3-go`ing

Next there's some `go` package management stuff.

Run:

```shell
go mod init github.com/ajfriend/h3llo_gorld
```

where, as far as I can tell, `github.com/ajfriend/h3llo_gorld` is any unique identifier. 

Then run:

```shell
go mod tidy
```

You'll end up with 2 new files: `go.mod` and `go.sum`.

Now, you should be able to run the example:

```shell
go run h3example.go
```