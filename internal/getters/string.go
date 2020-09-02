package getters

import (
	"fmt"
	"math/rand"

	"github.com/icrowley/fake"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n uint64) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandomString getter
type RandomString struct {
	name      string
	maxSize   int64
	allowNull bool
}

func (r *RandomString) Value() interface{} {
	if r.allowNull && rand.Int63n(100) < nilFrequency {
		return nil
	}
	var s string
	maxSize := uint64(r.maxSize)
	if maxSize == 0 {
		maxSize = uint64(rand.Int63n(100))
	}

	if maxSize < 30 {
		s = randSeq(maxSize)
	} else {
		s = fake.Sentence()
	}
	if len(s) > int(maxSize) {
		s = s[:int(maxSize)]
	}
	return s
}

func (r *RandomString) String() string {
	v := r.Value()
	if v == nil {
		return NULL
	}
	return v.(string)
}

func (r *RandomString) Quote() string {
	v := r.Value()
	if v == nil {
		return NULL
	}
	return fmt.Sprintf("%q", v)
}

func NewRandomString(name string, maxSize int64, allowNull bool) *RandomString {
	return &RandomString{name, maxSize, allowNull}
}
