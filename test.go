package main

import (
	"fmt"
)

func main() {
	const STR = `
`+ "\x60" +`
aa`
	fmt.Printf("%x", '世')
	fmt.Println(STR, len('世界'))
}
