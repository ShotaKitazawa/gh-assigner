package image

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	testDir  = "/tmp/images/"
	testFile = "test.txt"
)

func TestMain(m *testing.M) {
	// テスト用ディレクトリ作成
	if err := os.MkdirAll(testDir, 0777); err != nil {
		panic(err)
	}
	// テスト用ファイル作成
	if _, err := os.Create(testDir + testFile); err != nil {
		panic(err)
	}

	code := m.Run()

	// テスト用ディレクトリ削除
	if err := os.RemoveAll(testDir); err != nil {
		panic(err)
	}

	os.Exit(code)
}

func TestImageInfrastructure(t *testing.T) {
	// Initialize
	t.Parallel()
	infra := &ImageInfrastructure{
		Path: testDir,
	}

	t.Run("CreateGraphWithReviewWaitTime()", func(t *testing.T) {
		// Initialize
		t.Parallel()

		t.Run("Imageが作成されることのテスト", func(t *testing.T) {
			// Initialize
			t.Parallel()

			// メソッド呼び出し
			filepath, err := infra.CreateGraphWithReviewWaitTime(map[uint]time.Duration{1: time.Second})

			// 関数が正常終了してるか
			assert.Nil(t, err)

			// ファイルの存在確認
			_, err = os.Stat(filepath)
			assert.Nil(t, err)
		})
		t.Run("イメージ配置先ディレクトリと同名のファイルが既に存在する場合エラーすることのテスト", func(t *testing.T) {
			// Initialize
			t.Parallel()
			path := testDir + "hoge.txt"

			// テスト用ファイル作成
			if _, err := os.Create(path); err != nil {
				panic(err)
			}

			// メソッド呼び出し
			infra = &ImageInfrastructure{
				Path: path,
			}
			_, err := infra.CreateGraphWithReviewWaitTime(map[uint]time.Duration{1: time.Second})

			// エラーが発生することのテスト
			assert.NotNil(t, err)
		})
		t.Run("渡すデータの長さが0の場合エラーすることのテスト", func(t *testing.T) {
			// Initialize
			t.Parallel()

			// メソッド呼び出し
			_, err := infra.CreateGraphWithReviewWaitTime(map[uint]time.Duration{})

			// 関数が異常終了してるか
			assert.NotNil(t, err)
		})
	})
	t.Run("DeleteFile()", func(t *testing.T) {
		// Initialize
		t.Parallel()

		t.Run("ファイルが削除されることのテスト", func(t *testing.T) {
			// Initialize
			t.Parallel()
			path := testDir + testFile

			// メソッド呼び出し
			err := infra.DeleteFile(path)

			// 関数が正常終了してるか
			assert.Nil(t, err)

			// ファイルの存在確認
			_, err = os.Stat(path)
			assert.NotNil(t, err)
		})
	})
}
