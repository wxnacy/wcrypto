package wcrypto

import (
    "fmt"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "crypto/hmac"
    "errors"
    "os"
    "io"
    "hash"
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

func HmacHash(mode, key, msg string) (string, error) {

    var h hash.Hash
    switch mode {
        case HMACMD5: {
            h = hmac.New(md5.New, []byte(key))
        }
        case HMACSHA1: {
            h = hmac.New(sha1.New, []byte(key))
        }
        case HMACSHA256: {
            h = hmac.New(sha256.New, []byte(key))
        }
        case HMACSHA512: {
            h = hmac.New(sha512.New, []byte(key))
        }
        default: {
            return "", errors.New(mode + ": This mode is not supported")
        }
    }
    h.Write([]byte(msg))
    return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func HmacHashPath(mode, key, path string) (string, error) {

    f, err := os.Open(path)
	if err != nil {
        return "", err
	}
	defer f.Close()

    return HmacHashFile(mode, key, f)
}

func HmacHashFile(mode, key string, f *os.File) (string, error) {

    var h hash.Hash
    var err error

    switch mode {
        case HMACMD5: {
            h = hmac.New(md5.New, []byte(key))
        }
        case HMACSHA1: {
            h = hmac.New(sha1.New, []byte(key))
        }
        case HMACSHA256: {
            h = hmac.New(sha256.New, []byte(key))
        }
        case HMACSHA512: {
            h = hmac.New(sha512.New, []byte(key))
        }
        default: {
            return "", errors.New(mode + ": This mode is not supported")
        }
    }
	if _, err = io.Copy(h, f); err != nil {
        return "", err
	}
    return fmt.Sprintf("%x", h.Sum(nil)), nil
}
