package helpers

import (
	"crypto/rand"
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func DeriveKey(password string, salt []byte) ([]byte, []byte){
	if salt == nil {
		salt = make([]byte, 8)
		rand.Read(salt)
	}

	return pbkdf2.Key([]byte(password), salt, 1024, 32, sha256.New), salt
}