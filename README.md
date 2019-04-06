# getignore

A tool for getting .getignore files.

### Install

Installing from sources (Mac OS X instructions)

```bash
go get github.com/nvictor/getignore/cmd/getignore
go build github.com/nvictor/getignore/cmd/getignore
install getignore /usr/local/bin/
```

See TODO about `brew` formulae.

### Usage:

```bash
getignore python
```

The above command will get the `Python.getignore` file and save it to `.gitignore` in the current folder.

### TODO:

1. Brew formulae
1. Other platforms?

### Misc:

See https://github.com/github/gitignore for more.

Thanks [Github](https://github.com)!
