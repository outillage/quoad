package quoad

import "encoding/hex"

// Commit is a parsed commit that contains information about category, scope and heading
type Commit struct {
	Category string
	Scope    string
	Breaking bool
	Heading  string
	Body     string
	Hash     Hash
	Issues   []int
}

// Hash describes a commit hash
type Hash [20]byte

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}
