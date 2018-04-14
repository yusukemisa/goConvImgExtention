package convertor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//Convertor is struct
type Convertor struct {
	From string
	To   string
}

//New サポート対象外の拡張子が指定された場合nilを返す
func New(from, to string) *Convertor {
	if ImageExtention[from] && ImageExtention[to] {
		return &Convertor{
			From: from,
			To:   to,
		}
	}
	return nil
}

//Convert はtargetPath配下のファイルの拡張子をFrom->Toに変換します。
//targetPath配下は再帰的に処理されます。
func (c *Convertor) Convert(targetPath string) {
	filepath.Walk(targetPath, c.convFileExtention)
}

//convFileExtention はWalkから呼ばれるコールバック関数
func (c *Convertor) convFileExtention(path string, info os.FileInfo, err error) error {
	//ディレクトリの場合スルー
	if info.IsDir() {
		return nil
	}
	//変換対象の拡張子でなければスルー
	if filepath.Ext(path) != "."+c.From {
		return nil
	}
	rename := strings.TrimRight(path, c.From) + c.To
	fmt.Printf("convert %v -> %v\n", path, rename)
	//リネーム
	return os.Rename(path, rename)
}
