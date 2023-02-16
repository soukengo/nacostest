package nacostest

var version = "0.1.0"

func Version() string {
	return version
}

func SetVersion(v string) string {
	if len(v) > 0 {
		version = v
	}
	return version
}
