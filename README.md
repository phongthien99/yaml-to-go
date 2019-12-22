# YAML to Go

[![YAML-to-Go converts YAML to a Go struct](https://raw.githubusercontent.com/Zhwt/yaml-to-go/gh-pages/resources/images/yaml-to-go.png)](http://zhwt.github.io/yaml-to-go)

Parse and translates YAML into a Go struct definition. [Check it out!](http://zhwt.github.io/yaml-to-go)

This project is aimed to provide an online service. In case you need the js library only, [click here to download](https://raw.githubusercontent.com/Zhwt/yaml-to-go/gh-pages/resources/js/yaml-to-go.js).

In fact , there is already a yaml-to-go tool by mengzhuo, however [that project](https://github.com/mengzhuo/yaml-to-go) seems abandoned and no longer maintained, so I started my own version.

### Things to note:

- The yaml is parsed by [js-yaml](https://github.com/nodeca/js-yaml).
- The script sometimes has to make some assumptions, so give the output a once-over.
- In an array of objects, the generated object will share all properties among them.
- The output is indented, but not formatted. Use `go fmt`!

Contributions are welcome!

### Credits

The original [JSON-to-Go](https://github.com/mholt/json-to-go) tool is written by [Matt Holt](https://github.com/mholt/).

The Go Gopher is originally by Renee French. This artwork is an adaptation.
