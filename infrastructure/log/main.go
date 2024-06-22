package log

import (
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"os"
	"path/filepath"
)

func init() {
	// 设置日志级别为 Debug 及以上
	log.SetLevel(log.DebugLevel)

	// 创建日志文件
	logDir := filepath.Join(".", "logs")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Errorf("Failed to create log directory: %v", err)
		return
	}

	// 创建 error 日志文件
	errorLogFile, err := os.OpenFile(filepath.Join(logDir, "error.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("Failed to open error log file: %v", err)
		return
	}
	defer errorLogFile.Close()

	// 创建 info 日志文件
	infoLogFile, err := os.OpenFile(filepath.Join(logDir, "info.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("Failed to open info log file: %v", err)
		return
	}
	defer infoLogFile.Close()

	// 创建 warning 日志文件
	warnLogFile, err := os.OpenFile(filepath.Join(logDir, "warning.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("Failed to open warning log file: %v", err)
		return
	}
	defer warnLogFile.Close()

	// 设置 error 日志输出
	errorHook := &writer.Hook{
		Writer: errorLogFile,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
		},
	}
	log.AddHook(errorHook)

	// 设置 info 日志输出
	infoHook := &writer.Hook{
		Writer: infoLogFile,
		LogLevels: []log.Level{
			log.InfoLevel,
		},
	}
	log.AddHook(infoHook)

	// 设置 warning 日志输出
	warnHook := &writer.Hook{
		Writer: warnLogFile,
		LogLevels: []log.Level{
			log.WarnLevel,
		},
	}
	log.AddHook(warnHook)

	// 设置本地控制台日志输出
	log.SetOutput(os.Stdout)

	log.Debug("This is a debug message.")
	log.Info("This is an info message.")
	log.Warn("This is a warning message.")
	log.Error("This is an error message.")
}
