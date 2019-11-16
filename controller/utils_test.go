package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestControllerUtils(t *testing.T) {
	t.Run("trimNewlineChar()", func(t *testing.T) {
		t.Parallel()
		t.Run("改行や余分な空白が取り除かれているテスト", func(t *testing.T) {
			t.Parallel()

			expected := "hoge fuga piyo"
			actual := trimNewlineChar(`hoge
fuga
    piyo
`)
			assert.Equal(t, expected, actual)
		})
	})
}
