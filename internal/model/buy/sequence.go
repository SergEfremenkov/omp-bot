package buy

type Sequence struct {
	value uint64
}

func (s *Sequence) NextVal() uint64 {
	s.value++
	return s.value
}
