package convertor

import (
	"testing"
)

func Test_New_empty(t *testing.T) {
	if actual := New("", ""); actual != nil {
		t.Error("nilにならないとおかしい")
	}

}

func Test_New_not_supported(t *testing.T) {
	if actual := New("bmp", "png"); actual != nil {
		t.Error("サポート対象外拡張子なのでnilになるべき")
	}
}

func Test_New_supported(t *testing.T) {
	if actual := New("jpg", "png"); actual == nil {
		t.Error("サポート対象拡張子なのでnilになるとおかしい")
	} else {
		if actual.From != "jpg" || actual.To != "png" {
			t.Error("引数->構造体へのマッピングが意図した通りになってない")
		}
	}
}
