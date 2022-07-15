package main

import "fmt"

func Act(h Human) {
	h.Talk()
}

type Human interface {
	Talk()
}

type PoliceOfficer struct{}

func (p *PoliceOfficer) Talk() {
	fmt.Printf("逮捕だ！！")
}

type Firefighter struct{}

func (f *Firefighter) Talk() {
	fmt.Printf("火事だ！！")
}

func main() {
	p := &PoliceOfficer{}
	Act(p)

	f := &Firefighter{}
	Act(f)
}
