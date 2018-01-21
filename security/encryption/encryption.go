package encryption

type Encryptor interface{
	Encrypt(message string) string
}