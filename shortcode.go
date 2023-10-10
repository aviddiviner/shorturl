package shortcode

import (
	"errors"
	"math"
	"strings"
)

const (
	DefaultAlphabet  = "mn6j2c4rv8bpygw95z7hsdaetxuk3fq"
	DefaultBlockSize = 24
	MinLength        = 5
)

type Encoder struct {
	Alphabet  string
	BlockSize int
	Mask      int
}

func NewEncoder(alphabet string, blockSize int) (*Encoder, error) {
	n := len(alphabet)
	if n < 2 {
		return nil, errors.New("alphabet must contain at least 2 characters")
	}
	alphabet = dedupe(alphabet)
	if n != len(alphabet) {
		return nil, errors.New("alphabet must contain unique characters")
	}

	return &Encoder{
		Alphabet:  alphabet,
		BlockSize: blockSize,
		Mask:      (1 << blockSize) - 1,
	}, nil
}

func dedupe(s string) string {
	set := make(map[rune]bool)
	var builder strings.Builder
	for _, char := range s {
		if !set[char] {
			set[char] = true
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

func (e *Encoder) EncodeID(n, minLength int) string {
	return e.Enbase(e.Encode(n), minLength)
}

func (e *Encoder) DecodeID(s string) int {
	return e.Decode(e.Debase(s))
}

func (e *Encoder) Encode(n int) int {
	return (n & ^e.Mask) | e.encode(n&e.Mask)
}

func (e *Encoder) encode(n int) int {
	result := 0
	for i := 0; i < e.BlockSize; i++ {
		b := e.BlockSize - i - 1
		if n&(1<<i) != 0 {
			result |= 1 << b
		}
	}
	return result
}

func (e *Encoder) Decode(n int) int {
	return (n & ^e.Mask) | e.decode(n&e.Mask)
}

func (e *Encoder) decode(n int) int {
	result := 0
	for i := 0; i < e.BlockSize; i++ {
		b := e.BlockSize - i - 1
		if n&(1<<b) != 0 {
			result |= 1 << i
		}
	}
	return result
}

func (e *Encoder) Enbase(x, minLength int) string {
	result := e.enbase(x)
	paddingCount := minLength - len(result)
	if paddingCount < 0 {
		paddingCount = 0
	}
	padding := strings.Repeat(string(e.Alphabet[0]), paddingCount)
	return padding + result
}

func (e *Encoder) enbase(x int) string {
	n := len(e.Alphabet)
	if x < n {
		return string(e.Alphabet[x])
	}
	return e.enbase(x/n) + string(e.Alphabet[x%n])
}

func (e *Encoder) Debase(s string) int {
	n := len(e.Alphabet)
	result := 0
	for i := 0; i < len(s); i++ {
		c := s[len(s)-i-1]
		index := strings.IndexByte(e.Alphabet, c)
		if index != -1 {
			result += index * int(math.Pow(float64(n), float64(i)))
		}
	}
	return result
}

var defaultEncoder, _ = NewEncoder(DefaultAlphabet, DefaultBlockSize)

func Encode(n int) int {
	return defaultEncoder.Encode(n)
}

func Decode(n int) int {
	return defaultEncoder.Decode(n)
}

func Enbase(n int) string {
	return defaultEncoder.Enbase(n, MinLength)
}

func Debase(s string) int {
	return defaultEncoder.Debase(s)
}

func EncodeID(n int) string {
	return defaultEncoder.EncodeID(n, MinLength)
}

func DecodeID(s string) int {
	return defaultEncoder.DecodeID(s)
}
