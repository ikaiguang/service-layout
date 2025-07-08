package global

const (
	KeyPrefix = "test_service_"
)

func Key(k string) string {
	return KeyPrefix + k
}
