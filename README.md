# Testing main() in package main

This repo contains an example implementation of Go testing code that
tests the `main()` function in `package main`.

While the `go test` tool is designed primarily for unit testing of
packages, it can also be used for testing a command by calling its
`main()` function as if it were called on the command line.

Note that [`TestMain(m *testing.M)`](https://pkg.go.dev/testing#hdr-Main)
is not for testing the `main()` function of `package main` and serves
a different purpose.  Its main purpose is to facilitate extra setup and
teardown steps before and after running tests.

To try this repo, you can do the following:

```
git clone https://github.com/SaiedKazemi/testmain.git
cd testmain
go test -v .
go test -v -forcefailure .
```
