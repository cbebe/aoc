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
	var bins []int64
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
		bins = append(bins, bin)

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

	var oxy int64 = 0
	var co2 int64 = 0
	oxynums := make([]int64, len(bins))
	copy(oxynums, bins)
	co2nums := make([]int64, len(bins))
	copy(co2nums, bins)
	for i := numbits - 1; i >= 0; i-- {
		oxy <<= 1
		co2 <<= 1
		if bits[i] >= (numlines - bits[i]) {
			co2 |= 1
		} else {
			oxy |= 1
		}
		fmt.Printf("eps: %v, gam: %v\n", toNumBits(mask, oxy), toNumBits(mask, co2))
	}
	fmt.Printf("%v %v\n", oxy, co2)
	fmt.Println(oxy * co2)
}
