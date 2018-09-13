package main

import (
    "fmt"
    "flag"
    "github.com/wxnacy/wcrypto"
    "os"
    "strings"
)

const VERSION = "0.2.2"

var HELP  = fmt.Sprintf(`
Version: %s

Uage:
    wcrypto <md5|sha1|sha256|sha512> <string|file>

    Or

    wcrypto -k <key> <hmacmd5|hmacsha1|hmacsha256|hmacsha512> <string|file>


Or use pipeline, Like

cat file | wcrypto <md5|sha1|sha256|sha512>
`, VERSION)

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

func HmacFromFile(f *os.File) {
    h, err := wcrypto.HmacHashFile(Mode, Key, f)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(h)
}

func HmacFromMsg(msg string) {
    var h string
    var err error

    if _, flag := wcrypto.IsFile(msg); flag {
        h, err = wcrypto.HmacHashPath(Mode, Key, msg)
    } else {
        h, err = wcrypto.HmacHash(Mode, Key, msg)
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
var Key string              // 秘钥
var ModeType string         // 加密类型
var Ver bool                // 查看简单版本信息
var Version bool            // 查看版本信息

func InitArgs() bool {
    flag.StringVar(&Key, "k", "", "Security key")
    flag.BoolVar(&Ver, "v", false, "Security key")
    flag.BoolVar(&Version, "V", false, "Security key")

    flag.Parse()
    // fmt.Println(Key)

    Args = flag.Args()
    // fmt.Println(Args)
    ModeType = wcrypto.ModeTypeHash
    Stdin, IsStdin = HasStdin()

    return CheckArgs()
}

func CheckArgs() bool{

    if Ver {
        fmt.Println(VERSION)
        return false
    }

    if Version {
        fmt.Println(HELP)
        return false
    }

    argsNeedCount := 2
    if IsStdin {
        argsNeedCount = 1
    }

    if len(Args) < argsNeedCount {
        fmt.Println(HELP)
        return false
    }

    Mode = Args[0]

    if strings.HasPrefix(Mode, "hmac") && Key == "" {
        fmt.Printf("Uage: wcrypto -k <key> %s <string|file>\n", Mode)
        return false
    }

    if strings.HasPrefix(Mode, "hmac") {
        ModeType = wcrypto.ModeTypeHmac
    }
    return true
}

func main() {

    if !InitArgs() {
        return
    }

    msgs := make([]string, 0)

    msgs = Args[1:len(Args)]

    switch ModeType {
        case wcrypto.ModeTypeHash: {
            if IsStdin {
                HashFromFile(Stdin)
                return
            }
            for _, msg := range msgs {
                HashFromMsg(msg)
            }
        }
        case wcrypto.ModeTypeHmac: {
            if IsStdin {
                HmacFromFile(Stdin)
                return
            }
            for _, msg := range msgs {
                HmacFromMsg(msg)
            }
        }
    }

    _ = msgs
}

