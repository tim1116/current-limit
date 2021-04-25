package common

import go_logger "github.com/phachon/go-logger"

// 日志处理器
func Logger() *go_logger.Logger {
	logger := go_logger.NewLogger()
	logger.Detach("console")
	// 命令行输出配置
	consoleConfig := &go_logger.ConsoleConfig{
		Color: true, // 命令行输出字符串是否显示颜色
	}
	// 添加 console 为 logger 的一个输出
	logger.Attach("console", go_logger.LOGGER_LEVEL_INFO, consoleConfig)
	return logger
}
