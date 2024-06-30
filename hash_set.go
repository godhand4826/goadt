package main

var _ Set[string] = (*HashSet[string, string])(nil)

type HashSet[E any, H comparable] struct {
	equalFn     EqualFn[E]
	hashFn      HashFn[E, H]
	listFactory List[E]
	values      map[H]List[E]
	size        int
}

func NewHashSet[E any, H comparable](
	equalFn EqualFn[E],
	hashFn HashFn[E, H],
	listFactory List[E],
) *HashSet[E, H] {
	return &HashSet[E, H]{
		equalFn:     equalFn,
		hashFn:      hashFn,
		listFactory: listFactory,
		values:      make(map[H]List[E]),
	}
}

func (h *HashSet[E, H]) IsEmpty() bool {
	return h.Size() == 0
}

func (h *HashSet[E, H]) Size() int {
	return h.size
}

func (h *HashSet[E, H]) Clear() {
	clear(h.values)
	h.size = 0
}

func (h *HashSet[E, H]) Add(element E) {
	hash := h.hashFn(element)
	if list, ok := h.values[hash]; ok {
		if list.Contains(element) {
			return
		}
		list.Append(element)
		h.size++
	} else {
		list := h.listFactory.NewList()
		list.Append(element)
		h.values[hash] = list
		h.size++
	}
}

func (h *HashSet[E, H]) Remove(element E) bool {
	hash := h.hashFn(element)
	if list, ok := h.values[hash]; ok {
		if ok := list.Remove(element); ok {
			h.size--
			return true
		}
	}
	return false
}

func (h *HashSet[E, H]) Contains(element E) bool {
	hash := h.hashFn(element)
	if list, ok := h.values[hash]; ok {
		return list.Contains(element)
	}
	return false
}
