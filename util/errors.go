package util

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Errors : 複数のエラーを一括するエラー
type Errors []error

// SPush : 文字列でエラーを作成し、追加する
func (a *Errors) SPush(text string) {
	a.Push(errors.New(text))
}

// Pushf : 書式付き文字列でエラーを作成し、追加する
func (a *Errors) Pushf(format string, p ...interface{}) {
	a.Push(fmt.Errorf(format, p...))
}

// Push : エラーを追加する
func (a *Errors) Push(item error) {
	if item == nil {
		return
	}
	if arr, ok := item.(Errors); ok && arr != nil {
		for _, child := range arr {
			a.Push(child)
		}
	} else {
		*a = append(*a, item)
	}
}

func (a Errors) Len() int      { return len(a) }
func (a Errors) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Errors) Less(i, j int) bool {
	return a[i].Error() < a[j].Error()
}

// IsEmpty : エラーが空かどうか取得する
func (a *Errors) IsEmpty() bool {
	return a == nil || len(*a) == 0
}

// Error : エラー出力を得る
func (a Errors) Error() string {
	array := make([]string, len(a))
	for i, item := range a {
		array[i] = item.Error()
	}
	buf, _ := json.MarshalIndent(array, "", "  ")
	return string(buf)
}

// Err : 空の場合は nil 、それ以外の場合は自身を返す
func (a Errors) Err() error {
	if a.IsEmpty() {
		return nil
	}
	return a
}
