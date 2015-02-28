riotjs-todo-goemon-livereload-example
=====================================

A browser livereload example of [Riot todo app](https://muut.com/riotjs/guide/) with [mattn/goemon](https://github.com/mattn/goemon).

## Setup

- Install node.js and npm: [Node.js](http://nodejs.org/)
- Install riot globally: `npm install -g riot`
- Install go: [Getting Started - The Go Programming Language](http://golang.org/doc/install)
- Set [the GOPATH environment variable](https://golang.org/doc/code.html#GOPATH)
- Install [mattn/goemon](https://github.com/mattn/goemon) with the following command.

```
go get -u github.com/mattn/goemon/...
```

The command file `goemon` will be installed in `$GOPATH/bin`.

## Run

```
goemon go run main.go
```

- Open http://localhost:3000
- Edit assets/index.html or assets/todo.tag and see your browser to be livereloaded automatically!


## License
MIT
