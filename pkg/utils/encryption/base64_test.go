package encryption

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

func TestBase64(t *testing.T) {
	value := "4c321387f4f4fbdccba734f80239dfee0b041301aa23d1df5ea755c98d727642"
	b64 := EncryptBase64(value)

	fmt.Println(b64)

	res, err := DecryptBase64(b64)
	if err != nil {
		log.Fatal(err)
	}
	if res != value {
		log.Fatal(errors.New("result not equal"))
	}
}
