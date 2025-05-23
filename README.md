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

### 初期化

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
}
```

### ユーザー情報

```go
// 認証ユーザーの情報を取得
me, err := c.FetchMe()
```

### 入出金履歴

```go
// 入出金履歴を取得
money, err := c.FetchMoney()

// 入出金履歴を取得 (グルーピング形式)
money, err := c.FetchGroupedMoney()
```

### カテゴリ

```go
// カテゴリ一覧を取得
categories, err := c.FetchCategories()

// デフォルトカテゴリ一覧を取得
categories, err := c.FetchDefaultCategories()
```

### ジャンル

```go
// ジャンル一覧を取得する
genres, err := c.FetchGenres()
```

### 口座

```go
// 口座一覧を取得する
accounts, err := c.FetchAccounts()

// デフォルト口座一覧を取得する
accounts, err := c.FetchDefaultAccounts()
```

### 支払情報

```go
// 支払情報を登録する
res, err := c.CreatePayment(zaim.CreatePaymentParams{
    CategoryID: 102,
    GenreID:    10202,
    Amount:     1,
})

// 支払情報を更新する
res, err := c.UpdatePayment(381, zaim.UpdatePaymentParams{
    CategoryID: 102,
    GenreID:    10202,
    Amount:     1,
})

// 支払情報を削除する
res, err := c.DeletePayment(381)
```

### 収入情報

```go
// 収入情報を登録する
res, err := c.CreateIncome(zaim.CreateIncomeParams{
    CategoryID: 102,
    Amount:     1,
})

// 収入情報を更新する
res, err := c.UpdateIncome(381, zaim.UpdateIncomeParams{
    CategoryID: 102,
    Amount:     1,
})

// 収入情報を削除する
res, err := c.DeleteIncome(381)
```

### 振替情報

```go
// 振替情報を登録する
res, err := c.CreateTransfer(zaim.CreateTransferParams{
    Amount: 1,
    Date:   time.Now(),
    FromAccountID: 1,
    ToAccountID:   2,
})

// 振替情報を更新する
res, err := c.UpdateTransfer(381, zaim.CreateTransferParams{
    Amount: 1,
    Date:   time.Now(),
    FromAccountID: 1,
    ToAccountID:   2,
})

// 振替情報を削除する
res, err := c.DeleteTransfer(381)
```

## リンク

- [API リファレンス (Zaim developers)](https://dev.zaim.net/)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
