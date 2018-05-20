package game

type FixedSizeQueue struct {
	objects []*GameState
	size    int
	head    int
	count   int
}

func NewStateQueue(fixedSize int) *FixedSizeQueue {
	return &FixedSizeQueue{
		objects: make([]*GameState, fixedSize),
		size:    fixedSize,
	}
}

func (q *FixedSizeQueue) Push(state *GameState) {
	for q.count > q.size {
		q.head++
		q.count--
	}
	q.objects[(q.head+q.count)%q.size] = state
}

func (q *FixedSizeQueue) GetItem(depth int) *GameState {
	if depth < 0 {
		return nil
	}
	if depth >= q.count {
		return q.objects[q.head]
	}
	return q.objects[q.head+q.count-depth]
}
