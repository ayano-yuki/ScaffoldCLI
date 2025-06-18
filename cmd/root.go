package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"scaffold/application/usecase"
	"scaffold/domain/model"
	domainservice "scaffold/domain/service"
	infrastructure "scaffold/infrastructure/service"
)

var rootCmd = &cobra.Command{
	Use:   "scaffold",
	Short: "A simple project scaffolding tool",
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		templateType, _ := cmd.Flags().GetString("template")
		outputDir, _ := cmd.Flags().GetString("output")

		if name == "" || templateType == "" || outputDir == "" {
			return fmt.Errorf("all of --name, --template, and --output are required")
		}

		project := model.Project{
			Name:         name,
			TemplateType: templateType,
			OutputDir:    outputDir,
		}

		initUsecase := usecase.NewInitProjectUsecase(
			domainservice.NewTemplateService(),
			infrastructure.NewFileWriter(),
		)
		return initUsecase.Execute(project)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	initCmd.Flags().StringP("name", "n", "", "Project name")
	initCmd.Flags().StringP("template", "t", "go", "Template type (e.g., go)")
	initCmd.Flags().StringP("output", "o", "./output", "Output directory")

	rootCmd.AddCommand(initCmd)
}
