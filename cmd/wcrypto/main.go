package main

import (
    "fmt"
    "flag"
    "github.com/wxnacy/wcrypto"
    "os"
)

func HasStdin() (*os.File, bool) {
    fileInfo, _ := os.Stdin.Stat()
    if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
        return nil, false
    }
    return os.Stdin, true
}

const HELP  = `
Version: 0.1.0

Uage: wcrypto <param> <string|file>

Or use pipeline, Like

cat file | wcrypto <param>

Params has md5|sha1|sha256|sha512
`

func main() {
    flag.Parse()
    args := flag.Args()

    msgs := make([]string, 0, 0)

    argsNeedCount := 2
    stdin, flag := HasStdin()
    if flag {
        argsNeedCount = 1
    }

    if len(args) < argsNeedCount {
        fmt.Println(HELP)
        return
    }

    cmd := args[0]

    msgs = args[1:len(args)]
    if flag {
        h, err := wcrypto.HashFile(cmd, stdin)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(h)
        return
    }

    for _, msg := range msgs {
        h, err := wcrypto.Hash(cmd, msg)
        if _, flag := wcrypto.IsFile(msg); flag {
            h, err = wcrypto.HashPath(cmd, msg)
        }
        if err != nil {
            fmt.Println(err)
            return
        }
        if flag {
            fmt.Println(h)
        } else {
            fmt.Printf("%s %s\n", h, msg)
        }
    }


    _ = msgs
    _ = cmd
}

