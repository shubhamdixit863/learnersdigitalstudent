package services
import "fmt"

type Node struct {
	Data string
	Next *Node
	Prev *Node
}

type VersionManager struct {
	Current *Node
	Head     *Node
	Tail     *Node
	Length   int
	Capacity int
	StackR   []*Node
}

func NewVersionManager(capacity int) *VersionManager {
	return &VersionManager{
		Current: nil,
		Head:     nil,
		Tail:     nil,
		Length:   0,
		Capacity: capacity,
		StackR:   make([]*Node, 0),
	}
}

// type stack []int

// func (s *stack) Push(v int) {
//     *s = append(*s, v)
// }

// func (s *stack) Pop() int {
//     res:=(*s)[len(*s)-1]
//     *s=(*s)[:len(*s)-1]
//     return res
// }

// Add version
func (vm *VersionManager) AddVersion(version string) {
	newNode := &Node{
		Data: version,
		Next: nil,
		Prev: vm.Tail,
	}

	if vm.Head == nil {
		vm.Head = newNode
		vm.Tail = newNode
		vm.Current=newNode
	} else {
		vm.Current.Next=newNode
		newNode.Prev=vm.Current
		vm.Current=newNode
		vm.Tail = newNode
	}
	vm.Length++

	

	if vm.Length > vm.Capacity {
		vm.Head = vm.Head.Next
		vm.Head.Prev = nil
		vm.Length--
		fmt.Println("oldest version removed due to capacity limit")
	}
  vm.StackR=make([]*Node, 0)
	fmt.Println("version added",version)
}


func (vm *VersionManager) Undo()string{
	if vm.Current==nil || vm.Current.Prev==nil{
		return "no previous versions to undo"
	}
	vm.StackR=append(vm.StackR, vm.Current)
	vm.Current=vm.Current.Prev
	return vm.Current.Data
}

func (vm *VersionManager) Redo() string{
	if len(vm.StackR)==0{
		return "no redo available"
	}
	vm.Current=vm.StackR[len(vm.StackR)-1]
	vm.StackR=vm.StackR[:len(vm.StackR)-1]
	return vm.Current.Data

}

func (vm *VersionManager) GetCurrentVersion() string{
	if(vm.Current==nil){
		return "no version available"
}
 return vm.Current.Data
}

