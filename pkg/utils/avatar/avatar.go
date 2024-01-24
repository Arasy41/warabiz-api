package avatar

import (
	"image"
	"image/color"

	"warabiz/api/pkg/utils/avatar/src"
	"warabiz/api/pkg/utils/getter"
)

// func GetAvatar(cfg *config.Config, username string) (string, error) {

// 	initial := strings.ToUpper(src.GetInitialName(username))

// 	//* Get path
// 	path := fmt.Sprintf("%s/%s/AvatarImage/Initial/%s/%s", cfg.Connection.CloudStorage.GoogleStorage.AppName, cfg.Server.Env, initial, src.AvaVersion+".jpeg")
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
// 		//* Create AvatarImage
// 		avatarImg, err := createAvatar(initial)
// 		if err != nil {
// 			return "", err
// 		}
// 		//* Upload AvatarImage
// 		err = upload(cfg, client, avatarImg, path)
// 	}
// 	return fmt.Sprintf("%s/%s/%s", cfg.Connection.CloudStorage.GoogleStorage.GoogleCloudStorageURL, cfg.Connection.CloudStorage.GoogleStorage.GoogleCloudStorageBucket, path), nil
// }

func CreateAvatar(initial string) (image.Image, error) {
	c := src.PickRandomColor()
	bgImg := src.GenerateBackground(c)

	currentDir, _ := getter.GetCurrentDir()

	ava, err := src.AddLabel(bgImg, src.LabelConfiguration{
		Text:      initial,
		Font:      currentDir + "/pkg/fonts/Poppins-SemiBold.ttf",
		FontSize:  72,
		YPosition: 0,
		Color:     color.White,
	})
	if err != nil {
		return nil, err
	}
	return ava, nil
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
