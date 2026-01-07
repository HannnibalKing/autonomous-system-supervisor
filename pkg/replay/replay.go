package replay

import (
	"encoding/binary"
	"math/rand"
)

// StateDump captures binary state for deterministic replay.
type StateDump struct {
	Seed   int64
	States map[string][]byte
}

// RNG returns a deterministic random generator configured from the dump.
func (s StateDump) RNG() *rand.Rand {
	return rand.New(rand.NewSource(s.Seed))
}

// MarshalBinary encodes the dump for persistence.
func (s StateDump) MarshalBinary() ([]byte, error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, uint64(s.Seed))
	return buf, nil
}

// UnmarshalBinary decodes a dump.
func (s *StateDump) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return nil
	}
	s.Seed = int64(binary.LittleEndian.Uint64(data[:8]))
	return nil
}
