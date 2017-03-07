package tape

import (
  "sync"
)

const blockSize = 32

type Tape struct {
	size		 uint
	lock		 sync.Mutex
	cells		  []byte
	index		  uint
  totalOps  uint
}

func New() *Tape {
  t := new(Tape)
  t.size = blockSize
	t.lock = sync.Mutex{}
	t.cells = make([]byte, blockSize)
  t.totalOps = 0

	return t
}

func (t *Tape) Right() {
  t.lock.Lock()
  defer t.lock.Unlock()

  // Grow slice if needed
  if t.index == t.size-1 {
    t.size = t.size*2
    newCells := make([]byte, t.size)
    copy(newCells, t.cells)
    t.cells = newCells
  }

  t.index++
}

func (t *Tape) Left() {
  t.lock.Lock()
  defer t.lock.Unlock()

  // If index at 0, grow slice to the left and modify pointers accordingly
  if t.index == 0 {
    newCells := make([]byte, t.size * 2)
    copy(newCells[t.size:], t.cells)
    t.index = t.size
    t.size = t.size * 2
    t.cells = newCells
  }
  t.index--
  t.totalOps++
}

// Doesn't validate bounds yet
func (t *Tape) Sync(index uint) {
  t.lock.Lock()
  defer t.lock.Unlock()

  t.index = index
}

func (t *Tape) SyncToOpening() {
  var count uint
  for {
    t.Left()
    token := t.Read()
    if token == 93 {
      count++
    } else if token == 91 {
      if count == 0 {
        break
      } else {
        count--
      }
    }
  }
  t.Left()
}

func (t *Tape) SyncToClosing() {
  for t.Read() != 93 {
    t.Right()
  }
}

func (t *Tape) Read() byte {
  t.lock.Lock()
  defer t.lock.Unlock()
  t.totalOps++

  return t.cells[t.index]
}

func (t *Tape) Write(b byte) {
  t.lock.Lock()
  defer t.lock.Unlock()
  t.totalOps++

  t.cells[t.index] = b
}

func (t *Tape) Inc() {
  t.lock.Lock()
  defer t.lock.Unlock()
  t.totalOps++

  t.cells[t.index]++
}

func (t *Tape) Dec() {
  t.lock.Lock()
  defer t.lock.Unlock()
  t.totalOps++

  t.cells[t.index]--
}

func (t *Tape) GetIndex() uint {
  t.lock.Lock()
  defer t.lock.Unlock()

  return t.index
}

func (t *Tape) GetCells() []byte {
  t.lock.Lock()
  defer t.lock.Unlock()

  return t.cells
}
