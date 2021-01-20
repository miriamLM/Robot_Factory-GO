package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot interface {
	PowerOn() *Terminator
	Rename(*Terminator)
	checkName([]*Terminator, *Terminator, int) bool
	Reset([]*Terminator, *Terminator, int)
}

type Terminator struct {
	Name string
}

func (t Terminator) PowerOn() *Terminator {
	return new(Terminator)
}

func (t Terminator) Rename(terminator *Terminator) {
	var name = RandomString(1) + RandomInt(1)
	terminator.Name = name
}

func RandomInt(n int) string {
	var letter = []rune("0123456789")
	b := make([]rune, n)

	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)

}
func RandomString(n int) string {
	var letter = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)

	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}

func (t Terminator) checkName(s []*Terminator, terminator *Terminator, n int) bool {
	var same = false

	for i := 0; i < len(s); i++ {

		if terminator.Name == s[i].Name {
			same = true
			if n == 2 {
				fmt.Println("el robot " + terminator.Name + " ya existe")
			}

		}
	}
	return same
}

func (t Terminator) Reset(s []*Terminator, terminator *Terminator, random int) {
	fmt.Println("Reset robot " + s[random].Name + " por " + terminator.Name)
	s[random].Name = terminator.Name
}

func infiniteLoop(r Robot) {
	var s []*Terminator

	for true {

		if len(s) < 100 {
			var t = r.PowerOn()
			t.Rename(t)
			var same = r.checkName(s, t, 1)

			for same == true {
				t.Rename(t)
				same = r.checkName(s, t, 1)
			}

			s = append(s, t)

		} else {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10) // n will be between 0 and 10
			time.Sleep(time.Duration(n) * time.Second)
			var t = r.PowerOn()
			var num = rand.Intn(100)
			t.Rename(t)
			var same = r.checkName(s, t, 2)

			for same == true {
				t.Rename(t)
				same = r.checkName(s, t, 2)
			}

			r.Reset(s, t, num)

		}
	}
}

func main() {
	t1000 := &Terminator{}
	infiniteLoop(t1000)
}
