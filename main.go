package main

import (
	"flag"
	"fmt"

	"github.com/juniorgarcia/ascii-banner-cli/banner"
)

const appName = "ASCII Banner Generator"

func main() {

	msg := flag.String("m", "Hello world!", "The banner message")
	tl := flag.String("tl", "+", "The top left char")
	tr := flag.String("tr", "*", "The top right char")
	bl := flag.String("bl", "*", "The bottom left char")
	br := flag.String("br", "*", "The bottom right char")
	d := flag.String("d", "-", "The dashes character")
	s := flag.String("s", "|", "The sides character")
	p := flag.Int("p", 2, "Individual side padding")
	h := flag.Bool("h", false, "Show help")

	flag.Parse()

	if *h {
		printHelp()
		return
	}

	b := banner.NewBanner(*tl, *tr, *bl, *br, *d, *s, *msg, uint8(*p))

	bStr, err := b.PrintFullBanner()
	if err != nil {
		panic(err)
	}

	fmt.Println(bStr)
}

func printHelp() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", appName)
	flag.PrintDefaults()
}
