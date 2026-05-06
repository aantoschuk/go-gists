package bitSet

import (
	"bytes"
	"testing"
)

func TestBitSet_NewBitSet(t *testing.T) {
	bs := NewBitset(3)

	// 8 bits = 1 byte. So i the size of slice would be 1, as i specify only 3 bits size.
	if len(bs.data) != 1 {
		t.Fatalf("expected size to be 1 bute , but got: %d", len(bs.data))
	}

	if !bytes.Equal(bs.data, []byte{0}) {
		t.Fatalf("expected to have empty slice but got: %v", bs.data)
	}
}

func TestBitSet_Add(t *testing.T) {
	bs := NewBitset(3)
	bs.Add(2)
	r := bs.String()
	if r != "00000100" {
		t.Fatalf("should flip the 3 bit but got %s", r)
	}
}

func TestBitSet_Delete(t *testing.T) {
	bs := NewBitset(3)
	bs.Add(2)
	bs.Delete(2)
	r := bs.String()
	if r != "00000000" {
		t.Fatalf("Delete should flip 3 bit to 0 but got: %s", r)
	}

	bs.Delete(2)
	if r != "00000000" {
		t.Fatalf("test Delete again: should everything remain 0: %s", r)
	}
}

func TestBitSet_Check(t *testing.T) {
	bs := NewBitset(9)
	bs.Add(1)
	ok := bs.Check(1)

	if !ok {
		t.Fatalf("expected 1 to be in the set, but got: %v", ok)
	}

	ok = bs.Check(2)

	if ok {
		t.Fatalf("expected 2 not to be in the set")
	}
}

func TestBitSet_Union(t *testing.T) {
	a := NewBitset(3)
	b := NewBitset(4)
	a.Add(1)
	b.Add(2)

	r := a.Union(b)
	str := r.String()
	if str != "00000110" {
		t.Fatalf("result BitSet should have 1 and 2 but got: %s", str)
	}
}

func TestBitSet_Intersection(t *testing.T) {
	a := NewBitset(3)
	b := NewBitset(4)
	a.Add(1)
	a.Add(1)
	b.Add(1)

	r := a.Intersection(b)
	str := r.String()
	if str != "00000010" {
		t.Fatalf("intersecction result should have 1 but got: %s", str)
	}
}

func TestBitSet_Difference(t *testing.T) {
	a := NewBitset(4)
	b := NewBitset(4)

	a.Add(1)
	a.Add(2)
	a.Add(3)
	b.Add(3)
	b.Add(4)
	b.Add(5)

	r := a.Difference(b)
	str := r.String()
	if str != "00000110" {
		t.Fatalf("result BitSet should have 1 and 2 but got: %s", str)
	}
}

func TestBitSet_SymmertricDiff(t *testing.T) {
	a := NewBitset(3)
	b := NewBitset(4)

	a.Add(1)
	a.Add(2)
	a.Add(3)
	b.Add(3)
	b.Add(4)
	b.Add(5)

	r := a.SymmetricDiff(b)
	str := r.String()
	if str != "00110110" {
		t.Fatalf("result BitSet should have 1 and not 2 but got: %s", str)
	}
}
