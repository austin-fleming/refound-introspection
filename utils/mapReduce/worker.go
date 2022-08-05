package mapReduce

type KvPair[T string | uint16 | uint64] struct {
	Key   string
	Value T
}

type MapFn[T string | uint16 | uint64] func(string, T) KvPair[T]

func Worker[T string | uint16 | uint64](mapper MapFn[T]) {

}
