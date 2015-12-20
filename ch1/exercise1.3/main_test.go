package main

import "testing"

var args []string = []string{"main_test", "a", "b", "c", "d", "f"}

func BenchmarkEchoSlow(b *testing.B) {
  for i := 0; i < b.N; i++ {
    echoSlow(args)
  }
}

func BenchmarkEchoFast(b *testing.B) {
  for i := 0; i < b.N; i++ {
    echoFast(args)
  }
}
