# "Hello World" for `h3-go`

(If you've never used `go` before.)


## Get `go`ing

First we'll get a basic Go example running (without `h3-go`).

We'll try this example: https://gobyexample.com/values

On my Mac:

```
brew install go
```

Create the `values.go` file as in the example (also included in this repo),
and then run:

```
go run values.go
```

That should confirm that you can compile and run
basic `go` programs.

## Get `h3-go`ing

Next, there's some `go` package management stuff so our example program `h3example.go` can find the `h3-go` package.

Run:

```shell
go mod init github.com/ajfriend/h3llo_gorld
```

where, as far as I can tell, `github.com/ajfriend/h3llo_gorld` could be any unique identifier. 

Then run:

```shell
go mod tidy
```

You'll end up with 2 new files: `go.mod` and `go.sum`.

Finally, you should be able to run the example:

```shell
go run h3example.go
```

which should print out `8928308280fffff`.
