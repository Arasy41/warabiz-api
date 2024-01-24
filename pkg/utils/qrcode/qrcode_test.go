package qrcode

import (
	"log"
	"testing"

	"warabiz/api/pkg/utils/file"
	"warabiz/api/pkg/utils/qrcode/src"
)

func TestQRCode(t *testing.T) {

	content := src.Content{
		Logo:            "https://storage.googleapis.com/chakra-loyalty/kalcare/Testing/upload/images/productbrand/2190760b-3afd-4aa5-ab3c-fdfcb67d87c4.jpg",
		URL:             "https://lp-teman-prenagen-testing-d33dgvhu5a-as.a.run.app?qr=KC221226nKIx",
		BrandName:       "PRENAGEN",
		SpgName:         "syifa.prenagen",
		SpgSubmissionID: "JOGSFA",
	}

	//* Call configuration
	config := src.SetConfiguration(&content)

	//* Create QRImage
	qrImg, err := CreateQRImage(config)
	if err != nil {
		log.Fatal(err)
	}

	//* Store file in local
	file.CreateJPEGFile(qrImg, "example/qrcode-example.jpeg")
}
