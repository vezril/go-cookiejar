package tape

import (
  "sync"
)

const blockSize = 8192

type Tape struct {
	size		uint
	lock		sync.Mutex
	cells		[]byte
	index		uint
}

func New() *Tape {
  t := new(Tape)
  t.size = blockSize
	t.lock = sync.Mutex{}
	t.cells = make([]byte, blockSize)

	return t
}

func (t *Tape) Right() {
  t.lock.Lock()
  defer t.lock.Unlock()

  t.index = (t.index + 1) % t.size
}

func (t *Tape) Left() {
  t.lock.Lock()
  defer t.lock.Unlock()

  if t.index == 0 {
    t.index = t.size
  }
  t.index--
}

func (t *Tape) returnIndex() uint {
  return t.index
}

func (t *Tape) Read() byte {
  return t.cells[t.index]
}
