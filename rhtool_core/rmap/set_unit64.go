package rmap

import "strconv"

type uint64set struct {
	m map[uint64]struct{}
}

func NewUint64Set() *uint64set {
	return &uint64set{m: make(map[uint64]struct{})}
}

func (s *uint64set) Add(item uint64) {
	s.m[item] = struct{}{}
}

func (s *uint64set) AddList(items []uint64) {
	for _, item := range items {
		s.m[item] = struct{}{}
	}
}

func (s *uint64set) Size() int {
	return len(s.m)
}

func (s *uint64set) List() []uint64 {
	v := make([]uint64, 0, s.Size())
	for item := range s.m {
		v = append(v, item)
	}
	return v
}

func (s *uint64set) In(k uint64) bool {
	_, ok := s.m[k]
	return ok
}

func (s *uint64set) ListIn(ks []uint64) bool {
	for _, val := range ks {
		_, ok := s.m[val]
		if !ok {
			return false
		}
	}
	return true
}

func (s *uint64set) Delete(k uint64) {
	delete(s.m, k)
}

func (s *uint64set) DeleteList(ks []uint64) {
	for _, k := range ks {
		delete(s.m, k)
	}
}

func (s *uint64set) Merge2String(split string) (str string) {
	for item := range s.m {
		tmp := strconv.FormatUint(item, 10)
		if str == "" {
			str = tmp
		} else {
			str = str + split + tmp
		}
	}
	return
}
