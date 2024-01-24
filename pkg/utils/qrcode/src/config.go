package src

import (
	"fmt"
	"image/color"
	"strings"

	"warabiz/api/pkg/utils/getter"
)

func SetConfiguration(content *Content) *Configuration {

	brandNameContent := getBrandNameConfiguration(content.BrandName)
	identityContent := getIdentityConfiguration(fmt.Sprintf("%s | %s", content.SpgSubmissionID, content.SpgName))
	linkContent := getLinkConfiguration(content.URL)

	return &Configuration{
		Content: Content{
			Logo:            content.Logo,
			URL:             content.URL,
			BrandName:       content.BrandName,
			SpgName:         content.SpgName,
			SpgSubmissionID: content.SpgSubmissionID,
		},
		BgConf: bgConfiguration{
			Width:   420,
			Height:  594,
			BgColor: color.White,
		},
		QRConf: qrConfiguration{
			Size:    2400,
			Percent: 30,
		},
		QRResizeConf: qrResizeConfiguration{
			Size: 375,
		},
		AddQRToBgConf: addQrToBgConfiguration{
			YPosition: 75,
		},
		LayoutConf: layoutConfiguration{
			BrandNameConf: brandNameConfiguration{
				Line1:            brandNameContent.Line1,
				Line2:            brandNameContent.Line2,
				TotalLine:        brandNameContent.TotalLine,
				Font:             "fonts/Poppins-SemiBold.ttf",
				FontSize:         32,
				Line1YPositionV1: 245,
				Line1YPositionV2: 265,
				Line2YPosition:   225,
				Color:            color.Black,
			},
			DashLineConf: dashLineConfiguration{
				Text:      "---------------------------------------------------",
				Font:      "fonts/Poppins-Regular.ttf",
				FontSize:  14,
				YPosition: -175,
				Color:     color.Black,
			},
			IdentityConf: identityConfiguration{
				Line1:            identityContent.Line1,
				Line2:            identityContent.Line2,
				TotalLine:        identityContent.TotalLine,
				Font:             "fonts/Poppins-SemiBold.ttf",
				FontSize:         20,
				Line1YPositionV1: -208,
				Line1YPositionV2: -195,
				Line2YPosition:   -218,
				Color:            color.Black,
			},
			LinkConf: linkConfiguration{
				Line1:            linkContent.Line1,
				Line2:            linkContent.Line2,
				TotalLine:        linkContent.TotalLine,
				Font:             "fonts/Poppins-Regular.ttf",
				FontSize:         14,
				Line1YPositionV1: -245,
				Line1YPositionV2: -245,
				Line2YPosition:   -265,
				Color:            color.RGBA{195, 195, 195, 255},
			},
		},
	}
}

func getBrandNameConfiguration(brandName string) brandNameConfiguration {
	var line1 []string
	var line2 []string
	brandNameParts := strings.Split(brandName, " ")
	if len(brandNameParts) != 1 {
		for _, attr := range brandNameParts {
			brandName += " " + attr
		}
	} else {
		brandName = brandNameParts[0]
	}
	var counter1, counter2 int
	for i, attr := range brandNameParts {
		counter1 += getter.GetStringLength(attr)
		if counter1+i-1 <= maxBrandName {
			line1 = append(line1, attr)
		} else {
			counter2 += getter.GetStringLength(attr)
			if counter2+i-1 <= maxBrandNameWithDot {
				line2 = append(line2, attr)
			} else {
				counter2 += getter.GetStringLength(attr)
				if counter2+i+1 == maxBrandName {
					line2 = append(line2, attr)
				} else {
					line2[len(line2)-1] += "..."
					break
				}
			}
		}
	}
	var joinLine1, joinLine2 string
	if len(line1) == 1 {
		joinLine1 = line1[0]
	} else {
		for _, str1 := range line1 {
			joinLine1 = joinLine1 + str1 + " "
		}
	}
	if len(line2) == 0 {
		return brandNameConfiguration{
			Line1:     joinLine1,
			TotalLine: 1,
		}
	} else {
		if len(line2) == 1 {
			joinLine2 = line2[0] + " "
		} else {
			for _, str2 := range line2 {
				joinLine2 = joinLine2 + str2 + " "
			}
		}
		return brandNameConfiguration{
			Line1:     joinLine1[0 : getter.GetStringLength(joinLine1)-1],
			Line2:     joinLine2[0 : getter.GetStringLength(joinLine2)-1],
			TotalLine: 2,
		}
	}
}

func getIdentityConfiguration(identity string) identityConfiguration {
	var line1 []string
	var line2 []string
	identityParts := strings.Split(identity, " ")
	if len(identityParts) != 1 {
		for _, attr := range identityParts {
			identity += " " + attr
		}
	} else {
		identity = identityParts[0]
	}
	var counter1, counter2 int
	for i, attr := range identityParts {
		counter1 += getter.GetStringLength(attr)
		if counter1+i-1 <= maxIdentity {
			line1 = append(line1, attr)
		} else {
			counter2 += getter.GetStringLength(attr)
			if counter2+i-1 <= maxIdentityWithDot {
				line2 = append(line2, attr)
			} else {
				counter2 += getter.GetStringLength(attr)
				if counter2+i+1 == maxIdentity {
					line2 = append(line2, attr)
				} else {
					line2[len(line2)-1] += "..."
					break
				}
			}
		}
	}
	var joinLine1, joinLine2 string
	if len(line1) == 1 {
		joinLine1 = line1[0]
	} else {
		for _, str1 := range line1 {
			joinLine1 = joinLine1 + str1 + " "
		}
	}
	if len(line2) == 0 {
		return identityConfiguration{
			Line1:     joinLine1,
			TotalLine: 1,
		}
	} else {
		if len(line2) == 1 {
			joinLine2 = line2[0] + " "
		} else {
			for _, str2 := range line2 {
				joinLine2 = joinLine2 + str2 + " "
			}
		}
		return identityConfiguration{
			Line1:     joinLine1[0 : getter.GetStringLength(joinLine1)-1],
			Line2:     joinLine2[0 : getter.GetStringLength(joinLine2)-1],
			TotalLine: 2,
		}
	}
}

func getLinkConfiguration(link string) linkConfiguration {
	var line1 string
	var line2 string
	if getter.GetStringLength(link) <= maxLink {
		line1 = link
	}
	if getter.GetStringLength(link) > maxLink {
		line1 = link[0:maxLink]
		newLink := link[maxLink+1 : getter.GetStringLength(link)]
		if getter.GetStringLength(newLink) > maxIdentityWithDot {
			line2 = newLink[0:maxIdentityWithDot] + "..."
		} else {
			line2 = newLink
		}
	}
	if line2 == "" {
		return linkConfiguration{
			Line1:     line1,
			TotalLine: 1,
		}
	} else {
		return linkConfiguration{
			Line1:     line1,
			Line2:     line2,
			TotalLine: 2,
		}
	}
}
