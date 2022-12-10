package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bitfield/script"
)

type CPU struct {
	cycles  int
	visible []bool
	x       int
}

func visible(c CPU) bool {
	pos := c.x - 1
	col := c.cycles % 40

	return col == pos || col == pos+1 || col == pos+2
}

func inc(c *CPU) {
	c.visible = append(c.visible, visible(*c))
	c.cycles++
}

func main() {
	f := "input.txt"
	// f := "test.txt"
	input, err := script.File(f).String()
	if err != nil {
		log.Fatal(err)
	}

	cpu := CPU{0, []bool{}, 1}
	arr := strings.Split(input, "\n")
	for _, v := range arr {
		if v == "" {
			continue
		}
		inc(&cpu)
		if v != "noop" {
			a := strings.Split(v, " ")
			n, err := strconv.Atoi(a[1])
			if err != nil {
				log.Fatalln(err)
			}
			inc(&cpu)
			cpu.x += n
		}
	}
	for i, v := range cpu.visible {
		if i%40 == 0 {
			fmt.Println()
		}
		val := "#"
		if !v {
			val = "."
		}
		fmt.Print(val)
	}
}
