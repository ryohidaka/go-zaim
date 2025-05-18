package models

import (
	"fmt"
	"time"
)

// ZaimTime は Zaim API の日付文字列（"2006-01-02 15:04:05"）を
// time.Time に変換するためのカスタム型。
type ZaimTime struct {
	time.Time
}

// zaimLayout は Zaim API で使われる日時フォーマットを表す。
const zaimLayout = "2006-01-02 15:04:05"

// UnmarshalJSON は JSON の日時文字列を ZaimTime 型にパースする。
// 文字列の両端のダブルクォートを除去してから解析を行い、
// フォーマットに合わない場合はエラーを返す。
func (zt *ZaimTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	if len(s) < 2 {
		return fmt.Errorf("invalid time format: %s", s)
	}
	// JSONの文字列型のダブルクォート " を取り除く
	s = s[1 : len(s)-1]

	t, err := time.Parse(zaimLayout, s)
	if err != nil {
		return fmt.Errorf("time parse error: %w", err)
	}

	zt.Time = t
	return nil
}
