package main

import "time"

type snake struct {

	x int
	y int
	length int
	points int

	duration time.Duration

}

func (s *snake) locationUpdate(x, y int) {
	s.x = x
	s.y = y
}

func (s *snake) lengthUpdate(length int) {
	s.length = length
}

func (s *snake) durationUpdate(duration time.Duration) {
	s.duration = duration
}