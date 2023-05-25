# dump

Dumping ground to play with Go.

## Links
- Go package documentation: https://pkg.go.dev/, e.g. https://pkg.go.dev/encoding/json
- Examples for everything: https://gobyexample.com/
- Lorem ipsum text generator API: https://loripsum.net/api

## How to run

From the root of this repo:

    $ go run .
    {...,"quo":1,"recte":1,"sed":2,"semper":1,"si":3,...}

## Where to take this
- Refactor into a separate function
- Change the JSON output from
    ```json
    {"word1": 1, "word2": 2, ...}
    ```
    to
    ```json
    [{"word1": 1}, {"word2": 2}, ...]
    ```
- Sort the list in order of the most frequently used words.
- Add unit tests: https://go.dev/doc/code#Testing
- Take a [`io.Reader`](https://pkg.go.dev/io#Reader) as input.
- Create a more full-fledged CLI that can read from [`*os.File`s](https://pkg.go.dev/os#File) (like `os.Stdin`) and URLs.
    - See [`flag`](https://pkg.go.dev/flag) and [`os.Args`](https://pkg.go.dev/os#pkg-variables)
- Use a [`bufio.Scanner`](https://pkg.go.dev/bufio#Scanner) to split words.
- Expose this functionality from an HTTP endpoint.
