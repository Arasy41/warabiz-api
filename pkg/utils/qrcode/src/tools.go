package src

import (
	"image"
	"image/color"
	"image/draw"

	"warabiz/api/pkg/utils/encryption"
	"warabiz/api/pkg/utils/file"

	"github.com/fogleman/gg"
	"github.com/skip2/go-qrcode"
)

func GenerateBackground(cfg *Configuration) *image.RGBA {
	bgImg := image.NewRGBA(image.Rect(0, 0, cfg.BgConf.Width, cfg.BgConf.Height))
	draw.Draw(bgImg, bgImg.Bounds(), &image.Uniform{cfg.BgConf.BgColor}, image.ZP, draw.Src)
	return bgImg
}

func GenerateQRWithLogo(cfg *Configuration) (image.Image, error) {
	//* Create QR
	code, err := qrcode.New(cfg.Content.URL, qrcode.Highest)
	if err != nil {
		return nil, err
	}
	srcImage := code.Image(cfg.QRConf.Size)

	//* Download logo
	logoImg, err := file.DownloadImage(cfg.Content.Logo)
	if err != nil {
		return nil, err
	}
	bounds := logoImg.Bounds()
	if bounds.Dx() != 155 || bounds.Dy() != 155 {
		reImg, err := file.ResizeImage(logoImg, 155)
		if err != nil {
			return nil, err
		}
		logoImg = reImg
	}
	//* Convert logo into circular shape
	circularLogo := convertLogoToCircleShape(logoImg)
	//* Add logo into QR
	logoSize := float64(cfg.QRConf.Size) * float64(cfg.QRConf.Percent) / 100
	srcImage, err = addLogo(srcImage, circularLogo, int(logoSize))
	if err != nil {
		return nil, err
	}
	return srcImage, nil
}

func AddQrToBackground(bgImg image.Image, qrImg image.Image, yPosition int) image.Image {
	b := bgImg.Bounds()
	c := qrImg.Bounds()
	offset := image.Pt((b.Dx()-c.Dx())/2, yPosition)
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, bgImg, image.ZP, draw.Src)
	draw.Draw(image3, qrImg.Bounds().Add(offset), qrImg, image.ZP, draw.Over)
	return image3
}

func AddLabel(bgImg image.Image, label LabelConfiguration) (image.Image, error) {
	bounds := bgImg.Bounds()
	imgWidth := bounds.Dx()
	imgHeight := bounds.Dy()
	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImg, 0, 0)
	if err := dc.LoadFontFace(label.Font, label.FontSize); err != nil {
		return nil, err
	}
	x := float64(imgWidth / 2)
	y := float64((imgHeight / 2) - label.YPosition)
	maxWidth := float64(imgWidth) - 60.0
	dc.SetColor(label.Color)
	dc.DrawStringWrapped(label.Text, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)
	return dc.Image(), nil
}

func GenerateUniqueFromContent(content *Content) (string, error) {
	newUUID, err := encryption.GenerateUnique(content.URL + content.Logo + content.BrandName + content.SpgSubmissionID + content.SpgName)
	if err != nil {
		return "", err
	}
	return newUUID, nil
}

func convertLogoToCircleShape(logoImg image.Image) *image.RGBA {
	im := image.NewRGBA(image.Rect(0, 0, 200, 200))
	dc := gg.NewContextForRGBA(im)
	dc.SetColor(color.White)
	dc.DrawCircle(100, 100, 95)
	dc.Fill()
	b := logoImg.Bounds()
	x := 100 - b.Dx()/2
	y := 100 - b.Dy()/2
	draw.Draw(im, b.Add(image.Pt(x, y)), logoImg, image.ZP, draw.Over)
	return im
}

func addLogo(srcImage image.Image, logo image.Image, size int) (image.Image, error) {
	logoImage, err := file.ResizeImage(logo, uint(size))
	if err != nil {
		return nil, err
	}
	offset := image.Pt((srcImage.Bounds().Dx()-logoImage.Bounds().Dx())/2, (srcImage.Bounds().Dy()-logoImage.Bounds().Dy())/2)
	b := srcImage.Bounds()
	m := image.NewNRGBA(b)
	draw.Draw(m, b, srcImage, image.ZP, draw.Src)
	draw.Draw(m, logoImage.Bounds().Add(offset), logoImage, image.ZP, draw.Over)
	return m, nil
}
