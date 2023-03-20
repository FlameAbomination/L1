package main

import (
	"fmt"
)

type Human struct {
	height float32
	weight float32
}

// В структуру можно встроить как другую структуру,
// так и её указатель
type Action struct {
	Human
}

type ActionRef struct {
	*Human
}

func (h Human) GetBMI() float32 {
	return float32(h.weight) / (h.height * h.height)
}

func task1() {
	human := Human{
		1.84,
		60,
	}
	action := Action{
		human,
	}
	actionRef := Action{
		human,
	}

	fmt.Println(action.GetBMI())
	fmt.Println(actionRef.GetBMI())
}
