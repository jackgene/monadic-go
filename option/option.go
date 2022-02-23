package option

import "fmt"

type Option[T any] struct {
	defined bool
	item    T
}

func None[T any]() Option[T] {
	return Option[T]{defined: false}
}

func Wrap[T any](in T) Option[T] {
	return Option[T]{
		defined: true,
		item:    in,
	}
}

func UnwrapOrElse[T any](in Option[T], orElse T) T {
	if !in.defined {
		return orElse
	}

	return in.item
}

func Map[T any, U any](in Option[T], transform func(T) U) Option[U] {
	if !in.defined {
		return Option[U]{defined: false}
	}

	out := transform(in.item)

	return Option[U]{defined: true, item: out}
}

func AndThen[T any, U any](in Option[T], transform func(T) Option[U]) Option[U] {
	if !in.defined {
		return Option[U]{defined: false}
	}

	return transform(in.item)
}

func Filter[T any](in Option[T], predicate func(T) bool) Option[T] {
	if in.defined && !predicate(in.item) {
		in.defined = false
	}

	return in
}

func (o Option[T]) String() string {
	if !o.defined {
		return "None"
	}

	return fmt.Sprintf("Some(%v)", o.item)
}