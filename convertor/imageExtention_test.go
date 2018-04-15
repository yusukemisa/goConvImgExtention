package convertor

import (
	"testing"
)

func Test_keys(t *testing.T) {
	actual := keys(ImageExtention)
	if len(ImageExtention) != len(actual) {
		t.Error("期待するサイズではありません")
	}
	//比較元はmapに格納された情報なので並び順不定。なので含まれているかどうかで判断する
	for _, v := range actual {
		if !ImageExtention[v] {
			t.Error("期待する拡張子の定義と一致しません。")
		}
	}
}
