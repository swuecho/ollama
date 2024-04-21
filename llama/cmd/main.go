package main

import (
	"fmt"

	"github.com/ollama/ollama/llama"
)

func main() {
	fmt.Println(llama.SystemInfo())
	llama.Run("gemma-7b-it.gguf", "Hello, world!")
}