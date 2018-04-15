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
		fmt.Fprintln(os.Stderr, "<コマンド実行例>")
		fmt.Fprintln(os.Stderr, "------------------------------------------------")
		fmt.Fprintln(os.Stderr, "$goConvImgExtention -f jpg -t png {targetDir}")
		fmt.Fprintln(os.Stderr, "------------------------------------------------")
	}
	flag.Parse()
}

func main() {
	//引数指定なしは終了
	if len(flag.Args()) != 1 {

		fmt.Fprintln(os.Stderr, "変換対象とするディレクトリを１つ指定してください")
		flag.Usage()
		os.Exit(1)
	}
	rootPath := flag.Args()[0]
	//引数のディレクトリないと終了
	_, err := os.Stat(rootPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "指定されたディレクトリがありません:%v\n", rootPath)
		flag.Usage()
		os.Exit(1)
	}
	//フラグにサポート対象の拡張子が指定されている場合変換実行
	if c := convertor.New(*from, *to); c != nil {
		c.Convert(rootPath)
	} else {
		fmt.Fprintln(os.Stderr, "サポート対象外の画像形式が指定されています。")
		flag.Usage()
		os.Exit(1)
	}
}
