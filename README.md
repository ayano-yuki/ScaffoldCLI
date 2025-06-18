# Scaffold CLI
このプロジェクトは、**Golang製のCLIコマンドラインツールを作成するためのテンプレート**として設計・開発されました。

Go言語でCLIツールを素早く開発できるように、  
- オニオンアーキテクチャに基づく堅牢な設計  
- Cobraによる拡張しやすいコマンド実装  
- テンプレートによるプロジェクト初期化機能  

を備えています。


## 特徴
- CLIツール開発のベースとしてすぐ使える  
- 汎用的なプロジェクト初期化コマンド（`init`）をサポート  
- 将来的に多様なコマンド追加やテンプレート拡張が可能  
- Windows、Linux、macOSで動作


## インストール
初期状態は、Golangで「Hello, {プロジェクト名}!」と出力するプロジェクトを作成します。

### ソースから実行
```bash
# scaffold init --name プロジェクト名 --template テンプレート名 --output 出力先ディレクトリ
go run main.go init --name MyApp --template go --output ./MyApp
```
- `-name / -n`: プロジェクト名を指定
- `-template / -t`: 使用するテンプレート（現状は go のみ）
- `-output / -o`: プロジェクト生成先ディレクトリ（省略時は ./output）

### ビルドして実行ファイルを作成
```bash
go build -o scaffold.exe main.go
./scaffold.exe init -n MyApp -t go -o ./MyApp
```

### go install でインストール
```bash
go install github.com/ユーザ名/scaffold@latest
scaffold init -n MyApp -t go -o ./MyApp
```
※ GOPATH/bin が PATH に含まれていることを確認してください


## ディレクトリ構成
```bash
.
├── application/usecase       # ユースケース層
├── cmd                      # CLIエントリポイント
├── config                   # 設定管理
├── domain                   # ドメインモデルとドメインサービス
├── infrastructure           # 外部依存の実装（ファイル書き込みなど）
├── interface/cli            # CLIインターフェース層
├── shared                   # 共通ライブラリ・エラー定義
└── templates/go             # Goプロジェクトテンプレート
```

## 設計思想
- オニオンアーキテクチャを採用し、依存性逆転を徹底
- コマンドごとにユースケースを分離しテスト容易性を確保
- templates ディレクトリにテンプレートを分離し将来の拡張を見据えた構成
- ドメイン層とインフラ層を明確に分離


## テンプレート管理
- templates/go 以下にGoプロジェクトの雛形をテンプレートとして管理


## 展望
- React、Vue、Rust など多様なテンプレートを出力するコマンドライン作成ができる
- YAMLやJSONによる設定ファイルのサポートができる
- CI/CD連携やGitHub Actionsによる自動リリースをできるようにする

