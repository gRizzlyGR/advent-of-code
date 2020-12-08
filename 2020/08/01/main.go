package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	argument  string
	offset    int
}

func main() {
	acc := run(compile())
	fmt.Println(acc)
}

func compile() []Instruction {
	scanner := bufio.NewScanner(os.Stdin)

	program := make([]Instruction, 0)

	offset := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		instruction := Instruction{parts[0], parts[1], offset}
		program = append(program, instruction)
		offset++
	}

	return program
}

func run(program []Instruction) int {
	visited := make(map[Instruction]struct{})

	i, acc := 0, 0
	for i < len(program) {
		instruction := program[i]

		if _, ok := visited[instruction]; ok {
			break
		} else {
			visited[instruction] = struct{}{}
		}

		i, acc = exec(instruction.operation, instruction.argument, i, acc)
	}

	return acc
}

func exec(op string, arg string, i int, acc int) (int, int) {
	n, _ := strconv.Atoi(arg)
	switch op {
	case "acc":
		i++
		acc += n
	case "jmp":
		i += n
	case "nop":
		i++
	}
	return i, acc
}
