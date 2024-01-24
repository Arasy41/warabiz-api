package qrcode

import (
	"image"

	"warabiz/api/pkg/utils/file"
	"warabiz/api/pkg/utils/qrcode/src"
)

// func GetQRCode(cfg *config.Config, content *src.Content) (string, error) {
// 	//* Get unique file name
// 	unique, err := src.GenerateUniqueFromContent(content)
// 	if err != nil {
// 		return "", err
// 	}
// 	//* Get path
// 	cat1 := getter.GetThousandsCategory(float64(content.SpgQrCodeId))
// 	cat2 := getter.GetHundredsCategory(float64(content.SpgQrCodeId))
// 	path := fmt.Sprintf("%s/%s/QRCodeImage/%s/%s/%v/%s", cfg.Connection.CloudStorage.GoogleStorage.AppName, cfg.Server.Env, cat1, cat2, content.SpgQrCodeId, unique+src.QrVersion+".jpeg")
// 	//* Define new gcp client
// 	client, err := gcp.NewGCPClient(&cfg.Connection.CloudStorage.GoogleStorage)
// 	if err != nil {
// 		return "", err
// 	}
// 	//* Check if file is already exist
// 	isExist, err := gcp.CheckFile(client, path)
// 	if err != nil {
// 		return "", err
// 	}
// 	if !isExist {
// 		//* Call configuration
// 		config := src.SetConfiguration(content)
// 		//* Create QRImage
// 		qrImg, err := createQRImage(config)
// 		if err != nil {
// 			return "", err
// 		}
// 		//* Upload QRImage
// 		err = upload(cfg, client, qrImg, path)
// 	}
// 	return fmt.Sprintf("%s/%s/%s", cfg.Connection.CloudStorage.GoogleStorage.GoogleCloudStorageURL, cfg.Connection.CloudStorage.GoogleStorage.GoogleCloudStorageBucket, path), nil
// }

