package test

import (
  "batch"

  "testing"
)

func TestImportOnsenList(t *testing.T) {
  // 実行前のデータベースにある温泉データの数を取得
  oldResult := execShowSpaList()
  oldSize := len(oldResult.Spas)

  // 1件データを挿入
  batch.ImportOnsenList()

  // 実行後のデータベースにある温泉データの数を取得
  newResult := execShowSpaList()
  newSize := len(newResult.Spas)

  if newSize != (oldSize + 1) {
    t.Fatal()
  }
}
