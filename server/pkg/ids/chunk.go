package ids

const chunkPrefix = "ch-"

type Chunk int64

var _ identifier = Chunk(zeroID)

func (c Chunk) Int64() int64 {
	return int64(c)
}

func (c Chunk) String() string {
	return chunkPrefix + encode(uint64(c))
}

func ParseChunk(s string) (Chunk, error) {
	c, err := parseID(s, chunkPrefix)
	return Chunk(c), err
}
