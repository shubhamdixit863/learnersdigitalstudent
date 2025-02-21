package main

func Calculator(op string, a int, b int) int{
	if op == "Add" {
		return a + b
	}
	if op == "Sub" {
		return a - b
	}
	if op == "Mult" {
		return a * b
	}
	if op == "Div" {
		return a / b
	}
	return 0
}
