package env

import (
	"fmt"
	"os"
	"strings"
)

func Setup() {
	b, err := os.ReadFile(".env")
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		parts := strings.Split(line, "=")
		os.Setenv(parts[0], parts[1])
	}

	fmt.Printf("Environment has been set:\n\n%s\n\n", string(b))
}
