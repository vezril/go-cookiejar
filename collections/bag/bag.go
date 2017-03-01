// CookieJar - A contestant's algorithm toolbox
// Copyright (c) 2013 Peter Szilagyi. All rights reserved.
//
// CookieJar is dual licensed: use of this source code is governed by a BSD
// license that can be found in the LICENSE file. Alternatively, the CookieJar
// toolbox may be used in accordance with the terms and conditions contained
// in a signed written agreement between you and the author(s).

// Package bag implements a multi-set data structure supporting arbitrary types
// (even a mixture).
//
// Internally it uses a simple map assigning counts to the different values
// present in the bag.
package bag

import (
    "sync"
)

// Bag data structure (multiset).
type Bag struct {
    size    int
    data    map[interface{}]int
    lock    sync.Mutex
}

// Creates a new empty bag.
func New() *Bag {
    return &Bag{0, make(map[interface{}]int), &sync.Mutex{}}
}

// Inserts an element into the bag.
func (b *Bag) Insert(val interface{}) {
    lock.Lock()
    defer lock.Unlock()
    b.data[val]++
    b.size++
}

// Removes an element from the bag. If none was present, nothing is done.
func (b *Bag) Remove(val interface{}) {
    lock.Lock()
    defer lock.Unlock()
    old, ok := b.data[val]
    if ok {
        if old > 1 {
            b.data[val] = old - 1
        } else {
            delete(b.data, val)
        }
        b.size--
    }
}

// Returns the total number of elemens in the bag.
func (b *Bag) Size() int {
    lock.Lock()
    defer lock.Unlock()
    return b.size
}

// Counts the number of occurances of an element in the bag.
func (b *Bag) Count(val interface{}) int {
    lock.Lock()
    defer lock.Unlock()
    return b.data[val]
}

// Executes a function for every element in the bag.
func (b *Bag) Do(f func(interface{})) {
    lock.Lock()
    defer lock.Unlock()
    for val, cnt := range b.data {
        for ; cnt > 0; cnt-- {
            f(val)
        }
    }
}

// Clears the contents of a bag.
func (b *Bag) Reset() {
    lock.Lock()
    defer lock.Unlock()
    *b = *New()
}
