package wcrypto

import (
    "testing"
)

func TestIsExists(t *testing.T) {
    filename := "file_test.go"
    _, flag := IsExists(filename)
    if ! flag {
        t.Errorf("%s is exists", filename)
    }

    filename = "file_tests.go"
    _, flag = IsExists(filename)
    if flag {
        t.Errorf("%s is not exists", filename)
    }

}

func TestIsFile(t *testing.T) {

    filename := "file_test.go"
    _, flag := IsFile(filename)
    if ! flag {
        t.Errorf("%s is file", filename)
    }

    filename = "file_test1.go"
    _, flag = IsFile(filename)
    if flag {
        t.Errorf("%s is not file", filename)
    }

    filename = "cmd"
    _, flag = IsFile(filename)
    if flag {
        t.Errorf("%s is not file", filename)
    }

}

func TestIsDir(t *testing.T) {
    filename := "cmd"
    _, flag := IsDir(filename)
    if ! flag {
        t.Errorf("%s is dir", filename)
    }

    filename = "file_test.go"
    _, flag = IsDir(filename)
    if flag {
        t.Errorf("%s is not dir", filename)
    }

    filename = "cmds"
    _, flag = IsDir(filename)
    if flag {
        t.Errorf("%s is not dir", filename)
    }
}
