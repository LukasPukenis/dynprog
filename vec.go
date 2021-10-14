package main

type vec2 struct {
	x, y int
}

func (v vec2) add(o vec2) vec2 {
	return vec2{
		x: v.x + o.x,
		y: v.y + o.y,
	}
}

func (v vec2) clamp(min, max vec2) vec2 {
	if v.x < min.x {
		v.x = min.x
	}

	if v.y < min.y {
		v.y = min.y
	}

	if v.x > max.x {
		v.x = max.x
	}

	if v.y > max.y {
		v.y = max.y
	}

	return v
}

func (v vec2) equals(o vec2) bool {
	return v.x == o.x && v.y == o.y
}
