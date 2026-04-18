package conf

// IsDev 开发环境
func IsDev() bool {
	return env == "dev"
}

// IsTest 测试环境
func IsTest() bool {
	return env == "sim"
}

// IsSim 仿真环境
func IsSim() bool {
	return env == "sim"
}

// IsPre 预览环境
func IsPre() bool {
	return env == "pre"
}

// IsProd 线上环境
func IsProd() bool {
	return env == "prod"
}
