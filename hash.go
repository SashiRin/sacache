package sacache

type Hasher interface {
	Hash(string) uint64
}
