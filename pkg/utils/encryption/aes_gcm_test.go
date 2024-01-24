package encryption

import (
	"encoding/json"
	"fmt"
	"testing"
)

type RegistrationRequestRaw struct {
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	EmailOTP        string `json:"email_otp" validate:"required"`
	Whatsapp        string `json:"whatsapp" validate:"omitempty,phone"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

func TestAESGCM(t *testing.T) {

	data := RegistrationRequestRaw{
		FirstName:       "putra",
		LastName:        "rama",
		Email:           "putra1business@gmail.com",
		EmailOTP:        "582347",
		Password:        "SuretyBond2023!",
		ConfirmPassword: "SuretyBond2023!",
		Whatsapp:        "",
	}

	dataRaw, _ := json.Marshal(data)

	secretKey := "isth1sb0nd5ur3ty"

	encryptedData, err := EncryptAESGCM(string(dataRaw), secretKey)
	if err != nil {
		fmt.Println("Encrption error:", err)
		return
	}

	fmt.Println("Encrypted Data:", encryptedData)
	// fmt.Println("iv:", iv)

	// encryptedData := "rPGKTXfs5XYtEsXbkZ6Jts8oq365Jh1/FXLVX6M="
	// iv := "ZTptyWIXW/jjznqu"

	decryptedData, err := DecryptAESGCM(encryptedData, secretKey)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}

	fmt.Println("Decrypted Data:", decryptedData)

}
