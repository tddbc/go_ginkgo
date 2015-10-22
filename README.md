# TDDBC GO Project
_このプロジェクトは、TDDBC(Test Driven Development Boot Camp)で使用する GO 言語用のプロジェクトテンプレートです。_

## はじめに
サンプルプロジェクトでは、HelloWorldクラスを作成し、Sayメソッドが呼ばれたときに"Hello TDD BootCamp!!"を返すプログラムがついてきます。

## セットアップ
### GO言語のセットアップ
[本家サイト](https://golang.org/doc/install)を参考に環境をセットアップしてください。

**注意**
セットアップの際、環境変数に GOPATH の設定を忘れてしまいうケースが多いようです。
また、$GOPATH/bin を環境変数 PATH に含めてください。


#### MacPorts で簡単セットアップ
port コマンドを実行して、GO言語をインストールすることができます。
コマンドは、以下の通りです。

```bash
sudo port selfupdate
sudo port upgrade outdated 
sudo port install go
```

#### Homebrew で簡単セットアップ
brew コマンドを実行して、GO言語をインストールすることができます。
コマンドは、以下の通りです。

```bash
sudo brew update
sudo brew install go
```

#### apt-get で簡単セットアップ
apt-get コマンドを実行して、GO言語をインストールすることができます。
コマンドは、以下の通りです。

```bash
sudo apt-get update
sudo apt-get install go
```

### テスティングフレームワークのセットアップ
テスティングフレームワークに [ginkgo](http://onsi.github.io/ginkgo/)を使用しています。
テストをはじめる前に　go get コマンドを使ってセットアップしてください。

```bash
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega
```

### TDDテンプレートのセットアップ
go get コマンドを実行して TDDBC テンプレートをセットアップしてください。

```bash
go get github.com/tddbc/go_ginkgo
```

## テストの進め方（テンプレートと同じところまで）
1. テスト環境を作る

   `<user>`と`<project name>`は、各自の環境に読み替えてください。
  
   ```bash
   mkdir -p $GOPATH/src/github.com/<user>/<project name>
   cd $GOPATH/src/github.com/<user>/<project name>
   ginkgo bootstrap
   ginkgo generate sample 
   ```

2. テストケースを書く

   ```sample_test.go
   package go_ginkgo_test
   
   import (
       . "github.com/tddbc/go_ginkgo"
   
       . "github.com/onsi/ginkgo"
       . "github.com/onsi/gomega"
   )
   
   var _ = Describe("Sample", func() {
       Describe("#say", func() {
           It("Hello TDD BootCampを返すこと", func() {
               Expect(Say()).To(Equal("Hello TDD BootCamp!!"))
           })
       })
   })
   ```

3. 実装コードを書く

   ```sample.go
   package go_ginkgo
   
   func Say() string {
       return ""
   }
   ```

4. テストを実行する
   ```bash
$ ginkgo
~~~ 詳細略 ~~~
Summarizing 1 Failure:

[Fail] Sample #say [It] Hello TDD BootCampを返すこと
C:/Users/ALSI30001059/Documents/GitHub/src/github.com/tddbc/go_ginkgo/sample_test.go:13

Ran 1 of 1 Specs in 0.019 seconds
FAIL! -- 0 Passed | 1 Failed | 0 Pending | 0 Skipped --- FAIL: TestGoGinkgo (0.02s)
FAIL

Ginkgo ran 1 suite in 6.424s
Test Suite Failed
   ```

5. テストが通るコードを書く

   ```sample.go
   package go_ginkgo
   
   func Say() string {
       return "Hello TDD BootCamp!!"
   }
   ```

6. テストを実行する

   ```bash
$ ginkgo
Running Suite: GoGinkgo Suite
=============================
Random Seed: 1445227273
Will run 1 of 1 specs

•
Ran 1 of 1 Specs in 0.000 seconds
SUCCESS! -- 1 Passed | 0 Failed | 0 Pending | 0 Skipped PASS

Ginkgo ran 1 suite in 4.917s
Test Suite Passed
   ```

## おまけ
[ginkgo](http://onsi.github.io/ginkgo/)コマンドは、便利な機能がいくつかあります。
よく使うであろうパラメータの使い方を書いておきます。

### パラメータ
* watch

   ファイルの更新されると自動でテストを実行してくれます。
   テストの度にコマンドを実行するのが面倒な方は、利用するといいかも

* -r

   指定したディレクトリの配下のディレクトリを対象にテストを実行します。
   ただし、各ディレクトリで get bootstrap コマンドの実行が必要です。

* -cover

   カバレッジを測定します。どのくらい網羅されているかパーセンテージで表示してくれます。

* -notify

   OS X と Linux で対応しています。テストが完了するとOSにポップアップを表示してくれます。
   OS Xの場合は、terminal-notifier。Linux の場合は、notify-send　そインストールしておく必要があります。


## License
Copyright (c) 2015 TDD BootCamp and other contributors

http://devtesting.jp/tddbc/

https://github.com/tddbc

Licensed under the MIT license.

Created by [135yshr](https://github.com/135yshr)
