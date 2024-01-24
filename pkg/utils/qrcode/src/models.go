package src

import "image/color"

type Configuration struct {
	Content       Content
	BgConf        bgConfiguration
	LayoutConf    layoutConfiguration
	QRConf        qrConfiguration
	QRResizeConf  qrResizeConfiguration
	AddQRToBgConf addQrToBgConfiguration
}

type Content struct {
	SpgQrCodeId     int64
	Logo            string
	URL             string
	BrandName       string
	SpgName         string
	SpgSubmissionID string
}

type bgConfiguration struct {
	Width   int
	Height  int
	BgColor color.Color
}

type layoutConfiguration struct {
	BrandNameConf brandNameConfiguration
	DashLineConf  dashLineConfiguration
	IdentityConf  identityConfiguration
	LinkConf      linkConfiguration
}

type brandNameConfiguration struct {
	Line1            string
	Line2            string
	TotalLine        int
	Font             string
	FontSize         float64
	Line1YPositionV1 int
	Line1YPositionV2 int
	Line2YPosition   int
	Color            color.Color
}

type dashLineConfiguration struct {
	Text      string
	Font      string
	FontSize  float64
	YPosition int
	Color     color.Color
}

type identityConfiguration struct {
	Line1            string
	Line2            string
	TotalLine        int
	Font             string
	FontSize         float64
	Line1YPositionV1 int
	Line1YPositionV2 int
	Line2YPosition   int
	Color            color.Color
}

type linkConfiguration struct {
	Line1            string
	Line2            string
	TotalLine        int
	Font             string
	FontSize         float64
	Line1YPositionV1 int
	Line1YPositionV2 int
	Line2YPosition   int
	Color            color.Color
}

type qrConfiguration struct {
	Size    int
	Percent int
}

type qrResizeConfiguration struct {
	Size int
}

type addQrToBgConfiguration struct {
	YPosition int
}

type LabelConfiguration struct {
	Text      string
	Font      string
	FontSize  float64
	YPosition int
	Color     color.Color
}
