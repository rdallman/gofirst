# superbasicfragilisticexpialiGOcious

(I'm easier to read in a web browser, unless you get / like markdown syntax)

<http://github.com/rdallman/gofirst/blob/master/README.md>

Just in case, a `$` just denotes to type in shell, don't actually type the `$`

### go get it (words to live by)

`
$ go get github.com/rdallman/gofirst
`

__if this doesn't work, you'll need to fix your $GOPATH__

If something went wrong suggestions:

* `go help gopath`
* "No Git command" on Windows <http://git-scm.com/>

Raise your hand, we're here to help


### If you build it, they will come

`
$ cd $GOPATH/src/github.com/rdallman/gofirst
`

`
$ go install
`

### run the created binary

Linux/Mac

`
$ gofirst
`
Windows

`
$ gofirst.exe
`

If this didn't work, you'll need to add `$GOPATH/bin` to your `$PATH`, or
you can type `$ $GOPATH/bin/gofirst` in shell to run from here on out

### open a browser, type in `localhost:8080/gopher`

P.S. try some other urls, you'll catch on

### That's all there is to it! _go_ dig around in `firstserver.go`

`
$ EDITOR firstserver.go
`

where `EDITOR` = `vim`, `subl`, `nano`, `emacs`, etc.

### When you're comfortable with this...

```
$ go get github.com/rdallman/goserver
$ cd $GOPATH/src/github.com/rdallman/goserver
```

Open the README: <https://github.com/rdallman/goserver/blob/master/README.md>

You can close out of this readme now, but it might be handy to keep `firstserver.go` open in your
editor for the first few steps of the next tutorial.

### BONUS for vimmers (we gotta stick together)

Go is awesome and provides syntax highlighting and other support for `vim`

First:

`
$ echo $GOROOT
`

If nothing prints:

`
$ go env GOROOT
`

If you had to do `go env`, then copy that value in for `$GOROOT` below

Then, copy and paste the following into your `~/.vimrc`

```
filetype off
filetype plugin indent off
set rtp+=$GOROOT/misc/vim
filetype plugin indent on
syntax on
```

__Put that in exactly__, even if you have any of the other stuff already set, it's 
really weird about reloading everything.
