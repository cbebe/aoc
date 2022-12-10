package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"

	"github.com/cbebe/aoc"
)

func toNumBits(mask, num int64) string {
	return strconv.FormatInt(num&mask, 2)
}

func main() {
	f := "input.txt"

	_, filename, _, _ := runtime.Caller(0)
	lines := aoc.Lines(f, filename)
	// PartA(lines)
	PartB(lines)
}

func PartA(lines []string) {
	numbits := 0
	var bits []int64
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

func findBit(bins []string, most bool) int64 {
	cp := make([]string, len(bins))
	copy(cp, bins)
	for i := 0; i <= len(bins[0]); i++ {
		if len(cp) == 1 {
			num, err := strconv.ParseInt(cp[0], 2, 64)
			if err != nil {
				log.Fatalln(err)
			}
			return num
		}
		n := 0
		for j := 0; j < len(cp); j++ {
			q := cp[j][i]
			if q == byte('1') {
				n++
			}
		}
		fmt.Printf("%d ", n)
		var common byte
		if most {
			if n >= (len(cp) - n) {
				common = '1'
			} else {
				common = '0'
			}

		} else {
			if n >= (len(cp) - n) {
				common = '0'
			} else {
				common = '1'
			}
		}

		newCp := make([]string, 0)
		for j := 0; j < len(cp); j++ {
			q := cp[j][i]
			if q == common {
				newCp = append(newCp, cp[j])
			}
		}
		cp = newCp
		fmt.Printf("%v\n", cp)
	}
	return -1
}

func PartB(lines []string) {
	numbits := 0
	var numlines int = 0
	var bins []string
	for _, v := range lines {
		if v == "" {
			continue
		}
		numlines++
		w := strings.TrimSpace(v)
		if numbits == 0 {
			numbits = len(w)
		}
		bins = append(bins, w)
	}

	var oxy int64 = findBit(bins, true)
	var co2 int64 = findBit(bins, false)
	fmt.Printf("%v %v\n", oxy, co2)
	fmt.Println(oxy * co2)
}
