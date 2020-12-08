package main

import (
	"bufio"
	"errors"
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

func compile() ([]Instruction, []Instruction) {
	scanner := bufio.NewScanner(os.Stdin)

	program := make([]Instruction, 0)
	queue := make([]Instruction, 0) // Only jmp and nop

	offset := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		instruction := Instruction{parts[0], parts[1], offset}
		program = append(program, instruction)

		if instruction.operation == "jmp" || instruction.operation == "nop" {
			queue = append(queue, instruction)
		}

		offset++
	}

	return program, queue
}

func recompile(program []Instruction, queue []Instruction, j int) {
	// Get the next jmp or nop instruction to be changed
	next := queue[j]
	newOp, _ := changeOp(next.operation)
	next.operation = newOp
	program[next.offset] = next

	// Restore previously changed instruction
	k := j - 1
	if k >= 0 {
		prev := queue[j-1]
		program[prev.offset] = prev
	}
}

func changeOp(op string) (string, error) {
	var newOp string
	var err error

	switch op {
	case "jmp":
		newOp = "nop"
	case "nop":
		newOp = "jmp"
	default:
		err = errors.New("Unkown op: " + op)
	}

	return newOp, err
}


func run(program, queue []Instruction) int {
	visited := make(map[Instruction]struct{})
	i, j, acc := 0, 0, 0
	for i < len(program) {
		instruction := program[i]

		if _, ok := visited[instruction]; ok {
			recompile(program, queue, j)
			j++

			// Clean to restart
			i, acc = 0, 0
			visited = make(map[Instruction]struct{})

			continue
		} else {
			visited[instruction] = struct{}{}
		}

		i, acc = exec(instruction.operation, instruction.argument, i, acc)
	}

	return acc
}
