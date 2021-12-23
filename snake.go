package main

import "time"

// Movemnent Direction
type direction struct {
	down  int
	right int
}

var UP = direction{down: -1, right: 0}
var DOWN = direction{down: 1, right: 0}
var LEFT = direction{down: 0, right: -1}
var RIGHT = direction{down: 0, right: 1}

// Snake Struct
type snake struct {
	x      int // Keep in mind that 0 and n-1 (limits of terminal) have a border on them
	y      int // Keep in mind that 0 and n-1 (limits of terminal) have a border on them
	length int
	points int

	dur time.Duration
	dir direction

	symbol string

	body [][]int
}

func (s *snake) move() {
	s.locationUpdate(s.x + s.dir.right, s.y + s.dir.down)
}

func (s *snake) locationUpdate(x, y int) {
	area[y][x] = 0
	s.x = x
	s.y = y
	area[y][x] = 1
}

func (s *snake) lengthUpdate(length int) {
	s.length = length
}

func (s *snake) durationUpdate(duration time.Duration) {
	s.dur = duration
}

func (s snake) foodCollision() bool {
	return s.x == dinner.x && s.y == dinner.y
}

func (s snake) boundaryCollision() bool {
	return s.y + s.dir.down == -1 || s.y + s.dir.down == HEIGHT || s.x + s.dir.right == -1 || s.x + s.dir.right == WIDTH
}

// Snake's Body's Queue
func (s *snake) bodyPush(x, y int) {
	area[y][x] = 3
	s.body = append(s.body, []int{x, y})
}

func (s *snake) bodyPushFront(x, y int) {
	area[y][x] = 1
	s.body = append([][]int{{x, y}}, s.body[:]...)
}

func (s *snake) bodyPop() {
	
	area[s.body[0][1]][s.body[0][0]] = 0
	s.body = s.body[1:]

}

func (s *snake) lenInr() {
	s.length++
}

func (s *snake) pointsInr() {
	s.points++
}