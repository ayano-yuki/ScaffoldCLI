
# 生成ファイルの保存場所変更とテンプレート内変数挿入の仕組み解説

## 1. `templates/go` 以外のパスにテンプレートを置く場合の変更

---

### 現状

- テンプレートは `templates/go/` 以下に配置されています。
- 例： `templates/go/main.go` 、 `templates/go/go.mod.tmpl` など。

### 変更したい内容

- 例えば `templates/main.go.tmpl` のように、テンプレートファイルを直下の `templates/` に置きたい。
- そのため、テンプレートを読むパスの指定を変更する必要がある。

### 変更例（Goコード内）

```go
// 例: 旧パス
tplPath := filepath.Join("templates", "go", "main.go.tmpl")

// 新パス（テンプレート直下にファイルを置く場合）
tplPath := filepath.Join("templates", "main.go.tmpl")
```

---

### ファイル読み込み部分の例

```go
content, err := os.ReadFile(tplPath)
if err != nil {
    return err
}
tpl, err := template.New("main").Parse(string(content))
if err != nil {
    return err
}
```

- ここでのパス指定を適宜変更してください。

---

## 2. プロジェクト名を変更できる仕組み

---

- テンプレート内にプロジェクト名などの変数を埋め込むことで、生成ファイルを動的にカスタマイズ可能にする。

### 使うのは `text/template` パッケージの機能

例：

```go
type TemplateData struct {
    Name string
}

// 使用例
data := TemplateData{Name: "MyProject"}

tpl.Execute(outFile, data)
```

- テンプレートファイル中で

```go
fmt.Println("Hello, {{ .Name }}!")
```

のように `.Name` でアクセス可能。

---

## 3. テンプレートの変数と注意点

---

### 例

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, {{ .Name }}!")
}
```

- ここで、`{{ .Name }}` はテンプレート実行時に `TemplateData.Name` の値で置換される。

### 重要なポイント

- テンプレートで宣言していない変数（例えば `{{ .Age }}` を `TemplateData` に入れていない）を使うとエラーになります。
- 必ずテンプレートに渡すデータ構造で使うフィールドを揃えましょう。

---

## 4. 具体的なコード例

---

### 生成プログラム例

```go
package scaffold

import (
    "os"
    "path/filepath"
    "text/template"
)

type TemplateData struct {
    Name string
}

func GenerateMainGo(projectDir, projectName string) error {
    tplPath := filepath.Join("templates", "main.go.tmpl")

    tplContent, err := os.ReadFile(tplPath)
    if err != nil {
        return err
    }

    tpl, err := template.New("main").Parse(string(tplContent))
    if err != nil {
        return err
    }

    outFilePath := filepath.Join(projectDir, "main.go")
    outFile, err := os.Create(outFilePath)
    if err != nil {
        return err
    }
    defer outFile.Close()

    data := TemplateData{Name: projectName}

    return tpl.Execute(outFile, data)
}
```

---

### `templates/main.go.tmpl` の例

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, {{ .Name }}!")
}
```

---

## 5. まとめ

- **テンプレートの配置場所**はコード内のパス指定を変えるだけで対応可能。
- **テンプレートファイル内の変数**は、`text/template` の `{{ .FieldName }}` 形式で指定し、Go側で渡すデータ構造に対応したフィールドを用意する必要あり。
- テンプレート内で存在しない変数を使うとエラーになるため、テンプレートとデータの整合性を必ず保つこと。
- プロジェクト名などの動的要素は、`TemplateData`構造体のように管理すると拡張もしやすい。

---
