package models

import (
	"fmt"
	"strings"
	"time"
)

// ZaimTime は Zaim API の日付文字列（"2006-01-02 15:04:05"）を
// time.Time に変換するためのカスタム型。
type ZaimTime struct {
	time.Time
}

// Zaim API で使われる日時フォーマットを表す。
const (
	zaimLayoutDateTime = "2006-01-02 15:04:05"
	zaimLayoutDate     = "2006-01-02"
)

// UnmarshalJSON は JSON の日時文字列を ZaimTime 型にパースする。
// 文字列の両端のダブルクォートを除去してから解析を行い、
// フォーマットに合わない場合はエラーを返す。
func (zt *ZaimTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "null" || s == "" {
		return fmt.Errorf("invalid time format: empty or null")
	}

	// 日時付き -> パース
	if t, err := time.Parse(zaimLayoutDateTime, s); err == nil {
		zt.Time = t
		return nil
	}

	// 日付のみ -> パース
	if t, err := time.Parse(zaimLayoutDate, s); err == nil {
		zt.Time = t
		return nil
	}

	return fmt.Errorf("ZaimTime parse error: unknown format: %s", s)
}
