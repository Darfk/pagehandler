## Synopsis

Invokes https://github.com/Darfk/page to render HTML pages in response to requests routed by https://github.com/gorilla/mux

## Installation

`go get github.com/darfk/pagehandler`

## Usage

`PageHandler` implements the http.Handler interface.

    package main

    import (
        "github.com/gorilla/mux"
        "github.com/darfk/pagehandler"
        "net/http"
    )

    func main() {
        r := mux.NewRouter()
        r.Handle("/{page:([a-z]+)?}", &pagehandler.PageHandler{})
        http.Handle("/", r)
    }


## License

MIT
