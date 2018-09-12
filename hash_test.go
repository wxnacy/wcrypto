package wcrypto

import (
    "testing"
)

func TestHashPath(t *testing.T) {
    filepath := "LICENSE"

    modes := []string{MD5, SHA1, SHA256, SHA512}

    for _, d := range modes {

        h, err := HashPath(d, filepath)
        if err != nil {
            t.Error(err)
        }

        switch d {
            case  MD5: {
                if h != "06214ba9fecad17362f12ab37e720767" {
                    t.Errorf("%s %s 结果不正确", d, h)
                }
            }
            case  SHA1: {
                if h != "f9acfd6120ae751a6bcbfce5c74317cc85691e19" {
                    t.Errorf("%s %s 结果不正确", d, h)
                }
            }
            case  SHA256: {
                if h != "53ade92cf56ebdd6c5c2751c620c3a4302363c82c0c1eb66999f3df316262b0b" {
                    t.Errorf("%s %s 结果不正确", d, h)
                }
            }
            case  SHA512: {
                if h != "e6907fce8f15c4157101932485af674f6e9b416f9b7d28e4769748313fff53d8d416847013c62c4a40c51f894184fb02d1b971515f5a1a12e9b82b5cec1c52bf" {
                    t.Errorf("%s %s 结果不正确", d, h)
                }
            }
        }
    }
}

func TestHash(t *testing.T) {

    modes := []string{MD5, SHA1, SHA256, SHA512}

    for _, d := range modes {

        h, err := Hash(d, "wxnacy")
        if err != nil {
            t.Error(err)
        }

        flag := false

        switch d {
            case  MD5: {
                if h != "1f806eb48b670c40af49a3f764ba086f" {
                    flag = true
                }
            }
            case  SHA1: {
                if h != "ae80552bbe355867a1579ab25dfb3a49ac5ffae5" {
                    flag = true
                }
            }
            case  SHA256: {
                if h != "e272638378933bcd0921396695cc357a5f8ed7c136d06878d0b9c9ae0302c14a" {
                    flag = true
                }
            }
            case  SHA512: {
                if h != "edc44730889e61c7674c6f80c550a865d222ac9214cbb310e61303c5b1fc6bc3ea801a95a3dc2070d2c90aa7a5cae53bbc417b0c10be2e0d33d41d6a68cbf822" {
                    flag = true
                }
            }
        }
        if flag {
            t.Errorf("%s %s 结果不正确", d, h)
        }
    }
}

func TestHmacHash(t *testing.T) {

    modes := []string{HMACMD5, HMACSHA1, HMACSHA256, HMACSHA512}

    for _, d := range modes {

        h, err := HmacHash(d, "wxnacy", "wxnacy")
        if err != nil {
            t.Error(err)
        }

        flag := false



        switch d {
            case  HMACMD5: {
                if h != "84b12d9307fa423d41ae9719efc60ef3" {
                    flag = true
                }
            }
            case  HMACSHA1: {
                if h != "852c1f85f3fd1b4f9eb0f60647d3026773500896" {
                    flag = true
                }
            }
            case  HMACSHA256: {
                if h != "29e86b1e142731d6b69029a074ba1d8b8d37cd62f6f2d25a74409a3a27e90e1c" {
                    flag = true
                }
            }
            case  HMACSHA512: {
                if h != "b8dfd4f067a0b3697720032dd735dd1110f93ec87a351918629a533f9c81b8f30378934b287d4aca668df4a0edb270a962b745eaf85029f860bca8e2ca4abfca" {
                    flag = true
                }
            }
        }
        if flag {
            t.Errorf("%s %s 结果不正确", d, h)
        }
    }
}
