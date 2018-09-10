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
