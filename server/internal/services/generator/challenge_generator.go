package generator

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GenerateChallenge(difficulty int) string {
	randomInt := rand.Int()

	challenge := fmt.Sprintf("%d:%d", randomInt, difficulty)

	return challenge
}
