package encryption

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type sha256Encryptor struct{
	secretKey string
}

type Sha256EncryptorFactory struct{}

func (f *Sha256EncryptorFactory) NewSha256Encryptor(secretKey string)*sha256Encryptor{
	return &sha256Encryptor{secretKey:secretKey}
}

func (s *sha256Encryptor) Encrypt(message string) string{
	key := []byte(s.secretKey)
	sig := hmac.New(sha256.New, key)
	sig.Write([]byte(message))
	return hex.EncodeToString(sig.Sum(nil))
}