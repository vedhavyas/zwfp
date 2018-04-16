package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/vedhavyas/zwfp"
)

func main() {
	var key string
	flag.StringVar(&key, "k", "", "key string")
	flag.Usage = func() {
		name := os.Args[0]
		fmt.Println("Usage:")
		flag.PrintDefaults()
		fmt.Printf(`
  Embeding:
    $ cat covertext.txt | %[1]s -k Payload
    $ %[1]s -k Payload CoverText

  Extracting:
    $ cat covertext.txt | %[1]s
    $ %[1]s SteganoText
`, name)
	}
	flag.Parse()

	if key == "" {
		text := flag.Arg(0)
		if text == "" {
			b, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				log.Fatal(err)
			}
			_, key := zwfp.Extract(string(b))
			fmt.Println(key)
		} else {
			pt, key := zwfp.Extract(text)
			fmt.Println("Cover Text:", pt)
			fmt.Println("Payload:", key)
		}
	} else {
		switch flag.NArg() {
		case 0:
			if err := zwfp.Write(os.Stdout, os.Stdin, key); err != nil {
				log.Fatal(err)
			}
		case 1:
			embed := zwfp.Embed(flag.Arg(0), key)
			fmt.Println(embed)
		default:
			flag.Usage()
			os.Exit(1)
		}
	}
}
