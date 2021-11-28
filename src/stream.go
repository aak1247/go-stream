/*
Go generic stream

Due to the defect of go generic type, i am not able to implement a fullly-functional generic stream.


*/

package stream

// type IStream[T any] interface {
// 	map(func(T any) any) IStream(T)
// }


// type Steam
type Stream[T any] struct{
	data []T
}


func (s *Stream[T]) Foreach(f func(T)) *Stream[T] {
	for _, v := range s.data {
		f(v)
	}
	return s
}

// a fully functional map method is not supported yet
func (s *Stream[T]) Map(f func(T) interface{}) *Stream[interface{}] {
	var res []interface{}
	for _, v := range s.data {
		res = append(res, f(v))
	}
	return &Stream[interface{}]{data: res}
}

func (s *Stream[T]) Filter(f func(T) bool) *Stream[T] {
	var res []T
	for _, v := range s.data {
		if f(v) {
			res = append(res, v)
		}
	}
	return &Stream[T]{data: res}
}

func (s *Stream[T]) Reduce(identity T, binaryOperator func(T, T) T) T {
	var res T = identity
	for _, v := range s.data {
		res = binaryOperator(res, v)
	}
	return res
}

func (s *Stream[T]) ReduceRight(identity T, f func(T, T) T) T {
	var res T = identity
	for i := len(s.data) - 1; i >= 0; i-- {
		res = f(res, s.data[i])
	}
	return res
}

func (s *Stream[T]) Concat(other *Stream[T]) *Stream[T] {
	return &Stream[T]{data: append(s.data, other.data...)}
}

func (s *Stream[T]) Take(n int) (*Stream[T], bool) {
	if n > len(s.data) {
		return &Stream[T]{data: s.data}, false
	}
	return &Stream[T]{data: s.data[:n]}, true
}

func (s *Stream[T]) Skip(n int) (*Stream[T], bool) {
	if n > len(s.data) {
		return &Stream[T]{data: s.data}, false
	}
	return &Stream[T]{data: s.data[n:]}, true
}
func (s *Stream[T]) TakeWhile(f func(T) bool) *Stream[T] {
	var res []T
	for _, v := range s.data {
		if f(v) {
			res = append(res, v)
		} else {
			break
		}
	}
	return &Stream[T]{data: res}
}

func (s *Stream[T]) Drop(n int) (*Stream[T], bool) {
	if n > len(s.data) {
		return &Stream[T]{data: s.data}, false
	}
	return &Stream[T]{data: s.data[n:]}, true
}

func (s *Stream[T]) DropWhile(f func(T) bool) *Stream[T] {
	var res []T
	for _, v := range s.data {
		if !f(v) {
			res = append(res, v)
		}
	}
	return &Stream[T]{data: res}
}

func (s *Stream[T]) TakeLast(n int) (*Stream[T], bool) {
	if n > len(s.data) {
		return &Stream[T]{data: s.data}, false
	}
	return &Stream[T]{data: s.data[len(s.data)-n:]}, true
}

// not recommended, use Stream.FlatMap Instead
func (s *Stream[T]) FlatMap(mapper func(T) *Stream[interface{}]) *Stream[interface{}] {
	var res []interface{}
	for _, v := range s.data {
		res = append(res, (*mapper(v)).data...)
	}
	return &Stream[interface{}]{data: res}
}

// not recommended, use Stream.FlatMapConcat Instead
func (s *Stream[T]) FlatMapConcat(mapper func(T) *Stream[interface{}]) *Stream[T] {
	var res []T
	for _, v := range s.data {
		res = append(res, (*mapper(v)).data...)
	}
	return &Stream[interface{}]{data: res}
}

func (s *Stream[T]) ToArray() []T {
	return s.data
}