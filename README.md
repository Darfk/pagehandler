## PageHandler

Invokes https://github.com/Darfk/page to render HTML pages

## Installation

`go get github.com/darfk/pagehandler`

## Usage

`PageHandler` implements the http.Handler interface.

    package main

    import (
        "github.com/darfk/pagehandler"
        "net/http"
    )

    func main() {
		mux := http.NewServeMux()
		mux.Handle("/", &pagehandler.PageHandler{})
		server := &http.Server{}
		server.Handler = mux
		server.ListenAndServe()
    }


## License

MIT
