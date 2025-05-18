package models

// Requested は全てのレスポンスに含まれる requested フィールド
type Requested struct {
	Requested uint64 `json:"requested"`
}
