package main

import (
	"fmt"
	"github.com/DaviRain-Su/headfirstgo/src/chapter14"
)

func main() {
	values := []string{"apple", "orange", "pear", "banana"}
	fmt.Println(chapter14.JoinWithCommas(values))
}
