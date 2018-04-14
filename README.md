## goConvImgExtention
* [x] 引数としてディレクトリを指定する
* [x] 指定したディレクトリ以下のJPGファイルをPNGに変換
* [x] ディレクトリ以下は再帰的に処理する
* [x] 変換前と変換後の画像形式を-f,-tオプションで指定できる

## Features
以下を満たすように開発する
* [x] mainパッケージと変換機能のパッケージは分離する
* [x] 自作パッケージと標準パッケージと準標準パッケージのみ使う（準標準パッケージ：golang.org/x以下のパッケージ）
* [x] ユーザ定義型を作ってみる

## Packages
* convertor
** 画像形式を変換する機能を持つパッケージ

## How to use

```
// get
$go get github.com/yusukemisa/goConvImgExtention

// execute
$goConvImgExtention -f png -t webp imageDir
convert imageDir/2.png -> imageDir/2.webp
convert imageDir/depth1/depth2/111.png -> imageDir/depth1/depth2/111.webp
```
