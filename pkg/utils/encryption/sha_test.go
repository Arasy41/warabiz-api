package encryption

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {

	text := "putra1business@gmail.com:SuretyBond2023!"
	hash := GenerateSHA256(text)
	fmt.Println(hash)
}
