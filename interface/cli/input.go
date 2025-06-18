package cli

import "fmt"

type InitInput struct {
	Name         string
	TemplateType string
	OutputDir    string
}

// 必要であれば、ここで入力チェックなどを加える
func (i InitInput) Validate() error {
	if i.Name == "" {
		return fmt.Errorf("project name is required")
	}
	return nil
}
