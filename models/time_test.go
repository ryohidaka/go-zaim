package models_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/ryohidaka/go-zaim/models"
)

func TestZaimTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		jsonInput string    // テスト用JSON入力文字列
		expected  time.Time // 期待するtime.Timeの値
		wantErr   bool      // エラー発生の有無
	}{
		{
			jsonInput: `"2023-05-10 15:30:45"`,
			expected:  time.Date(2023, 5, 10, 15, 30, 45, 0, time.UTC),
			wantErr:   false,
		},
		{
			jsonInput: `"invalid time"`,
			expected:  time.Time{},
			wantErr:   true,
		},
		{
			jsonInput: `""`,
			expected:  time.Time{},
			wantErr:   true,
		},
		{
			jsonInput: `null`,
			expected:  time.Time{},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		var zt models.ZaimTime
		err := json.Unmarshal([]byte(tt.jsonInput), &zt)
		if (err != nil) != tt.wantErr {
			t.Errorf("UnmarshalJSON(%s) のエラー = %v, 期待するエラー有無 = %v", tt.jsonInput, err, tt.wantErr)
			continue
		}
		if err == nil && !zt.Time.Equal(tt.expected) {
			t.Errorf("UnmarshalJSON(%s) の結果 = %v, 期待値 = %v", tt.jsonInput, zt.Time, tt.expected)
		}
	}
}
