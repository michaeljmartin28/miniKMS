package core

import "fmt"

func ParseAlgorithm(s string) (Algorithm, error) {
	switch Algorithm(s) {
	case AES256GCM:
		return AES256GCM, nil
	default:
		return "", fmt.Errorf("unknown algorithm: %s", s)
	}
}
