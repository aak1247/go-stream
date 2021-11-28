package stream



func Of[T any](data []T) *Stream[T] {
	return &Stream[T]{data}
}


func Zip[S any, T any](streams ...*Stream[S], mapper func(...S) *Stream[T]) (*Stream[T], bool) {
	var args []S
	var res []T
	for _, v := range streams[0].data {
		for j, s := range streams {
			if j >= len(s.data) {
				return nil, false
			}
			cur := s.Take(1)
			args = append(args, s)
			res = append(res, mapper(...args))
		}
	}
	return &Stream[T]{data: res}, true
}


func Map[S any, T any](stream *Stream[S], mapper func(S) T) *Stream[T] {
	var res []T
	for _, v := range stream.data {
		res = append(res, mapper(v))
	}
	return &Stream[T]{data: res}
}

func FlatMap[S any, T any](stream *Stream[S], mapper func(S) *Stream[T]) *Stream[T] {
	var res []T
	for _, v := range stream.data {
		res = append(res, (*mapper(v)).data...)
	}
	return &Stream[T]{data: res}
}

func FlatMapConcat[S any, T any](stream *Stream[S], mapper func(S) *Stream[T]) *Stream[T] {
	var res []T
	for _, v := range stream.data {
		res = append(res, (*mapper(v)).data...)
	}
	return &Stream[T]{data: res}
}
