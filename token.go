package token

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost int = 10 // the minimum allowable cost as passed in to Hash
	MaxCost int = 31 // the maximum allowable cost as passed in to Hash
)

// Token creates a token hash
// key (string) - a key
// cost (int) - which denotes the algorithmic cost that should be used
func Hash(key string, cost int) string {
	cost = fixCost(cost)
	h, _ := bcrypt.GenerateFromPassword([]byte(key), cost)
	return base64.StdEncoding.EncodeToString(h)
}

// Verify verifies that a token hash matches a key and cost
// key (string) - a key
// cost (int) - which denotes the algorithmic cost that should be used
// hash (string) - a token hash created by Hash
func Verify(key string, cost int, hash string) bool {
	h, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false
	}
	cost = fixCost(cost)
	if c, _ := bcrypt.Cost([]byte(h)); c != cost {
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(h), []byte(key)) == nil
}

func fixCost(c int) int {
	if c < MinCost {
		return MinCost
	}
	if c > MaxCost {
		return MaxCost
	}
	return c
}
