package models_test

import (
	"encoding/json"
	"testing"

	"github.com/ryohidaka/go-zaim/models"
)

func TestBoolInt_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		jsonInput string         // テスト用のJSON入力文字列
		expected  models.BoolInt // 期待するBoolIntの値
		wantErr   bool           // エラー発生の有無
	}{
		{`0`, false, false},       // 0 は false になる
		{`1`, true, false},        // 1 は true になる
		{`2`, true, false},        // 0以外の整数は true になる
		{`-1`, true, false},       // マイナスも true になる
		{`"string"`, false, true}, // 整数以外はエラーになる
	}

	for _, tt := range tests {
		var b models.BoolInt
		err := json.Unmarshal([]byte(tt.jsonInput), &b)
		if (err != nil) != tt.wantErr {
			t.Errorf("UnmarshalJSON(%s) でエラー = %v, 期待するエラーの有無 = %v", tt.jsonInput, err, tt.wantErr)
			continue
		}
		if err == nil && b != tt.expected {
			t.Errorf("UnmarshalJSON(%s) = %v, 期待値 = %v", tt.jsonInput, b, tt.expected)
		}
	}
}
