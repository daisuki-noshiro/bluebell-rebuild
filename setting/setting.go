package setting

import "github.com/spf13/viper"

// Conf 保存读取到的全局配置
var Conf = new(AppConfig)

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

func Init(filePath string) error {
	// 指定要读取的配置文件
	viper.SetConfigFile(filePath)

	//把配置文件读取进Viper
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	// 把配置内容解析到 Conf 指向的 AppConfig 中
	if err := viper.Unmarshal(Conf); err != nil {
		return err
	}
	return nil
}
