package state

type set[T comparable] struct {
	items map[T]struct{}
}

func initialize[T comparable]() *set[T] {
	return &set[T]{
		items: make(map[T]struct{}),
	}
}

func (s *set[T]) add(item T) {
	s.items[item] = struct{}{}
}

func (s *set[T]) delete(item T) {
	delete(s.items, item)
}

func (s *set[T]) contains(item T) bool {
	_, ok := s.items[item]
	return ok
}
