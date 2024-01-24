package encryption

import (
	"fmt"
	"log"
	"testing"
)

func TestPasword(t *testing.T) {
	password := "SuretyBond2023!"
	hash, err := EncryptPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hash)

	err = ComparePassword(hash, password)
	if err != nil {
		log.Fatal(err)
	}
}
