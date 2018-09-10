# wcrypto

wxnacy crypto 是一款类似 md5sum 进行 md5、sha1 等多种模式加密的学习作品。

## 安装

### by go

```bash
$ go get github.com/wxnacy/wcrypto
$ go install github.com/wxnacy/wcrypto/cmd/wcrypto
```

### by wget

```bash
$ wget https://raw.githubusercontent.com/wxnacy/wcrypto/master/bin/wcrypto
$ sudo cp wcrypto /usr/bin
```

## 使用

**语法**

```bash
$ wcrypto <mode> <string|file>...
```

**例子**

```bash
$ wcrypto md5 wxnacy
1f806eb48b670c40af49a3f764ba086f wxnacy
```

```bash
$ wcrypto md5 wxnacy README.md
1f806eb48b670c40af49a3f764ba086f wxnacy
114fcca532346c70863ab8deb6c1271c README.md
```

```bash
$ curl -s http://baidu.com | wcrypto md5
a6c4b5d58389762e8e7f67c8a3515d3f
```

**mode**

mode 取值为：`md5, sha1, sha256, sha512`
