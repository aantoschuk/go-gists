package bitSet

import (
	"fmt"
	"strings"
)

type BitSet struct {
	data []byte
}

// Allocate BitSet object.
func NewBitset(n int) *BitSet {
	// (n+7)/8 is equivalent to ceil(n/8) but avoids floating-point arithmetic
	// by performing integer-only ceiling division.
	return &BitSet{data: make([]byte, (n+7)/8)}
}

// Add function ads ( flip element ) to the BitSet
// return true if operation is successful.
func (b *BitSet) Add(v int) bool {
	if v < 0 {
		return false
	}
	index := v / 8

	if index >= len(b.data) {
		return false
	}
	offset := v % 8
	// <<  = pick a bit position (mask)
	// |= switch 0 to 1
	b.data[index] |= (1 << offset)
	return true
}

func (b *BitSet) Delete(v int) bool {
	if v < 0 {
		return false
	}
	index := v / 8

	if index >= len(b.data) {
		return false
	}
	offset := v % 8
	// &^= switch 1 to 0
	b.data[index] &^= (1 << offset)
	return true
}

// Check function checks bits at position
// and returns true if value is present in the set
func (b BitSet) Check(v int) bool {
	if v < 0 {
		return false
	}

	index := v / 8
	if index >= len(b.data) {
		return false
	}
	offset := v % 8
	return (b.data[index] & (1 << offset)) != 0
}

func (b *BitSet) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(b.data); i++ {
		fmt.Fprintf(&sb, "%08b", b.data[i])
	}
	return sb.String()
}

// Union function returns union of both BitSets
func (a *BitSet) Union(b *BitSet) BitSet {
	// calculate capacity. Which is bytes * 8, as 1 byte is 8 bits.
	// So that way we could properly create a new bitset.
	capacity := max(len(a.data), len(b.data))

	bs := *NewBitset(capacity * 8)

	for i := range capacity {
		var x byte

		// apply to temp variable current value.
		// If one of them have 1 and other 0, it should remain 1.
		if i < len(a.data) {
			x |= a.data[i]
		}

		if i < len(b.data) {
			x |= b.data[i]
		}
		// set to the final result
		bs.data[i] = x

	}
	return bs
}

// Intersection returns a new bitset with common values between A and B
func (a *BitSet) Intersection(b *BitSet) BitSet {
	capacity := max(len(a.data), len(b.data))
	bs := *NewBitset(capacity * 8)

	for i := range capacity {
		var av, bv byte

		if i < len(a.data) {
			av = a.data[i]
		}
		if i < len(b.data) {
			bv = b.data[i]
		}

		// if A 1 and B 1 = 1
		// otherwise would results in 0
		bs.data[i] = av & bv
	}

	return bs
}

// Difference function returns difference of both BitSets
func (a *BitSet) Difference(b *BitSet) BitSet {
	capacity := len(a.data)
	bs := *NewBitset(capacity * 8)

	for i := range capacity {
		var av, bv byte

		if i < len(a.data) {
			av = a.data[i]
		}
		if i < len(b.data) {
			bv = b.data[i]
		}
		// &^ operation is difference, so check the difference between two bitsets
		// if a bitset is is 1 and b is 0 = result is 1. if a is 1 and b 1 = results is 0.
		// and if a 0 and b 1 result is 0
		bs.data[i] = av &^ bv
	}
	return bs
}

// Symetric difference is when both A and B contain non-overlapping values, otherwise return 0
func (a *BitSet) SymmetricDiff(b *BitSet) BitSet {

	capacity := max(len(a.data), len(b.data))
	bs := *NewBitset(capacity * 8)

	for i := range capacity {
		var av, bv byte

		if i < len(a.data) {
			av = a.data[i]
		}
		if i < len(b.data) {
			bv = b.data[i]
		}

		// if both A and B has the same balue = remove
		// otherwise keep
		bs.data[i] = av ^ bv
	}

	return bs
}
