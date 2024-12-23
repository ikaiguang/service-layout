package po

const (
	KeyPrefix = "t_"
)

func Key(k string) string {
	return KeyPrefix + k
}
