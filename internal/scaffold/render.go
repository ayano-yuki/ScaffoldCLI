package scaffold

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TemplateData struct {
	Name string
	Age  int
}

func GenerateProjectFromTemplateDir(templateDir, outputDir string, data TemplateData) error {
	return filepath.WalkDir(templateDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 相対パスを取得
		relPath, err := filepath.Rel(templateDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(outputDir, removeTmplExt(relPath))

		// ディレクトリの場合は作成してスキップ
		if d.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}

		// .tmpl 以外はスキップ
		if !strings.HasSuffix(d.Name(), ".tmpl") {
			return nil
		}

		// ファイル出力
		return renderFile(path, destPath, data)
	})
}

func renderFile(templatePath, outputPath string, data TemplateData) error {
	content, err := os.ReadFile(templatePath)
	if err != nil {
		return err
	}

	tpl, err := template.New(filepath.Base(templatePath)).Parse(string(content))
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
		return err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return tpl.Execute(outFile, data)
}

func removeTmplExt(path string) string {
	if strings.HasSuffix(path, ".tmpl") {
		return strings.TrimSuffix(path, ".tmpl")
	}
	return path
}
