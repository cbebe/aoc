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

func main() {
	input, err := script.File("input.txt").String()
	// input, err := script.File("test.txt").String()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(input, "\n")
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
