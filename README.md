## goConvImgExtention
* [ ] 引数としてディレクトリを指定する
* [ ] 指定したディレクトリ以下のJPGファイルをPNGに変換
* [ ] ディレクトリ以下は再帰的に処理する
* [ ] 変換前と変換後の画像形式を指定できる

## Features
以下を満たすように開発する
* [ ] mainパッケージと変換機能のパッケージは分離する
* [ ] 自作パッケージと標準パッケージと準標準パッケージのみ使う
* [ ] 準標準パッケージ：golang.org/x以下のパッケージ
* [ ] ユーザ定義型を作ってみる
* [ ] GoDocを生成してみる

## How to use

```
// get
$go get github.com/yusukemisa/goConvImgExtention

// before convert
$ls imageDir
1.jpg	2.jpg	3.jpg

// execute
$ goConvImgExtention imageDir

// result
$ls imageDir
1.png	2.png	3.png
```
