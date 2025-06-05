package auth

import "github.com/alexedwards/argon2id"

func HashPassword(password string) (string, error) {
	return argon2id.CreateHash(password, &argon2id.Params{
		Memory:      256 * 1024,
		Iterations:  12,
		Parallelism: 8,
		SaltLength:  18,
		KeyLength:   32,
	})
}

func ComparePassword(password, hashed string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, hashed)
	if err != nil {
		return false
	}
	return match
}
