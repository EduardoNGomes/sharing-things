package cryptography

type Encrypter interface {
	Encrypt(password string) (string, error)
	Comparer(password, hash string) bool
}
