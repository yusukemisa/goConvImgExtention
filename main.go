package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yusukemisa/goConvImgExtention/convertor"
)

var from = flag.String("f", "jpg", "convert from")
var to = flag.String("t", "png", "convert to")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		convertor.PrintSupportImageExtention()
		fmt.Println("<コマンド実行例>")
		fmt.Println("------------------------------------------------")
		fmt.Println("$goConvImgExtention -f jpg -t png {targetDir}")
		fmt.Println("------------------------------------------------")
	}
	flag.Parse()
}

func main() {
	//引数指定なしは終了
	if len(flag.Args()) != 1 {
		fmt.Println("変換対象とするディレクトリを１つ指定してください")
		flag.Usage()
		os.Exit(1)
	}
	rootPath := flag.Args()[0]
	//引数のディレクトリないと終了
	_, err := os.Stat(rootPath)
	if err != nil {
		fmt.Printf("指定されたディレクトリがありません:%v\n", rootPath)
		flag.Usage()
		os.Exit(1)
	}
	//フラグにサポート対象の拡張子が指定されている場合変換実行
	if c := convertor.New(*from, *to); c != nil {
		c.Convert(rootPath)
	} else {
		fmt.Println("サポート対象外の画像形式が指定されています。")
		flag.Usage()
		os.Exit(1)
	}
}
