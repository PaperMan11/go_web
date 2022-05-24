package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config") // 指定配置文件名称（不需要要带后缀）
	viper.SetConfigType("yaml")   // 指定配置文件类型
	viper.AddConfigPath(".")      // 指定查找配置文件的路径
	err = viper.ReadInConfig()    // 读取配置信息
	if err != nil {
		fmt.Println("viper.ReadInConfig failed err:", err)
		return
	}
	viper.WatchConfig()                            // 热加载
	viper.OnConfigChange(func(in fsnotify.Event) { // hockfunc
		fmt.Println("配置文件修改了")
	})
	return
}
