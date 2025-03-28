package services

import "fmt"

// we have to implement doubly linked list
type VM struct {
	L         *LinkedList
	Stack     []*LinkedList
	UndoStack []*LinkedList
}
type LinkedList struct {
	Storage  string
	Next     *LinkedList
	Previous *LinkedList
}

func NewVm() *VM {
	return &VM{}
}
func (v *VM) AddVersion(s string) {

	l := LinkedList{s, nil, nil}
	v.L = &l
	if len(v.Stack) <= 5 {
		v.Stack = append(v.Stack, &l)
	} else {
		v.Stack = append(v.Stack[1:], &l)
	}
	fmt.Println("new vwersion added")
	fmt.Println(v.L)
	fmt.Println(v.Stack)
	//fmt.Println(*v.Stack[0])
	v.GetCurrentversion()
}
func (v *VM) GetCurrentversion() {
	fmt.Println(v.L.Storage)
}

func (v *VM) UNDO() {
	if len(v.Stack) > 0 {
		n := v.Stack[len(v.Stack)-1]
		v.L = n
		v.UndoStack = append(v.UndoStack, v.L)
		v.Stack = v.Stack[:len(v.Stack)-2]
		fmt.Println(v.L)
		fmt.Println(v.Stack)
	}
}
func (v *VM) Redo() {
	if len(v.UndoStack) > 0 {
		n := v.UndoStack[len(v.UndoStack)-1]
		v.L = n
		v.Stack = append(v.Stack, n)
		v.UndoStack = v.UndoStack[:len(v.UndoStack)-1]
		v.GetCurrentversion()
	}
}
