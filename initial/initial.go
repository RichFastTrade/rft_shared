package initial

func Init() error {
	// 初始化日志
	initLog()
	// 初始化配置
	if err := initConfig(); err != nil {
		return err
	}
	// 初始化队列
	if err := initRabbitMQ(); err != nil {
		return err
	}
	return nil
}
