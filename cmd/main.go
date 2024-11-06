package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/juniorgarcia/ascii-banner-cli/pkg/banner"
)

const appName = "ASCII Banner Generator"

var configFilePath = func() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("Unable to get user home dir.")
		panic(err)
	}

	return homeDir + "/.ascii-banner-cli"
}()

var defaultConfigFileContent = []byte(`TOP_LEFT_CHAR=╭
TOP_RIGHT_CHAR=╮
BOTTOM_LEFT_CHAR=╰
BOTTOM_RIGHT_CHAR=╯
DASHES_CHAR=─
SIDE_CHAR=│
INDIVIDUAL_SIDE_PADDING=2
`)

func createConfigFile(f string) error {
	file, err := os.Create(f)

	if err != nil {
		log.Fatalf("Unable to create default config file at %s. Error: %s", f, err)
		panic(err)
	}

	defer file.Close()

	if _, err := file.Write(defaultConfigFileContent); err != nil {
		log.Fatalf("Unable to write to config file: %s", err)
		panic(err)
	}

	_ = file.Sync()

	return nil
}

func checkConfigFile() bool {
	if _, err := os.Stat(configFilePath); errors.Is(err, fs.ErrNotExist) {
		return false
	}

	return true
}

func main() {
	if !checkConfigFile() {
		if err := createConfigFile(configFilePath); err != nil {
			log.Fatal("Unable to create default config file!", err)
		}
	}

	err := godotenv.Load(configFilePath)

	if err != nil {
		log.Fatal("Error loading env file")
		panic(err)
	}

	msg := flag.String("m", "Hello world!", "The banner message")
	tl := flag.String("tl", os.Getenv("TOP_LEFT_CHAR"), "The top left char")
	tr := flag.String("tr", os.Getenv("TOP_RIGHT_CHAR"), "The top right char")
	bl := flag.String("bl", os.Getenv("BOTTOM_LEFT_CHAR"), "The bottom left char")
	br := flag.String("br", os.Getenv("BOTTOM_RIGHT_CHAR"), "The bottom right char")
	d := flag.String("d", os.Getenv("DASHES_CHAR"), "The dashes character")
	s := flag.String("s", os.Getenv("SIDE_CHAR"), "The sides character")

	padding, _ := strconv.Atoi(os.Getenv("INDIVIDUAL_SIDE_PADDING"))

	p := flag.Int("p", padding, "Individual side padding")
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
