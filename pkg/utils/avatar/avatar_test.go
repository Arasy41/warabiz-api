package avatar

import (
	"fmt"
	"image/color"
	"log"
	"strings"
	"testing"

	"warabiz/api/pkg/utils/avatar/src"
	"warabiz/api/pkg/utils/file"
)

func TestAvatar(t *testing.T) {

	var err error

	username := "putra.rama"
	initial := strings.ToUpper(src.GetInitialName(username))
	c := src.PickRandomColor()

	bgImg := src.GenerateBackground(c)
	ava, err := src.AddLabel(bgImg, src.LabelConfiguration{
		Text:      initial,
		Font:      "../../fonts/Poppins-SemiBold.ttf",
		FontSize:  72,
		YPosition: 0,
		Color:     color.White,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = file.CreateJPEGFile(ava, fmt.Sprintf("./example/%s.jpeg", initial))
	if err != nil {
		log.Fatal(err)
	}
}
