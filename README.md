# go-zaim

[![Go Reference](https://pkg.go.dev/badge/github.com/ryohidaka/go-zaim.svg)](https://pkg.go.dev/github.com/ryohidaka/go-zaim)
![GitHub Release](https://img.shields.io/github/v/release/ryohidaka/go-zaim)
[![codecov](https://codecov.io/gh/ryohidaka/go-zaim/graph/badge.svg?token=4S2qaMq8BY)](https://codecov.io/gh/ryohidaka/go-zaim)
[![Go Report Card](https://goreportcard.com/badge/github.com/ryohidaka/go-zaim)](https://goreportcard.com/report/github.com/ryohidaka/go-zaim)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go 用 Zaim API クライアント

## インストール

```bash
go get github.com/ryohidaka/go-zaim
```

## ドキュメント

Read [GoDoc](https://pkg.go.dev/github.com/ryohidaka/go-zaim)

## 使用例

> [!IMPORTANT]
> 一部の API エンドポイントでは `OAuth 1.0a` での認証が必須です。

```go
import "github.com/ryohidaka/go-zaim"


func main() {
    // 認証情報を取得
    p := zaim.ZaimParams{
        ConsumerKey:    "<CONSUMER_KEY>",
        ConsumerSecret: "<CONSUMER_SECRET>",
        Token:          "<TOKEN>",
        TokenSecret:    "<TOKEN_SECRET>",
    }

    // クライアントを初期化
    c := zaim.NewZaimClient(p)

    // 認証ユーザーの情報を取得
    me, err := c.FetchMe()

    // 入出金履歴を取得
    money, err := c.FetchMoney()
}
```

## リンク

- [API リファレンス (Zaim developers)](https://dev.zaim.net/)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
