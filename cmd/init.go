package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"io/ioutil"
	"scaffold/application/usecase"

	"github.com/manifoldco/promptui"
)

func InitCommand() error {
	reader := bufio.NewReader(os.Stdin)

	// プロジェクト名の入力
	fmt.Print("Enter project name: ")
	projectName, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	projectName = strings.TrimSpace(projectName)

	// テンプレート一覧の取得
	templateType, err := selectTemplate()
	if err != nil {
		return err
	}

	// プロジェクト初期化の実行
	return usecase.InitProject(projectName, templateType)
}

// テンプレート選択プロンプト
func selectTemplate() (string, error) {
	const templatesDir = "templates"

	entries, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		return "", fmt.Errorf("failed to read templates directory: %w", err)
	}

	var templates []string
	for _, entry := range entries {
		if entry.IsDir() {
			templates = append(templates, entry.Name())
		}
	}

	if len(templates) == 0 {
		return "", fmt.Errorf("no templates found in '%s'", templatesDir)
	}

	prompt := promptui.Select{
		Label: "Select a template",
		Items: templates,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return result, nil
}
