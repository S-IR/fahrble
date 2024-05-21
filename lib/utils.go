package lib

import "math/rand"

func RandomString(length uint) string {
	if length < 1 {
		return ""
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		num := rand.Intn(int(length))
		result[i] = charset[num]
	}
	return string(result)
}
