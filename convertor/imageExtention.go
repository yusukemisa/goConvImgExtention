package convertor

import (
	"fmt"
	"os"
)

//ImageExtention サポートする画像形式
var ImageExtention = map[string]bool{
	"jpg":  true,
	"jpeg": true,
	"png":  true,
	"webp": true,
	"gif":  true,
}

//PrintSupportImageExtention サポートする画像形式出力
func PrintSupportImageExtention() {
	fmt.Fprintf(os.Stderr, "<サポートする画像形式>\n%v\n", keys(ImageExtention))
}

//ImageExtentionのキーだけ取り出す
var keys = func(m map[string]bool) []string {
	ks := []string{}
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
