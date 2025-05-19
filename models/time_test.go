package models_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/ryohidaka/go-zaim/models"
)

func TestZaimTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		jsonInput string    // テスト用JSON入力文字列
		expected  time.Time // 期待するtime.Timeの値
		wantErr   bool      // エラー発生の有無
	}{
		{
			name:      "日時形式",
			jsonInput: `"2023-05-10 15:30:45"`,
			expected:  time.Date(2023, 5, 10, 15, 30, 45, 0, time.UTC),
			wantErr:   false,
		},
		{
			name:      "日付のみ形式",
			jsonInput: `"2023-05-10"`,
			expected:  time.Date(2023, 5, 10, 0, 0, 0, 0, time.UTC),
			wantErr:   false,
		},
		{
			name:      "無効な文字列",
			jsonInput: `"invalid time"`,
			expected:  time.Time{},
			wantErr:   true,
		},
		{
			name:      "空文字列",
			jsonInput: `""`,
			expected:  time.Time{},
			wantErr:   true,
		},
		{
			name:      "null",
			jsonInput: `null`,
			expected:  time.Time{},
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var zt models.ZaimTime
			err := json.Unmarshal([]byte(tt.jsonInput), &zt)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON(%s) エラー = %v, 期待値 = %v", tt.jsonInput, err, tt.wantErr)
				return
			}
			if err == nil && !zt.Time.Equal(tt.expected) {
				t.Errorf("UnmarshalJSON(%s) 結果 = %v, 期待値 = %v", tt.jsonInput, zt.Time, tt.expected)
			}
		})
	}
}
