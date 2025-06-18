package usecase

import (
	"fmt"
	"os"
	"path/filepath"

	"scaffold/domain/model"
	"scaffold/domain/service"
	infra "scaffold/infrastructure/service"
)

type InitProjectUsecase struct {
	templateService service.TemplateService
	fileWriter      infra.FileWriter
}

func NewInitProjectUsecase(
	templateService service.TemplateService,
	fileWriter infra.FileWriter,
) *InitProjectUsecase {
	return &InitProjectUsecase{
		templateService: templateService,
		fileWriter:      fileWriter,
	}
}

// Execute プロジェクト初期化処理
func (u *InitProjectUsecase) Execute(project model.Project) error {
	// ① テンプレートディレクトリの特定
	templateRoot := filepath.Join("templates", project.TemplateType)

	// ② テンプレート内のファイルを再帰的に走査
	err := filepath.Walk(templateRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// ディレクトリはスキップ
		if info.IsDir() {
			return nil
		}

		// ③ ファイル読み込み
		contentBytes, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read template file: %w", err)
		}

		// ④ テンプレートレンダリング
		rendered, err := u.templateService.Render(string(contentBytes), project)
		if err != nil {
			return fmt.Errorf("failed to render template: %w", err)
		}

		// ⑤ 出力先パス計算（テンプレートルートをベースに相対パス化）
		relPath, err := filepath.Rel(templateRoot, path)
		if err != nil {
			return err
		}
		outputPath := filepath.Join(project.OutputDir, relPath)

		// ⑥ ファイル書き出し
		if err := u.fileWriter.WriteFile(outputPath, rendered); err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}

		return nil
	})

	return err
}
