// CookieJar - A contestant's algorithm toolbox
// Copyright (c) 2013 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: use of this source code is governed by a BSD
// license that can be found in the LICENSE file. Alternatively, the CookieJar
// toolbox may be used in accordance with the terms and conditions contained
// in a signed written agreement between you and the author(s).

// Package queue implements a FIFO (first in first out) data structure supporting
// arbitrary types (even a mixture).
//
// Internally it uses a dynamically growing circular slice of blocks, resulting
// in faster resizes than a simple dynamic array/slice would allow.
package queue

import (
  "sync"
)

// The size of a block of data
const blockSize = 4096

// First in, first out data structure.
type Queue struct {
	tailIdx int
	headIdx int
	tailOff int
	headOff int
  lock    sync.Mutex

	blocks [][]interface{}
	head   []interface{}
	tail   []interface{}
}

// Creates a new, empty queue.
func New() *Queue {
	result := new(Queue)
  result.lock = sync.Mutex{}
	result.blocks = [][]interface{}{make([]interface{}, blockSize)}
	result.head = result.blocks[0]
	result.tail = result.blocks[0]
	return result
}

// Pushes a new element into the queue, expanding it if necessary.
func (q *Queue) Push(data interface{}) {
  q.lock.Lock()
  defer q.lock.Unlock()

	q.tail[q.tailOff] = data
	q.tailOff++
	if q.tailOff == blockSize {
		q.tailOff = 0
		q.tailIdx = (q.tailIdx + 1) % len(q.blocks)

		// If we wrapped over to the end, insert a new block and update indices
		if q.tailIdx == q.headIdx {
			buffer := make([][]interface{}, len(q.blocks)+1)
			copy(buffer[:q.tailIdx], q.blocks[:q.tailIdx])
			buffer[q.tailIdx] = make([]interface{}, blockSize)
			copy(buffer[q.tailIdx+1:], q.blocks[q.tailIdx:])
			q.blocks = buffer
			q.headIdx++
			q.head = q.blocks[q.headIdx]
		}
		q.tail = q.blocks[q.tailIdx]
	}
}

// Pops out an element from the queue.
func (q *Queue) Pop() (res interface{}) {
  q.lock.Lock()
  defer q.lock.Unlock()

	res, q.head[q.headOff] = q.head[q.headOff], nil
	q.headOff++
	if q.headOff == blockSize {
		q.headOff = 0
		q.headIdx = (q.headIdx + 1) % len(q.blocks)
		q.head = q.blocks[q.headIdx]
	}
	return
}

// Returns the first element in the queue.
func (q *Queue) Front() interface{} {
  q.lock.Lock()
  defer q.lock.Unlock()

	return q.head[q.headOff]
}

// Checks whether the queue is empty.
func (q *Queue) Empty() bool {
  q.lock.Lock()
  defer q.lock.Unlock()
	return q.headIdx == q.tailIdx && q.headOff == q.tailOff
}

// Returns the number of elements in the queue.
func (q *Queue) Size() int {
  q.lock.Lock()
  defer q.lock.Unlock()
	if q.tailIdx > q.headIdx {
		return (q.tailIdx-q.headIdx)*blockSize - q.headOff + q.tailOff
	} else if q.tailIdx < q.headIdx {
		return (len(q.blocks)-q.headIdx+q.tailIdx)*blockSize - q.headOff + q.tailOff
	} else {
		return q.tailOff - q.headOff
	}
}

// Clears out the contents of the queue.
func (q *Queue) Reset() {
  q.lock.Lock()
  defer q.lock.Unlock()
	*q = *New()
}
