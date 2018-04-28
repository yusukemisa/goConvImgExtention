[![Build Status](https://travis-ci.org/yusukemisa/goConvImgExtention.svg?branch=master)](https://travis-ci.org/yusukemisa/goConvImgExtention)
[![codecov](https://codecov.io/gh/yusukemisa/goConvImgExtention/branch/master/graph/badge.svg)](https://codecov.io/gh/yusukemisa/goConvImgExtention)
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
* [x] テストコードを作ってみる
* [x] テストコード書きやすいようにリファクタ(独立性を高める)
* [x] テストコードのパッケージは「{test対象package名}_test」にする
* [ ] Exampleテストを作成する
* [x] サブテストを作成する
* [x] テーブル駆動テストを作成する
* [x] カバレッジを取得、可視化する


* [x] Travis CIで自動テストされるようにしてみる


## Packages
* convertor 画像形式を変換する機能を持つパッケージ
[![GoDoc](https://godoc.org/github.com/yusukemisa/goConvImgExtention/convertor?status.svg)](https://godoc.org/github.com/yusukemisa/goConvImgExtention/convertor)


## How to use

```
// get
$go get github.com/yusukemisa/goConvImgExtention

// execute
$goConvImgExtention -f png -t jpg imageDir
convert imageDir/2.png -> imageDir/2.jpg
convert imageDir/depth1/depth2/111.png -> imageDir/depth1/depth2/111.jpg
```

### テスト実行
```
$ go test {test対象パッケージ名}
```

### Exampleテスト
```
func ExampleHex_String() {
	fmt.Println(mypkg.Hex(10))
	// Output: a
}
```

### サブテストを作成する
サブテストの定義
```
func TestAB(t *testing.T) {
	t.Run("A", func(t *testing.T) { t.Error("error") })
	t.Run("B", func(t *testing.T) { t.Error("error") })
}
```
実行
```
go test -v sample -run TestAB/A
=== RUN   TestAB
=== RUN   TestAB/A
--- FAIL: TestAB (0.00s)
   --- FAIL: TestAB/A (0.00s)
   	sample_test.go:10: error
FAIL
exit status 1
FAIL	sample	0.007s
```
### テーブル駆動テストを作成する

```
var flagtests = []struct {
	in  string
	out string
}{
	{"%a", "[%a]"}, {"%-a", "[%-a]"}, {"%+a", "[%+a]"},
	{"%#a", "[%#a]"}, {"% a", "[% a]"},
}

func TestFlagParser(t *testing.T) {
	var flagprinter flagPrinter
	for _, tt := range flagtests {
		s := Sprintf(tt.in, &flagprinter)
		if s != tt.out {
			t.Errorf("Sprintf(%q, &flagprinter) => %q, want %q", tt.in, s, tt.out)
		}
	}
}
```

### カバレッジ測定
```
$go test -coverprofile=profile github.com/yusukemisa/goConvImgExtention/convertor
ok      github.com/yusukemisa/goConvImgExtention/convertor      0.009s  coverage: 26.4% of statements
$go tool cover -html=profile
```


### テストヘルパー関数
* ヘルパー関数はエラーを返さない
* *testing.Tを受け取ってテストを落とす
* Go 1.9からはT.Helperを使って情報を補足する
  * https://tip.golang.org/pkg/testing/#T.Helper

```
func testTempFile(t *testing.T) string {
 t.Helper()
 tf, err := ioutil.TempFile("", "test")
 if err != nil {
   t.Fatal("err %s", err)
 }
 tf.Close()
 return tf.Name()
}
```
