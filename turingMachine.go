package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	turingMachine := initMachine()
	turingMachine.printHelp()
	for turingMachine.loop {
		var input string
		fmt.Scan(&input)
		turingMachine.choice(input)
	}
}


type Element struct{
	prev, next *Element
	value bool
}

type Machine struct { 
	head, currentHead *Element 
	loop bool
}

func initMachine() Machine {
	nElements := 45
	newElement := &Element{ prev: nil, next: nil, value: false }
	current := newElement
	for i := 0; i < nElements-1; i++{
		(*current).next = &Element{ prev: current, next: nil, value: false }
		current = current.next
	}
	currentHead := Machine{ head: newElement, currentHead: newElement, loop: true }
	return currentHead
}

func (m *Machine) choice(s string){
	switch s {
	case "a":
		m.left()
	case "d":
		m.right()
	case "w":
		m.change()
	case "h":
		m.printHelp()
	case "p":
		m.printAll()
	case "q":
		m.quit()
	default:
		m.printAll()
	}
}

func (m *Machine) left(){ 
	if m.currentHead.prev == nil{
		m.printAll()
		return
	}
	m.currentHead = m.currentHead.prev 
	m.printAll()
}

func (m *Machine) right(){ 
	if m.currentHead.next == nil{
		m.printAll()
		return
	}
	m.currentHead = m.currentHead.next 
	m.printAll()
}


func (m *Machine) change() {
	if m.currentHead.value == true { 
		m.currentHead.value = false
	}else{
		m.currentHead.value = true
	}
	m.printAll()
}

func (m Machine) printHelp(){ fmt.Println("\033[37m", "   *** HELP ***\n\n  h -> help\n  a -> left\n  d -> right\n  w -> change\n  p -> print all\n  q -> QUIT") }

func (m Machine) printAll() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	for current := m.head; current != nil;{
		if m.currentHead == current{
			fmt.Print("\033[31m", formatBoolInInt(current.value), "\033[m")
		}else{
			fmt.Print(formatBoolInInt(current.value))
		}
		fmt.Print(" - ")
		current = current.next
	}
	fmt.Println("nil")
}

func formatBoolInInt(b bool) int {
	if b{
		return 1
	}
	return 0
}

func (m *Machine) quit() {
	m.loop = false
	m.printAll()
	fmt.Println("\nFIM ....")
}