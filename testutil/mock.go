package testutil

import (
	"fmt"
	"io"
	"os"

	"github.com/jarcoal/httpmock"
)

// MockResponseFromFile は、指定された URL パスおよび HTTP メソッド（GET など）に対して、
// 外部の JSON ファイルから API レスポンスをモック（模擬）する
func MockResponseFromFile(method, url, path string) error {
	// JSON ファイルを開く
	file, err := os.Open("testutil/fixtures/json/" + path + ".json")
	if err != nil {
		return fmt.Errorf("モックレスポンスファイルを開けませんでした: %v", err)
	}
	defer file.Close()

	// ファイルの内容を読み込む
	data, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("モックレスポンスファイルを読み取れませんでした: %v", err)
	}

	// 指定された URL パスに対して、GET メソッドのモックレスポンスを登録
	httpmock.RegisterResponder(method, url,
		httpmock.NewStringResponder(200, string(data)))

	return nil
}
