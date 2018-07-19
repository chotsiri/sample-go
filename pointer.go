package main

import "fmt"

type human struct {
	name    string
	age     int
	isAdult bool // Zero Value คือ false
}

// setAdult รับ human เข้ามา
// หากอายุเกิน 18 isAdult จะมีค่าเป็น true
// นอกนั้นมีค่าเป็น false
func setAdult(h *human) {
	h.isAdult = h.age >= 18
}

func main() {
	somchai := human{ name: "Somchai", age: 23 }
	setAdult(&somchai)
	fmt.Println(somchai)

	aom := human{ name: "Sorasak Chot", age: 19 }
	setAdult(&aom)
	fmt.Println(aom)


}
