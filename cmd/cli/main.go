package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Flags struct {
	Counter string
}

func main() {

	var flags Flags
	open, err := os.Open("/dev/tty")
	if err != nil {
		return
	}
	//flagMap := make(map[string]string)
	//fileName, _ := os.Stdin
	sc := bufio.NewScanner(open)
	fmt.Println(sc)

	//fmt.Println(string(all))
	if len(os.Args[1:]) >= 1 {
		arg := strings.Replace(os.Args[1], "-", "", 1)
		flag.StringVar(&flags.Counter, arg, "", "Counts the number of words in the file")
		flag.Usage = func() {
			_, err := fmt.Fprintf(os.Stderr, "Word count implemeng\n")
			if err != nil {
				return
			}
		}
		flag.Parse()
		outputFormatter(flags, arg)
	} else {
		fmt.Println("No argument passed to command")

	}
}

func outputFormatter(flags Flags, arg string) {
	bytes, err := fileParser(flags.Counter)

	switch arg {
	case "c":
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d %s\n", len(bytes), flags.Counter)
		break
	case "w":

		fmt.Printf("%d %s\n", len(strings.SplitAfter(string(bytes), " ")), flags.Counter)
		break
	case "m":
		fmt.Printf("%d %s\n", len(strings.Split(string(bytes), "")), flags.Counter)
		break
	case "h":
		fmt.Printf("%d %s\n", len(strings.Split(string(bytes), "")), flags.Counter)
		break
	case "l":
		counts := len(wordCounter(flags))
		fmt.Printf("%d %s\n", counts, flags.Counter)
		break
	default:
		parser, err := fileParser(flags.Counter)

		if err != nil {

		}
		fmt.Printf("%d %d %d %s\n", len(bytes), len(wordCounter(flags)), len(strings.Fields(string(parser))), flags.Counter)
		break
	}
}

func wordCounter(flags Flags) map[string]int {
	bytes, err := fileParser(flags.Counter)
	if err != nil {
		log.Printf("Error parsing file: %v", err)
		return nil
	}
	words := strings.Fields(string(bytes))
	word := make(map[string]int)
	for _, v := range words {
		word[v] += 1
	}
	fmt.Printf("%d %s\n", len(word), flags.Counter)
	return word
}

func fileParser(file string) ([]byte, error) {
	open, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	all, err := io.ReadAll(open)
	if err != nil {
		return nil, err
	}
	return all, nil
}
