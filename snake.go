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

func (s *snake) locationUpdate(x, y int) {
	s.x = x
	s.y = y
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
	return s.y == 0 || s.y == len(area) - 1 || s.x == 0 || s.x == len(area[0]) - 1
}

// Snake's Body's Queue
func (s *snake) bodyPush(x, y int) {
	s.body = append(s.body, []int{x, y})
}

func (s *snake) bodyPushFront(x, y int) {
	s.body = append([][]int{{x, y}}, s.body[:]...)
}
func (s *snake) bodyPop() {
	s.body = s.body[1:]
}

func (s *snake) lenInr() {
	s.length++
}

func (s *snake) pointsInr() {
	s.points++
}