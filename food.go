package main

type food struct {
	x int
	y int
}

func (f *food) init() {

	f.x = random.Intn(pty.width-2) + 1
	f.y = random.Intn(pty.height-2) + 1

}
