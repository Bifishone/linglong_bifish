package global

import "linglong/pkg/setting"
import "time"

var (
	ServerSetting *ServerSettingS
	JWTSetting    *JWTSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	MasscanSetting *setting.MasscanSettingS
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type JWTSettingS struct {
	Secret string
	Expire time.Duration
}
