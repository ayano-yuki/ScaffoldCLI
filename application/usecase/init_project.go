package usecase

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// InitProject は、テンプレートディレクトリ配下のファイル・ディレクトリ構造を
// そのままコピーしつつ、Goのテンプレートとして展開してプロジェクトを初期化します。
func InitProject(projectName, templateName string) error {
	templateDir := filepath.Join("templates", templateName)
	destDir := filepath.Join(".", projectName)

	data := map[string]interface{}{
		"Name": projectName,
	}

	return copyTemplateDir(templateDir, destDir, data)
}

// copyTemplateDir は templateDir 配下のすべてのファイル・フォルダを再帰的に走査し、
// destDir に同じ構造でテンプレート展開しながら生成します。
func copyTemplateDir(templateDir, destDir string, data map[string]interface{}) error {
	return filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(templateDir, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(destDir, relPath)

		if info.IsDir() {
			// ディレクトリを作成
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		// ファイルの拡張子が .tmpl なら除去する
		if strings.HasSuffix(targetPath, ".tmpl") {
			targetPath = strings.TrimSuffix(targetPath, ".tmpl")
		}

		tmpl, err := template.ParseFiles(path)
		if err != nil {
			return err
		}

		f, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer f.Close()

		return tmpl.Execute(f, data)
	})
}
