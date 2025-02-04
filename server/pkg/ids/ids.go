package ids

import (
	"fmt"
	"strconv"
	"strings"
)

const zeroID = 0

type identifier interface {
	Int64() int64
	String() string
}

func encode(i uint64) string {
	return fmt.Sprintf("%016x", i)
}

func decode(s string) (uint64, error) {
	return strconv.ParseUint(s, 16, 64)
}

func isZero(id int64) bool {
	return id == zeroID
}

func parseID(s string, prefix string) (uint64, error) {
	if s == "" {
		return zeroID, nil
	}
	if !strings.HasPrefix(s, chunkPrefix) {
		return 0, fmt.Errorf("invalid prefix in discovery ID '%s'", s)
	}
	c, err := decode(s[len(chunkPrefix):])
	if err != nil {
		return 0, fmt.Errorf("unable to parse discovery ID '%s': %w", s, err)
	}
	return c, nil
}
