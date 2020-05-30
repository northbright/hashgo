package main

import (
	"crypto"
	"strconv"
)

// getHashString returns the string of given hash.
// It should use Hash.String() of crypto package instead when Go 1.15 releases.
func getHashString(h crypto.Hash) string {
	switch h {
	case crypto.MD4:
		return "MD4"
	case crypto.MD5:
		return "MD5"
	case crypto.SHA1:
		return "SHA-1"
	case crypto.SHA224:
		return "SHA-224"
	case crypto.SHA256:
		return "SHA-256"
	case crypto.SHA384:
		return "SHA-384"
	case crypto.SHA512:
		return "SHA-512"
	case crypto.MD5SHA1:
		return "MD5+SHA1"
	case crypto.RIPEMD160:
		return "RIPEMD-160"
	case crypto.SHA3_224:
		return "SHA3-224"
	case crypto.SHA3_256:
		return "SHA3-256"
	case crypto.SHA3_384:
		return "SHA3-384"
	case crypto.SHA3_512:
		return "SHA3-512"
	case crypto.SHA512_224:
		return "SHA-512/224"
	case crypto.SHA512_256:
		return "SHA-512/256"
	case crypto.BLAKE2s_256:
		return "BLAKE2s-256"
	case crypto.BLAKE2b_256:
		return "BLAKE2b-256"
	case crypto.BLAKE2b_384:
		return "BLAKE2b-384"
	case crypto.BLAKE2b_512:
		return "BLAKE2b-512"
	default:
		return "unknown hash value " + strconv.Itoa(int(h))
	}
}
