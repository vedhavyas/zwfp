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
		fmt.Println("Cover Text:", pt)
		fmt.Println("Payload:", key)
		return
	default:
		fmt.Println("Usage:")
		fmt.Println("\t", os.Args[0], "CoverText Payload")
		fmt.Println("\t\t", "Embeds Payload into CoverText")
		fmt.Println("")
		fmt.Println("\t", os.Args[0], "StegoText")
		fmt.Println("\t\t", "Extracts Payload from StegoText")
		os.Exit(1)
	}

}
