package hasher

type PwHasher interface {
	Hash(plain string) (string, error)
	Verify(plain, hashed string) (bool, error)
}
