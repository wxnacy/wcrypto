package main

import (
    "fmt"
    "flag"
    "github.com/wxnacy/wcrypto"
    "os"
)

const HELP  = `
Version: 0.1.1

Uage: wcrypto <param> <string|file>

Or use pipeline, Like

cat file | wcrypto <param>

Params has md5|sha1|sha256|sha512
`

func HasStdin() (*os.File, bool) {
    fileInfo, _ := os.Stdin.Stat()
    if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
        return nil, false
    }
    return os.Stdin, true
}


func HashFromFile(f *os.File) {

    h, err := wcrypto.HashFile(Mode, f)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(h)
}

func HashFromMsg(msg string) {
    var h string
    var err error

    if _, flag := wcrypto.IsFile(msg); flag {
        h, err = wcrypto.HashPath(Mode, msg)
    } else {
        h, err = wcrypto.Hash(Mode, msg)
    }
    if err != nil {
        fmt.Println(err)
        return
    }
    PrintResult(h, msg)
}

func PrintResult(h, msg string) {
    if IsStdin {
        fmt.Println(h)
    } else {
        fmt.Printf("%s %s\n", h, msg)
    }
}

var Stdin *os.File          // 管道数据
var IsStdin bool            // 是否有管道数据
var Args = make([]string, 0)
var Mode string             // 加密模式

func InitArgs() {

    flag.Parse()

    Args = flag.Args()
    Stdin, IsStdin = HasStdin()
}

func CheckArgs() bool{
    argsNeedCount := 2
    if IsStdin {
        argsNeedCount = 1
    }

    if len(Args) < argsNeedCount {
        fmt.Println(HELP)
        return false
    }
    return true
}

func main() {

    InitArgs()
    if !CheckArgs() {
        return
    }

    msgs := make([]string, 0)
    Mode = Args[0]

    msgs = Args[1:len(Args)]
    if IsStdin {
        HashFromFile(Stdin)
        return
    }

    for _, msg := range msgs {
        HashFromMsg(msg)
    }

    _ = msgs
}

