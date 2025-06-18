
# runコマンド追加フルセット実装例

## 1. domain/service/run_service.go

```go
package service

import "fmt"

// RunService はプロジェクトの実行ロジックを提供します。
type RunService struct{}

// NewRunService は RunService のコンストラクタです。
func NewRunService() *RunService {
    return &RunService{}
}

// Run は指定されたプロジェクト名で処理を実行します。
func (s *RunService) Run(projectName string) error {
    if projectName == "" {
        return fmt.Errorf("project name must not be empty")
    }
    // 実際の実行ロジックはここに実装
    fmt.Printf("Running project: %s\n", projectName)
    return nil
}
```

## 2. infrastructure/service/run_executor.go

```go
package service

import (
    "fmt"
    "os/exec"
)

// RunExecutor は外部コマンド実行などの具体的処理を担当します。
type RunExecutor struct{}

// NewRunExecutor は RunExecutor のコンストラクタです。
func NewRunExecutor() *RunExecutor {
    return &RunExecutor{}
}

// Execute はプロジェクトの実行処理例（外部コマンドの呼び出しなど）を実装します。
func (r *RunExecutor) Execute(projectName string) error {
    // ここは例。実際はビルドや実行コマンドをここで実装
    fmt.Printf("Executing project: %s\n", projectName)

    // 例えば `go run main.go` を実行する場合（適宜変更）
    cmd := exec.Command("go", "run", "main.go")
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to execute project: %w\noutput: %s", err, string(output))
    }

    fmt.Printf("Execution output:\n%s\n", string(output))
    return nil
}
```

## 3. application/usecase/run_project.go

```go
package usecase

import (
    "yourmodule/domain/service"
    infra "yourmodule/infrastructure/service"
)

type RunProjectUsecase struct {
    runService  *service.RunService
    runExecutor *infra.RunExecutor
}

func NewRunProjectUsecase(rs *service.RunService, re *infra.RunExecutor) *RunProjectUsecase {
    return &RunProjectUsecase{
        runService:  rs,
        runExecutor: re,
    }
}

// RunProject はプロジェクト実行のユースケース処理
func (uc *RunProjectUsecase) RunProject(projectName string) error {
    if err := uc.runService.Run(projectName); err != nil {
        return err
    }
    if err := uc.runExecutor.Execute(projectName); err != nil {
        return err
    }
    return nil
}
```

## 4. interface/cli/run_input.go

```go
package cli

import (
    "fmt"

    "github.com/spf13/cobra"
)

func NewRunCommand(runFunc func(string) error) *cobra.Command {
    var projectName string

    cmd := &cobra.Command{
        Use:   "run",
        Short: "Run the specified project",
        RunE: func(cmd *cobra.Command, args []string) error {
            if projectName == "" {
                return fmt.Errorf("project name is required")
            }
            return runFunc(projectName)
        },
    }

    cmd.Flags().StringVarP(&projectName, "name", "n", "", "Project name to run")
    return cmd
}
```

## 5. cmd/root.go

```go
package cmd

import (
    "github.com/spf13/cobra"

    "yourmodule/application/usecase"
    "yourmodule/domain/service"
    infra "yourmodule/infrastructure/service"
    "yourmodule/interface/cli"
)

func NewRootCommand() *cobra.Command {
    rootCmd := &cobra.Command{
        Use:   "scaffold",
        Short: "CLI tool scaffold",
    }

    // 既存 init コマンド登録など

    // run用のユースケース初期化
    runUsecase := usecase.NewRunProjectUsecase(service.NewRunService(), infra.NewRunExecutor())
    runCmd := cli.NewRunCommand(runUsecase.RunProject)
    rootCmd.AddCommand(runCmd)

    return rootCmd
}
```

## 6. application/usecase/run_project_test.go

```go
package usecase

import (
    "errors"
    "testing"

    "github.com/stretchr/testify/assert"
)

type mockRunService struct {
    err error
}

func (m *mockRunService) Run(name string) error {
    return m.err
}

type mockRunExecutor struct {
    err error
}

func (m *mockRunExecutor) Execute(name string) error {
    return m.err
}

func TestRunProject_Success(t *testing.T) {
    rs := &mockRunService{}
    re := &mockRunExecutor{}

    uc := NewRunProjectUsecase(rs, re)
    err := uc.RunProject("MyProject")

    assert.NoError(t, err)
}

func TestRunProject_FailRunService(t *testing.T) {
    rs := &mockRunService{err: errors.New("run service error")}
    re := &mockRunExecutor{}

    uc := NewRunProjectUsecase(rs, re)
    err := uc.RunProject("MyProject")

    assert.Error(t, err)
    assert.Equal(t, "run service error", err.Error())
}

func TestRunProject_FailRunExecutor(t *testing.T) {
    rs := &mockRunService{}
    re := &mockRunExecutor{err: errors.New("executor error")}

    uc := NewRunProjectUsecase(rs, re)
    err := uc.RunProject("MyProject")

    assert.Error(t, err)
    assert.Equal(t, "executor error", err.Error())
}
```

## 7. 動作確認

```powershell
go build -o scaffold.exe ./cmd
.\scaffold.exe run --name MyProject
```

期待される出力例：

```
Running project: MyProject
Executing project: MyProject
...（外部コマンドの出力など）...
```

---

# 注意

- `yourmodule` はご自身のGoモジュール名に置き換えてください。
- 実際の外部コマンド呼び出しは `run_executor.go` の `Execute` メソッドでカスタマイズしてください。

---

