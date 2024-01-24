package general

const (
	RedisConnectionError = "failed to get redis client"
	RedisNilError        = "Unable to GET data. error: redis: nil"

	GcpConnectionError = "failed to get gcp client"

	DBTimeLayout       string = "2006-01-02 15:04:05"
	ResponseTimeLayout string = "2006-01-02T15:04:05-0700"
	RequestTimeLayout  string = "2006-01-02T15:04:05"

	LayoutDateOnly string = "2006-01-02"
	LayoutTimeOnly string = "15:04:00"
)
