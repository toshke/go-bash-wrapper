package main

import (
  "log"
  "os"
  "os/exec"
)

func main() {
  env := append(os.Environ(),"GO_SH_WRAPPER=1")
  args := os.Args[1:]
  cmd := exec.Command("/bin/bash", args...)
  cmd.Env = env
  cmd.Stderr = os.Stderr
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  if err := cmd.Run(); err != nil {
    log.Fatal(err)
  }
}
