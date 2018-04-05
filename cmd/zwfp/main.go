package main

import (
	"fmt"
	"os"

	"github.com/vedhavyas/zwfp"
)

func main() {
	switch len(os.Args) {
	case 3:
		embed := zwfp.Embed(os.Args[1], os.Args[2])
		fmt.Println(embed)
		return
	case 2:
		pt, key := zwfp.Extract(os.Args[1])
		fmt.Println("Plain Text:", pt)
		fmt.Println("Embed Key:", key)
		return
	default:
		fmt.Println("Usage:")
		fmt.Println("\t", os.Args[0], "PlainText Key")
		fmt.Println("\t\t", "Embeds Key into PlainText")
		fmt.Println("")
		fmt.Println("\t", os.Args[0], "EmbedText")
		fmt.Println("\t\t", "Extracts Key from EmbedText")
		os.Exit(1)
	}

}
