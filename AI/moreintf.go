package main

import "fmt"

type Lesson interface {
	numbers() int
	letters() string
}

type Math struct{
	number int
}
type Eng struct{
	letter string
}
type Trksh struct{
	letter string
}
type Science struct{
	number int
}

func(m *Math) numbers() int{
	return m.number*15
}
func (e *Eng) letters() string{
	return e.letter + "fdlgndfg"
}
func (t *Trksh) letters() string{
	return t.letter + "ffgdbgdfbdf"
}
func(s *Science) numbers() int{
	return s.number*19
}
func main() {
	m1 := &Math{}
	e1 := &Eng{}
	t1 := &Trksh{}
	s1 := &Science{}
	
	fmt.Println(m1,e1,t1,s1)
}