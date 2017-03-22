package main

import (
  "api"
  "batch"
  "fmt"
  "lib"
  "os"
)

func main() {
  lib.LoadEnv()

  if len(os.Args) == 1 {
    api.Run()
  } else if len(os.Args) == 2 {
    funcs := map[string]func(){
      "ImportOnsenList": batch.ImportOnsenList,
    }
    funcs[os.Args[1]]()
  } else {
    fmt.Println("too many args")
  }
}
