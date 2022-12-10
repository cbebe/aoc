package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bitfield/script"
)

type CPU struct {
	cycles int
	total  int
	x      int
}

func inc(c *CPU) {
	c.cycles++
	if (c.cycles % 40) == 20 {
		fmt.Printf("cycles: %d, x: %d, ", c.cycles, c.x)
		fmt.Printf("to add: %d, ", c.cycles*c.x)
		c.total += c.cycles * c.x
		fmt.Printf("total: %d\n", c.total)
	}
}

func main() {
	f := "input.txt"
	// f := "test.txt"
	input, err := script.File(f).String()
	if err != nil {
		log.Fatal(err)
	}

	cpu := CPU{0, 0, 1}
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
	fmt.Println(cpu.total)
}
