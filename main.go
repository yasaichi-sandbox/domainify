package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

type listFlag []string

func (list listFlag) String() string {
	return strings.Join(list, ",")
}

func (list *listFlag) Set(s string) error {
	*list = append(*list, strings.Split(s, ",")...)
	return nil
}

var tlds = []string{"com", "net"}

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	var tldListFlag listFlag

	flag.Var(&tldListFlag, "tlds", "comma-separated list of TLDs")
	flag.Parse()
	if len(tldListFlag) != 0 {
		tlds = tldListFlag
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune

		// NOTE: `r` is Rune type
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}

			newText = append(newText, r)
		}

		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
	if err := s.Err(); err != nil {
		log.Fatalln(err)
	}
}
