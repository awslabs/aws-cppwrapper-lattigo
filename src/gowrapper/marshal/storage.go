// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Adapted from https://github.com/justinfx/gofileseq/blob/97d877d68cdef5939f4db199f52e4424fe9a691a/exp/cpp/export/storage.go,
// commit 97d877d68cdef5939f4db199f52e4424fe9a691a, MIT License.

package marshal

import "C"

import (
	"errors"
	"strconv"
	"sync"
	"sync/atomic"
	"unsafe"
)

type Handle = uint64

var CrossLangObjMap *xlangRefMap = &xlangRefMap{
	lock: new(sync.RWMutex),
	m:    make(map[Handle]*xlangRef),
	// in C++, the default constructor uses a handle value of 0,
	// which should never point to a real object for safety.
	// Start the counter here at 1.
	idCtr: 1,
}

type xlangRefMap struct {
	lock  *sync.RWMutex
	m     map[Handle]*xlangRef
	idCtr uint64
}

type xlangRef struct {
	Ptr  unsafe.Pointer
	refs uint32
}

func (m *xlangRefMap) Len() int {
	m.lock.RLock()
	l := len(m.m)
	m.lock.RUnlock()
	return l
}

func (m *xlangRefMap) Add(fset unsafe.Pointer) Handle {
	m.lock.Lock()
	m.idCtr = m.idCtr + 1
	id := m.idCtr
	m.m[id] = &xlangRef{fset, 1}
	m.lock.Unlock()
	return id
}

//export incref
func incref(handle Handle) {
	CrossLangObjMap.Incref(handle)
}

func (m *xlangRefMap) Incref(id Handle) {
	m.lock.RLock()
	ref, ok := m.m[id]
	m.lock.RUnlock()

	if !ok {
		panic(errors.New("Cannot find object for specified handle: " + strconv.FormatUint(id, 10)))
	}

	atomic.AddUint32(&ref.refs, 1)
}

//export decref
func decref(handle Handle) {
	CrossLangObjMap.Decref(handle)
}

//export refCount
func refCount(handle Handle) uint32 {
	return CrossLangObjMap.RefCount(handle)
}

func (m *xlangRefMap) RefCount(id Handle) uint32 {
	m.lock.RLock()
	ref, ok := m.m[id]
	m.lock.RUnlock()

	if !ok {
		panic(errors.New("Cannot find object for specified handle: " + strconv.FormatUint(id, 10)))
	}

	return atomic.LoadUint32(&ref.refs)
}

func (m *xlangRefMap) Decref(id Handle) {
	m.lock.RLock()
	ref, ok := m.m[id]
	m.lock.RUnlock()

	if !ok {
		panic(errors.New("Cannot find object for specified handle: " + strconv.FormatUint(id, 10)))
	}

	refs := atomic.AddUint32(&ref.refs, ^uint32(0))
	if refs != 0 {
		return
	}

	m.lock.Lock()
	if atomic.LoadUint32(&ref.refs) == 0 {
		delete(m.m, id)
	}
	m.lock.Unlock()
}

func (m *xlangRefMap) Get(id Handle) *xlangRef {
	m.lock.RLock()
	ref, ok := m.m[id]
	m.lock.RUnlock()

	if !ok {
		panic(errors.New("Fatal error: unable to retrieve object for specified handle: " + strconv.FormatUint(id, 10)))
	}

	return ref
}
