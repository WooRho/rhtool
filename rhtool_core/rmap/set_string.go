package rmap

type stringSet struct {
	m map[string]struct{}
}

func NewStringSet() *stringSet {
	return &stringSet{m: make(map[string]struct{})}
}

func (s *stringSet) Add(item string) {
	s.m[item] = struct{}{}
}

func (s *stringSet) AddList(items []string) {
	for _, item := range items {
		s.m[item] = struct{}{}
	}
}

func (s *stringSet) Size() int {
	return len(s.m)
}

func (s *stringSet) List() []string {
	v := make([]string, 0, s.Size())
	for item := range s.m {
		v = append(v, item)
	}
	return v
}

func (s *stringSet) In(k string) bool {
	_, ok := s.m[k]
	return ok
}

func (s *stringSet) ListIn(ks []string) bool {
	for _, val := range ks {
		_, ok := s.m[val]
		if !ok {
			return false
		}
	}
	return true
}

func (s *stringSet) Delete(k string) {
	delete(s.m, k)
}

func (s *stringSet) DeleteList(ks []string) {
	for _, k := range ks {
		delete(s.m, k)
	}
}

func (s *stringSet) Merge2String(split string) (str string) {
	for item := range s.m {
		if str == "" {
			str = item
		} else {
			str = str + split + item
		}
	}
	return
}