func CreateQRImage(cfg *src.Configuration) (image.Image, error) {
	//* Create background
	bgImg := src.GenerateBackground(cfg)
	//* Create QR
	qrImg, err := src.GenerateQRWithLogo(cfg)
	if err != nil {
		return nil, err
	}
	//* Resize QR
	newQrImg, err := file.ResizeImage(qrImg, uint(cfg.QRResizeConf.Size))
	if err != nil {
		return nil, err
	}
	//* Add QR into background
	newImgWithQr := src.AddQrToBackground(bgImg, newQrImg, cfg.AddQRToBgConf.YPosition)
	//* Add BrandName Label
	var newImg image.Image
	if cfg.LayoutConf.BrandNameConf.TotalLine == 1 {
		brandName := src.LabelConfiguration{
			Text:      cfg.LayoutConf.BrandNameConf.Line1,
			Font:      cfg.LayoutConf.BrandNameConf.Font,
			FontSize:  cfg.LayoutConf.BrandNameConf.FontSize,
			YPosition: cfg.LayoutConf.BrandNameConf.Line1YPositionV1,
			Color:     cfg.LayoutConf.BrandNameConf.Color,
		}
		newImg, err = src.AddLabel(newImgWithQr, brandName)
		if err != nil {
			return nil, err
		}
	} else {
		brandName := src.LabelConfiguration{
			Text:      cfg.LayoutConf.BrandNameConf.Line1,
			Font:      cfg.LayoutConf.BrandNameConf.Font,
			FontSize:  cfg.LayoutConf.BrandNameConf.FontSize,
			YPosition: cfg.LayoutConf.BrandNameConf.Line1YPositionV2,
			Color:     cfg.LayoutConf.BrandNameConf.Color,
		}
		newImg, err = src.AddLabel(newImgWithQr, brandName)
		if err != nil {
			return nil, err
		}
		brandName2 := src.LabelConfiguration{
			Text:      cfg.LayoutConf.BrandNameConf.Line2,
			Font:      cfg.LayoutConf.BrandNameConf.Font,
			FontSize:  cfg.LayoutConf.BrandNameConf.FontSize,
			YPosition: cfg.LayoutConf.BrandNameConf.Line2YPosition,
			Color:     cfg.LayoutConf.BrandNameConf.Color,
		}
		newImg, err = src.AddLabel(newImg, brandName2)
		if err != nil {
			return nil, err
		}
	}
	//* Add dash line
	dashLine := src.LabelConfiguration{
		Text:      cfg.LayoutConf.DashLineConf.Text,
		Font:      cfg.LayoutConf.DashLineConf.Font,
		FontSize:  cfg.LayoutConf.DashLineConf.FontSize,
		YPosition: cfg.LayoutConf.DashLineConf.YPosition,
		Color:     cfg.LayoutConf.DashLineConf.Color,
	}
	newImg, err = src.AddLabel(newImg, dashLine)
	if err != nil {
		return nil, err
	}
	//* Add Identity Label
	if cfg.LayoutConf.IdentityConf.TotalLine == 1 {
		identity := src.LabelConfiguration{
			Text:      cfg.LayoutConf.IdentityConf.Line1,
			Font:      cfg.LayoutConf.IdentityConf.Font,
			FontSize:  cfg.LayoutConf.IdentityConf.FontSize,
			YPosition: cfg.LayoutConf.IdentityConf.Line1YPositionV1,
			Color:     cfg.LayoutConf.IdentityConf.Color,
		}
		newImg, err = src.AddLabel(newImg, identity)
		if err != nil {
			return nil, err
		}
	} else {
		identity := src.LabelConfiguration{
			Text:      cfg.LayoutConf.IdentityConf.Line1,
			Font:      cfg.LayoutConf.IdentityConf.Font,
			FontSize:  cfg.LayoutConf.IdentityConf.FontSize,
			YPosition: cfg.LayoutConf.IdentityConf.Line1YPositionV2,
			Color:     cfg.LayoutConf.IdentityConf.Color,
		}
		newImg, err = src.AddLabel(newImg, identity)
		if err != nil {
			return nil, err
		}
		identity2 := src.LabelConfiguration{
			Text:      cfg.LayoutConf.IdentityConf.Line2,
			Font:      cfg.LayoutConf.IdentityConf.Font,
			FontSize:  cfg.LayoutConf.IdentityConf.FontSize,
			YPosition: cfg.LayoutConf.IdentityConf.Line2YPosition,
			Color:     cfg.LayoutConf.IdentityConf.Color,
		}
		newImg, err = src.AddLabel(newImg, identity2)
		if err != nil {
			return nil, err
		}
	}
	//* Add Link Label
	if cfg.LayoutConf.LinkConf.TotalLine == 1 {
		link := src.LabelConfiguration{
			Text:      cfg.LayoutConf.LinkConf.Line1,
			Font:      cfg.LayoutConf.LinkConf.Font,
			FontSize:  cfg.LayoutConf.LinkConf.FontSize,
			YPosition: cfg.LayoutConf.LinkConf.Line1YPositionV1,
			Color:     cfg.LayoutConf.LinkConf.Color,
		}
		newImg, err = src.AddLabel(newImg, link)
		if err != nil {
			return nil, err
		}
	} else {
		link := src.LabelConfiguration{
			Text:      cfg.LayoutConf.LinkConf.Line1,
			Font:      cfg.LayoutConf.LinkConf.Font,
			FontSize:  cfg.LayoutConf.LinkConf.FontSize,
			YPosition: cfg.LayoutConf.LinkConf.Line1YPositionV2,
			Color:     cfg.LayoutConf.LinkConf.Color,
		}
		newImg, err = src.AddLabel(newImg, link)
		if err != nil {
			return nil, err
		}
		link2 := src.LabelConfiguration{
			Text:      cfg.LayoutConf.LinkConf.Line2,
			Font:      cfg.LayoutConf.LinkConf.Font,
			FontSize:  cfg.LayoutConf.LinkConf.FontSize,
			YPosition: cfg.LayoutConf.LinkConf.Line2YPosition,
			Color:     cfg.LayoutConf.LinkConf.Color,
		}
		newImg, err = src.AddLabel(newImg, link2)
		if err != nil {
			return nil, err
		}
	}
	return newImg, nil
}

// func upload(cfg *config.Config, client *gcp.ClientUploader, img image.Image, path string) error {
// 	//* Create context
// 	ctx := context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
// 	defer cancel()
// 	//* Create Buff
// 	buf := new(bytes.Buffer)
// 	if err := jpeg.Encode(buf, img, nil); err != nil {
// 		return err
// 	}
// 	reader := bytes.NewReader(buf.Bytes())
// 	//* Upload
// 	wc := client.Cl.Bucket(client.BucketName).Object(path).NewWriter(ctx)
// 	if _, err := io.Copy(wc, reader); err != nil {
// 		return fmt.Errorf("io.Copy: %v", err)
// 	}
// 	if err := wc.Close(); err != nil {
// 		return fmt.Errorf("Writer.Close: %v", err)
// 	}
// 	return nil
// }
