# ASCII Banner CLI

## A simple tool to create ASCII Banners

### Install

1. Make sure [Go is installed on your machine](https://golang.org/doc/install#download)
2. Run on terminal `go get github.com/juniorgarcia/ascii-banner-cli && go install github.com/juniorgarcia/ascii-banner-cli`
3. Be sure that `$GOPATH/bin` is on your `$PATH` variable.

The binary will be available under the name `ascii-banner-cli`

### Usage

With the CLI installed, you can run `ascii-banner-cli -h` to obtain help on how
to use it. It's really simple.

A simple banner could be created like this:

```
$ ascii-banner-cli -m 'Go rocks!'
╭─────────────╮
│  Go rocks!  │
╰─────────────╯
```

### Config file

When ran, this CLI will create a config file at `~/.ascii-banner-cli` if this
file does not exist.

It is a simple dotenv file and will contain the default configuration to create
your banners. Check this file to see what you can customize.

### Suggestions

Since this a repository only to play around and learn more Go, you may open
issues to comment about improvements, best code practices, bugs, etc. I'll be
glad to talk to you.
