package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/cbebe/aoc"
)

type CPU struct {
	pre     func(*CPU)
	post    func(*CPU)
	display func(*CPU)
	visible []bool
	cycles  int
	total   int
	x       int
}

func NewCPU(pre, post, disp func(*CPU)) *CPU {
	return &CPU{pre, post, disp, []bool{}, 0, 0, 1}
}

func PreA(c *CPU) {}
func PostA(c *CPU) {
	if (c.cycles % 40) == 20 {
		fmt.Printf("cycles: %d, x: %d, ", c.cycles, c.x)
		fmt.Printf("to add: %d, ", c.cycles*c.x)
		c.total += c.cycles * c.x
		fmt.Printf("total: %d\n", c.total)
	}
}
func PrintA(c *CPU) {
	fmt.Println(c.total)
}

func visible(c CPU) bool {
	pos := c.x - 1
	col := c.cycles % 40

	return col == pos || col == pos+1 || col == pos+2
}
func PreB(c *CPU)  { c.visible = append(c.visible, visible(*c)) }
func PostB(c *CPU) {}
func PrintB(c *CPU) {
	for i, v := range c.visible {
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

func (c *CPU) inc() {
	c.pre(c)
	c.cycles++
	c.post(c)
}

func main() {
	// f := "input.txt"
	f := "test.txt"

	// cpu := NewCPU(PreA, PostA, PrintA)
	cpu := NewCPU(PreB, PostB, PrintB)
	_, filename, _, _ := runtime.Caller(0)
	for _, v := range aoc.Lines(f, filename) {
		if v == "" {
			continue
		}
		cpu.inc()
		if v != "noop" {
			cpu.inc()
			cpu.x += aoc.ParseInt(strings.Split(v, " ")[1])
		}
	}
	cpu.display(cpu)
}
