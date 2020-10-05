package binaryheap

import (
	"reflect"
	"testing"
)

type testStruct struct {
	v int
}

var minCompareFunc = func(A, B interface{}) int {
	a := A.(int)
	b := B.(int)

	if a < b {
		return 1
	}
	if a > b {
		return -1
	}
	return 0
}

var structCompareFunc = func(A, B interface{}) int {
	a := A.(testStruct)
	b := B.(testStruct)

	if a.v < b.v {
		return -1
	}
	if a.v > b.v {
		return 1
	}
	return 0
}

func TestBHInsert(t *testing.T) {
	h := newTestBH()
	got := h.ToSlice()
	exp := []interface{}{3, 1, 2, 0, -2, -3, -1}

	if !reflect.DeepEqual(got, exp) {
		t.Errorf("invalid heap: expected %v, got %v ", exp, got)
	}
}

func TestBHShift(t *testing.T) {
	h := newTestBH()
	exp := []interface{}{3, 2, 1, 0, -1, -2, -3}
	got := func() []interface{} {
		n := h.Len()
		values := []interface{}{}
		for i := 0; i < n; i++ {
			values = append(values, h.Shift())
		}
		return values
	}()

	if !reflect.DeepEqual(got, exp) {
		t.Errorf("shift: expected %v, got %v ", exp, got)
	}
}

func TestBHMinCompareFunc(t *testing.T) {
	h := New().SetCompareFunc(minCompareFunc).InsertMany(0, 1, -3, 2, -2, 3, -1)
	exp := []interface{}{-3, -2, -1, 2, 1, 3, 0}
	got := h.ToSlice()

	if !reflect.DeepEqual(got, exp) {
		t.Errorf("minCompareFunc: expected %v, got %v ", exp, got)
	}
}

func TestBHStructCompareFunc(t *testing.T) {
	values := []interface{}{testStruct{0}, testStruct{1}, testStruct{-3},
		testStruct{2}, testStruct{-2}, testStruct{3}, testStruct{-1}}
	h := New().SetCompareFunc(structCompareFunc).InsertMany(values...)
	exp := []interface{}{testStruct{3}, testStruct{1}, testStruct{2},
		testStruct{0}, testStruct{-2}, testStruct{-3}, testStruct{-1}}
	got := h.ToSlice()

	if !reflect.DeepEqual(got, exp) {
		t.Errorf("minCompareFunc: expected %v, got %v ", exp, got)
	}
}

func newTestBH() Interface {
	h := New()
	h.InsertMany(0, 1, -3, 2, -2).Insert(3).InsertMany(-1)
	return h
}
