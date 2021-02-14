package bloom

import (
	// "fmt"

	"github.com/willf/bitset"
)

const defaultSize = 2 << 24

var seeds = []uint{7, 11, 13, 31, 37, 61}

type BloomFilter struct {
	Set   *bitset.BitSet
	Funcs [6]simpleHash
}

func NewBloomFilter() *BloomFilter {
	bf := new(BloomFilter)
	for i := 0; i < len(bf.Funcs); i++ {
		bf.Funcs[i] = simpleHash{defaultSize, seeds[i]}
	}
	bf.Set = bitset.New(defaultSize)
	return bf
}

func (bf *BloomFilter) Add(value string) {
	for _, f := range bf.Funcs {
		bf.Set.Set(f.hash(value))
	}
}

func (bf *BloomFilter) Contains(value string) bool {
	if value == "" {
		return false
	}
	ret := true
	for _, f := range bf.Funcs {
		ret = ret && bf.Set.Test(f.hash(value))
	}
	return ret
}

type simpleHash struct {
	Cap  uint
	Seed uint
}

func (s *simpleHash) hash(value string) uint {
	var result uint = 0
	for i := 0; i < len(value); i++ {
		result = result*s.Seed + uint(value[i])
	}
	return (s.Cap - 1) & result
}
