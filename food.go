package main

type food struct {
	x int
	y int
}

func (f *food) init() {

	f.x = random.Intn(pty.width-4) + 2
	f.y = random.Intn(pty.height-4) + 2

}
