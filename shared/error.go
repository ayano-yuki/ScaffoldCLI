package shared

import "errors"

var (
	// ErrNotFound は対象が見つからなかったエラーの例
	ErrNotFound = errors.New("not found")

	// ErrInvalidInput は入力値が不正な場合のエラー例
	ErrInvalidInput = errors.New("invalid input")
)
