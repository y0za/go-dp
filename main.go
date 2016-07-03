package main

import (
	"fmt"
	"time"
)

var pineapple = []string{
	` \ヘ､Ww〃へ`,
	`  ヘヾヾV7〃`,
	`   ヾV〃へ`,
	`  γ^i^i^i^ヽ`,
	` {^i^i^i^i^｝`,
	`｛^i^i^i^i^｝`,
	` ヾ,^i^i^iノ`,
	`   ｀￣￣'`,
}

func main() {
	fmt.Printf("\x1b[2J")
	for i := 68; i > 0; i-- {
		for j, row := range pineapple {
			fmt.Printf("\x1b[%d;%dH\x1b[2K", j+1, i)
			fmt.Println(row)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
