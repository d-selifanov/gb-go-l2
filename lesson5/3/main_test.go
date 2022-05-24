package main

import (
	"math/rand"
	"sync"
	"testing"
)

type Set struct {
	sync.Mutex
	mm map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		mm: map[int]struct{}{},
	}
}

func (s *Set) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *Set) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

type RWSet struct {
	sync.RWMutex
	mm map[int]struct{}
}

func NewRWSet() *RWSet {
	return &RWSet{
		mm: map[int]struct{}{},
	}
}

func (s *RWSet) Add(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
}

func (s *RWSet) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
}

// WRITE 10%, READ 90%
func BenchmarkMutex_w10_r90(b *testing.B) {
	var set = NewSet()
	b.Run("Bench Mutex write 10% read 90%", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() <= 0.1 {
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
		})
	})
}
func BenchmarkRWMutex_w10_r90(b *testing.B) {
	var set = NewRWSet()
	b.Run("Bench RWMutex write 10% read 90%", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() <= 0.1 {
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
		})
	})
}

// WRITE 50%, READ 50%
func BenchmarkMutex_w50_r50(b *testing.B) {
	var set = NewSet()
	b.Run("Bench Mutex write 50% read 50%", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() <= 0.5 {
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
		})
	})
}
func BenchmarkRWMutex_w50_r50(b *testing.B) {
	var set = NewRWSet()
	b.Run("Bench RWMutex write 50% read 50%", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() <= 0.5 {
					set.Add(1)
				} else {
					set.Has(1)
				}
			}
		})
	})
}

// WRITE 90%, READ 10%
func BenchmarkMutex_w90_r10(b *testing.B) {
	var set = NewSet()
	b.Run("Bench Mutex write 90% read 10%", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() <= 0.1 {
					set.Has(1)
				} else {
					set.Add(1)
				}
			}
		})
	})
}

func BenchmarkRWMutex_w90_r10(b *testing.B) {
	var set = NewRWSet()
	b.Run("Bench RWMutex write 90% read 10%", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if rand.Float32() <= 0.1 {
					set.Has(1)
				} else {
					set.Add(1)
				}
			}
		})
	})
}
