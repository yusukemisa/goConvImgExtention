package convertor

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

//Convertor is struct
type Convertor struct {
	From string
	To   string
}

var count int

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
	filepath.Walk(targetPath, c.process)
	fmt.Printf("%v件処理しました\n", count)
}

//process はWalkから呼ばれるコールバック関数
func (c *Convertor) process(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error has occurred in process\n%v\n", err.Error())
		return err
	}
	//ディレクトリの場合スルー
	if info.IsDir() {
		return nil
	}
	//変換対象の拡張子でなければスルー
	ext := strings.TrimLeft(filepath.Ext(path), ".")
	if ext != c.From {
		return nil
	}

	//画像形式変換
	imgFileFrom, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v could not open:%v\n", path, err.Error())
		return err
	}
	defer imgFileFrom.Close()

	imgFrom, _, err := image.Decode(imgFileFrom)
	if err != nil {
		//変換対象の拡張子を持つファイルがデコードできない場合はスキップする
		fmt.Fprintf(os.Stderr, "%v could not decode:%v\n", path, err.Error())
		return nil
	}
	//変換先のファイルの作成
	rename := strings.TrimRight(path, c.From) + c.To
	imgFileTo, err := os.Create(rename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v could not create:%v\n", rename, err.Error())
		return err
	}
	defer imgFileTo.Close()

	/*
		imageExtention.goにサポートする画像形式を定義したけど
		ここの分岐でベタ書きしちゃってるので活かせてない・・・
	*/
	switch ext {
	case "jpeg", "jpg":
		err = jpeg.Encode(imgFileTo, imgFrom, nil)
	case "gif":
		err = gif.Encode(imgFileTo, imgFrom, nil)
	case "png":
		err = png.Encode(imgFileTo, imgFrom)
	}
	fmt.Printf("convert %v -> %v\n", path, rename)
	count++
	return err
}
