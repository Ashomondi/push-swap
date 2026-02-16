package ops

import "push-swap/internal/stack"

type Executor struct {
	A *stack.Stack
	B *stack.Stack

	// If Record=true, we store operations here (push-swap mode)
	Record bool
	Out    []string
}

func NewExecutor(a, b *stack.Stack, record bool) *Executor {
	return &Executor{A: a, B: b, Record: record}
}

func (e *Executor) emit(op string) {
	if e.Record {
		e.Out = append(e.Out, op)
	}
}

// ---- Swap ----
func (e *Executor) Sa() {
	if e.A.Len() < 2 {
		return
	}
	e.A.Data[0], e.A.Data[1] = e.A.Data[1], e.A.Data[0]
	e.emit("sa")
}

func (e *Executor) Sb() {
	if e.B.Len() < 2 {
		return
	}
	e.B.Data[0], e.B.Data[1] = e.B.Data[1], e.B.Data[0]
	e.emit("sb")
}

func (e *Executor) Ss() {
	if e.A.Len() >= 2 {
		e.A.Data[0], e.A.Data[1] = e.A.Data[1], e.A.Data[0]
	}
	if e.B.Len() >= 2 {
		e.B.Data[0], e.B.Data[1] = e.B.Data[1], e.B.Data[0]
	}
	e.emit("ss")
}

// ---- Push ----
func (e *Executor) Pa() {
	v, ok := e.B.PopTop()
	if !ok {
		return
	}
	e.A.PushTop(v)
	e.emit("pa")
}

func (e *Executor) Pb() {
	v, ok := e.A.PopTop()
	if !ok {
		return
	}
	e.B.PushTop(v)
	e.emit("pb")
}

// ---- Rotate ----
func rotateUp(s *stack.Stack) {
	if s.Len() < 2 {
		return
	}
	first := s.Data[0]
	s.Data = append(s.Data[1:], first)
}

func rotateDown(s *stack.Stack) {
	if s.Len() < 2 {
		return
	}
	last := s.Data[len(s.Data)-1]
	s.Data = append([]int{last}, s.Data[:len(s.Data)-1]...)
}

func (e *Executor) Ra() {
	rotateUp(e.A)
	e.emit("ra")
}

func (e *Executor) Rb() {
	rotateUp(e.B)
	e.emit("rb")
}

func (e *Executor) Rr() {
	rotateUp(e.A)
	rotateUp(e.B)
	e.emit("rr")
}

func (e *Executor) Rra() {
	rotateDown(e.A)
	e.emit("rra")
}

func (e *Executor) Rrb() {
	rotateDown(e.B)
	e.emit("rrb")
}

func (e *Executor) Rrr() {
	rotateDown(e.A)
	rotateDown(e.B)
	e.emit("rrr")
}
func (e *Executor) Apply(op string) bool {
	switch op {
	case "sa":
		e.Sa()
	case "sb":
		e.Sb()
	case "ss":
		e.Ss()
	case "pa":
		e.Pa()
	case "pb":
		e.Pb()
	case "ra":
		e.Ra()
	case "rb":
		e.Rb()
	case "rr":
		e.Rr()
	case "rra":
		e.Rra()
	case "rrb":
		e.Rrb()
	case "rrr":
		e.Rrr()
	default:
		return false
	}
	return true
}
