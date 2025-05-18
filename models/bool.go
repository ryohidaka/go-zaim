package models

import (
	"encoding/json"
)

// BoolInt は JSON の整数 (0 または 1) を bool に変換するための型
// JSON デコード時に 0 は false、0 以外は true として扱う
type BoolInt bool

// UnmarshalJSON は JSON の整数値を BoolInt に変換する
// 0 は false、0 以外は true に設定される。
// JSON の整数以外が渡された場合はエラーを返す。
func (b *BoolInt) UnmarshalJSON(data []byte) error {
	var tmp int
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = tmp != 0
	return nil
}
