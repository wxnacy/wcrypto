package wcrypto

import (
    "fmt"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "errors"
    "os"
    "io"
    "hash"
)

const (
    MD5 = "md5"
    SHA1 = "sha1"
    SHA256 = "sha256"
    SHA512 = "sha512"
)

func HashFile(mode string, f *os.File) (string, error) {
    var result string
    var h hash.Hash
    var err error
    switch mode {
        case MD5: {
            h = md5.New()
        }
        case SHA1: {
            h = sha1.New()
        }
        case SHA256: {
            h = sha256.New()
        }
        case SHA512: {
            h = sha512.New()
        }
        default: {
            return "", errors.New(mode + ": This mode is not supported")
        }
    }
	if _, err = io.Copy(h, f); err != nil {
        return "", err
	}
    result = fmt.Sprintf("%x", h.Sum(nil))
    return result, nil
}

func HashPath(mode, filepath string) (string, error) {
    f, err := os.Open(filepath)
	if err != nil {
        return "", err
	}
	defer f.Close()
    return HashFile(mode, f)
}

func Hash(mode, msg string) (string, error) {
    var h hash.Hash
    switch mode {
        case MD5: {
            h = md5.New()
        }
        case SHA1: {
            h = sha1.New()
        }
        case SHA256: {
            h = sha256.New()
        }
        case SHA512: {
            h = sha512.New()
        }
        default: {
            return "", errors.New(mode + ": This mode is not supported")
        }
    }
    io.WriteString(h, msg)
    return fmt.Sprintf("%x", h.Sum(nil)), nil
}
