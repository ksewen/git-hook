package main

import (
	logging "github.com/op/go-logging"
	"io"
	"os"
)

/**
 *日志级别
 */
const (
	critical int = iota
	err_or
	warning
	notice
	info
	debug
)

var (
	// 日志管理器
	mlog *Logger
	/**
	 * 日志输出格式
	 */
	logFormat = []string{
		`%{shortfunc} ▶ %{level:.4s} %{message}`,
		`%{time:15:04:05.00} %{shortfunc} ▶ %{level:.4s} %{id:03x} %{message}`,
		`%{color}%{time:15:04:05.00} %{shortfunc} %{shortfile} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	}

	/**
	 * 日志级别与 string类型映射
	 */
	LogLevelMap = map[string]int{
		"CRITICAL": critical,
		"ERROR":    err_or,
		"WARNING":  warning,
		"NOTICE":   notice,
		"INFO":     info,
		"DEBUG":    debug,
	}
)

type Logger struct {
	log      *logging.Logger
	level    string
	filePath string
	ModeName string
	format   int
}

/**
 * 初始化日志
 * @param logLevel The arguments could be INFO, DEGUE, ERROR
 */
func Init(logLevel string, filePath string, modeName string) {
	mlog = newLog(logLevel, filePath, modeName)
	return
}

func newLog(level string, filePath string, modeName string) *Logger {
	log := new(Logger)
	log.level = level
	log.filePath = filePath
	log.ModeName = modeName
	log.format = 2
	log.log = logging.MustGetLogger(log.ModeName)
	log.AddLogBackend()
	return log
}

/**
 *添加日志输出终端，可以是文件，控制台，还有网络输出等。
 */
func (l *Logger) AddLogBackend() {
	l.log.ExtraCalldepth = 2
	// 打开文件输出终端
	backend1 := l.getFileBackend()
	backend2 := l.getStdOutBackend()
	logging.SetBackend(backend1, backend2)
	return
}

func (l *Logger) getFileBackend() logging.LeveledBackend {
	// 打开一个文件
	file, err := os.OpenFile(l.filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	backend := l.getLogBackend(file, LogLevelMap[l.level])
	logging.SetBackend(backend)
	return backend
}

func (l *Logger) getStdOutBackend() logging.LeveledBackend {
	bked := l.getLogBackend(os.Stderr, LogLevelMap[l.level])
	return bked
}

/**
 * 获取终端
 */
func (l *Logger) getLogBackend(out io.Writer, level int) logging.LeveledBackend {
	backend := logging.NewLogBackend(out, "", 1)
	format := logging.MustStringFormatter(logFormat[2])
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(logging.Level(level), "")
	return backendLeveled
}

func (l *Logger) logI(infmt string, args ...interface{}) {
	l.log.Infof(infmt, args...)
	return
}
func (l *Logger) logE(infmt string, args ...interface{}) {
	l.log.Errorf(infmt, args...)
	return
}
func (l *Logger) logW(infmt string, args ...interface{}) {
	l.log.Warningf(infmt, args...)
	return
}
func (l *Logger) logD(infmt string, args ...interface{}) {
	l.log.Debugf(infmt, args...)
	return
}

func LogI(fmtstr string, args ...interface{}) {
	mlog.logI(fmtstr, args...)
	return
}

func LogW(fmtstr string, args ...interface{}) {
	mlog.logW(fmtstr, args...)
	return
}

func LogE(fmtstr string, args ...interface{}) {
	mlog.logE(fmtstr, args...)
	return
}

func LogD(fmtstr string, args ...interface{}) {
	mlog.logD(fmtstr, args...)
	return
}
