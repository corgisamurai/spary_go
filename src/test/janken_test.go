package test

import (
  "testing"
)

func Judge(my string, you string) string {
  if my == you {
    return "あいこ"
  }
  if (my == "goo" && you == "choki") ||
    (my == "choki" && you == "par") ||
    (my == "par" && you == "goo") {
    return "勝ち"
  }
  return "負け"
}

func TestJudge(t *testing.T) {
  if Judge("goo", "choki") != "勝ち" {
    t.Fatalf("not 勝ち")
  }
  if Judge("par", "choki") != "負け" {
    t.Fatalf("not 負け")
  }
  if Judge("goo", "par") != "負け" {
    t.Fatalf("not 負け")
  }
  if Judge("goo", "goo") != "あいこ" {
    t.Fatalf("not あいこ")
  }
  if Judge("choki", "goo") != "負け" {
    t.Fatalf("not 負け")
  }
  if Judge("choki", "par") != "勝ち" {
    t.Fatalf("not 勝ち")
  }
  if Judge("par", "goo") != "勝ち" {
    t.Fatalf("not 勝ち")
  }
}
