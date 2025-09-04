package global

var (
	JWTSetting = &JWTSettingS{
		Secret: "linglong_secret", // JWT 密钥（实际建议用环境变量）
		Expire: time.Hour * 24,    // Token 有效期 24 小时
	}
)

type JWTSettingS struct {
	Secret string
	Expire time.Duration
}
