package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bitfield/script"
)

func toNumBits(mask, num int64) string {
	return strconv.FormatInt(num&mask, 2)
}

func main() {
	input, err := script.File("input.txt").String()
	if err != nil {
		log.Fatal(err)
	}

	numbits := 0
	var bits []int64
	lines := strings.Split(input, "\n")
	var numlines int64 = 0
	for _, v := range lines {
		if v == "" {
			continue
		}
		numlines++
		if numbits == 0 {
			numbits = len(v)
			bits = make([]int64, numbits)
		}
		bin, err := strconv.ParseInt(v, 2, 64)
		if err != nil {
			log.Fatal(err)
		}

		for i := 0; i < numbits; i++ {
			if bin&(1<<i) != 0 {
				bits[i]++
			}
		}
	}

	var mask int64 = 0
	for i := 0; i < numbits; i++ {
		mask <<= 1
		mask |= 1
	}
	fmt.Printf("%v\n", mask)

	var eps int64 = 0
	var gam int64 = 0
	for i := numbits - 1; i >= 0; i-- {
		eps <<= 1
		gam <<= 1
		if bits[i] > (numlines - bits[i]) {
			gam |= 1
		} else {
			eps |= 1
		}
		fmt.Printf("eps: %v, gam: %v\n", toNumBits(mask, eps), toNumBits(mask, gam))
	}
	fmt.Printf("bits: %v, lines: %v\n", bits, numlines)
	fmt.Printf("%v %v\n", eps, gam)
	fmt.Println(eps * gam)
}
