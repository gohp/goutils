package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"os"
)

// Md5Byte ...
func Md5Byte(s []byte) string {
	h := md5.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// Md5String ...
func Md5String(s string) string {
	return Md5Byte([]byte(s))
}

// Md5File ...
func Md5File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Sha1Byte ...
func Sha1Byte(s []byte) string {
	h := sha1.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1String ...
func Sha1String(s string) string {
	return Sha1Byte([]byte(s))
}

// Sha1File ...
func Sha1File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Sha256Byte ...
func Sha256Byte(s []byte) string {
	h := sha256.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// Sha256String ...
func Sha256String(s string) string {
	return Sha256Byte([]byte(s))
}

// Sha256File ...
func Sha256File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Sha512Byte ...
func Sha512Byte(s []byte) string {
	h := sha512.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// Sha512String ...
func Sha512String(s string) string {
	return Sha512Byte([]byte(s))
}

// Sha512File ...
func Sha512File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha512.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
