package convert

import (
	"fmt"
	"testing"
)

func TestGetDefaultChinese(t *testing.T) {
	words := getDefaultChinese()
	fmt.Println(words)
}
