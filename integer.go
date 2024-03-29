package types

import "golang.org/x/exp/constraints"

type Integer[T constraints.Integer] struct {
	basic[T]
}

func NewInteger[T constraints.Integer](i T) Integer[T] {
	return Integer[T]{
		newBasic(i),
	}
}

func (i Integer[T]) Int() int {
	return int(i.Value())
}

func (i Integer[T]) Int8() int8 {
	return int8(i.Value())
}

func (i Integer[T]) Int16() int16 {
	return int16(i.Value())
}

func (i Integer[T]) Int32() int32 {
	return int32(i.Value())
}

func (i Integer[T]) Int64() int64 {
	return int64(i.Value())
}

func (i Integer[T]) UInt8() uint8 {
	return uint8(i.Value())
}

func (i Integer[T]) UInt16() uint16 {
	return uint16(i.Value())
}

func (i Integer[T]) UInt32() uint32 {
	return uint32(i.Value())
}

func (i Integer[T]) UInt64() uint64 {
	return uint64(i.Value())
}

func (i Integer[T]) IsZero() bool {
	return i.Value() == 0
}

func (i Integer[T]) IsNegative() bool {
	return i.Value() < 0
}

func (i Integer[T]) IsPositive() bool {
	return i.Value() > 0
}

func (i Integer[T]) Str() Str {
	return fromStringer(i)
}

func (i Integer[T]) Float() Float[float64] {
	return NewFloat(float64(i.Value()))
}
