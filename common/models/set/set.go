package set

type Set[T comparable] struct {
	items map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

func (s *Set[T]) Delete(item T) {
	delete(s.items, item)
}

func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set[T]) GetAll() []T {
	items := make([]T, 0, len(s.items))
	for item := range s.items {
		items = append(items, item)
	}
	return items
}

func (s *Set[T]) Length() int {
	return len(s.items)
}
