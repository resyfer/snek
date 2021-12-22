package main

type food struct {
	x int
	y int
}

func (f *food) init() {

	f.x = random.Intn(WIDTH)
	f.y = random.Intn(HEIGHT)

}
