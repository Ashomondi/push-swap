package stack

type Stack struct {
	Data []int
}

func New(nums []int) *Stack {
	cp := make([]int, len(nums))
	copy(cp, nums)
	return &Stack{Data: cp}
}

func (s *Stack) Len() int {
	return len(s.Data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}

func (s *Stack) PushTop(v int) {
	s.Data = append([]int{v}, s.Data...)
}

func (s *Stack) PopTop() (int, bool) {
	if len(s.Data) == 0 {
		return 0, false
	}
	v := s.Data[0]
	s.Data = s.Data[1:]
	return v, true
}
