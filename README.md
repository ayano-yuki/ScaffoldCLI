# Scaffold CLI

このプロジェクトは、**Golang製のCLIコマンドラインツールを作成するためのテンプレート**として設計・開発されました。

Go言語でCLIツールを素早く開発できるように、  
- オニオンアーキテクチャに基づく堅牢な設計  
- CobraやシンプルなCLI入力に対応した拡張しやすいコマンド構成  
- テンプレートによるプロジェクト初期化機能  

を備えています。

---

## 特徴

- CLIツール開発のベースとしてすぐ使える  
- 汎用的なプロジェクト初期化コマンド（`init`）をサポート  
- 将来的に多様なコマンド追加やテンプレート拡張が可能  
- Windows、Linux、macOSで動作  
- 対話形式でテンプレートやプロジェクト名を選択できるUX  

---

## 使い方（対話モード）

```bash
go run main.go init
```

実行後、以下のようなプロンプトが表示されます：

```
Enter project name: MyApp
Select template (go): 
```

選択に応じて `templates/go` などからテンプレートが適用されます。  
拡張子 `.tmpl` のファイルは `.tmpl` を除いた名前で出力されます。

---

## 使い方（オプション指定）

```bash
go run main.go init --name MyApp --template go --output ./MyApp
```

- `--name / -n`: プロジェクト名を指定
- `--template / -t`: 使用するテンプレート名（例: `go`）
- `--output / -o`: プロジェクト生成先ディレクトリ（省略時は `./output`）

---

## インストール方法

### ビルドして実行ファイルを作成

```bash
go build -o scaffold.exe main.go
./scaffold.exe init
```

### `go install` でインストール

```bash
go install github.com/ユーザ名/scaffold@latest
scaffold init
```

> ※ `GOPATH/bin` が `PATH` に含まれていることを確認してください

---

## ディレクトリ構成（アーキテクチャ）

```
.
├── application/usecase       # ユースケース層（InitProjectなど）
├── cmd                       # CLIエントリポイント
├── config                    # 設定管理（将来的に拡張予定）
├── domain
│   ├── model                 # ドメインモデル
│   └── service               # ドメインサービス（template操作）
├── infrastructure/service    # 外部依存の実装（file出力など）
├── interface/cli             # CLIからの入力インターフェース
├── shared                    # 共通定義（エラーハンドリングなど）
└── templates/
    ├── go/                   # goテンプレートの構成ディレクトリ
    └── vue/                  # （将来的に）Vueテンプレートなど
```

---

## テンプレート構成のルール

- `templates/{templateName}/` にテンプレート一式を格納
- ディレクトリ構造はそのままプロジェクトに展開されます
- 拡張子 `.tmpl` のファイルはテンプレートとして扱い `.tmpl` を除いたファイル名で生成
- Goの `text/template` 構文が使用可能（例： `{{ .Name }}` など）

### 例: `main.go.tmpl`

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, {{ .Name }}!")
}
```

このテンプレートは、プロジェクト名に応じて `Hello, MyApp!` のように動的に変化します。

---

## 拡張性

- `templateName` を動的に選択できる仕組み（将来的に Vue/Rustなども追加可）
- `interface/cli` → `application/usecase` → `domain/service` → `infrastructure` のように責務を分離
- オニオンアーキテクチャで依存性の方向が明確
- テスト容易性と保守性を意識した構造

---

## 今後の展望

- 対話UIの改善（例: テンプレート一覧を十字キー選択）
- テンプレート追加（React、Vue、Rust、Next.js など）
- GitHub Actions でのCIテンプレート
- JSON/YAML による設定反映
- カスタム変数の注入 `.Port`, `.Author` など

---

## ライセンス

MIT License
