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
	t.lock = sync.Mutex{}
	t.cells = make([]byte, blockSize)
	return t
}

func (t *Tape) Left() {
  t.index = t.index++ % t.size
}

func (t *Tape) Right() {
  if t.index == 0
    t.index = t.size
  t.index--
}

func (t *Tape) Read() (data byte) {
  return t.cells[t.index]
}
