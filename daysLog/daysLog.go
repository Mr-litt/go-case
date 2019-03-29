package daysLog

import (
	"os"
	"time"
	"sync"
	"path/filepath"
)

type DaysLog struct {
	mutex        *sync.Mutex
	FilePath     string
	FileWriteMap map[string]*FileWrite
}

type FileWrite struct {
	date  string
	write *os.File
}

func NewDaysLog() *DaysLog {
	return &DaysLog{
		mutex:        new(sync.Mutex),
		FilePath:     "/data/log/daysLog/",
		FileWriteMap: make(map[string]*FileWrite),
	}
}

func (daysLog *DaysLog) SetFilePath(filePath string) {
	daysLog.FilePath = filePath
}

func (daysLog *DaysLog) Info(message string, moduleType string, channel string) error {
	return daysLog.Write(message, moduleType, channel, "INFO")
}

func (daysLog *DaysLog) Error(message string, moduleType string, channel string, ) error {
	return daysLog.Write(message, moduleType, channel, "ERROR")
}

func (daysLog *DaysLog) Write(message string, moduleType string, channel string, level string) error {

	// 锁
	daysLog.mutex.Lock()
	defer daysLog.mutex.Unlock()

	// 获取文件句柄
	fileWrite, err := daysLog.getFileWrite(moduleType, level)
	if err != nil {
		return err
	}

	// 获取消息格式
	msg := daysLog.getMsgFormat(channel, level, message)

	// 写入文件
	_, err = fileWrite.WriteString(msg)

	return err
}

func (daysLog *DaysLog) getFileWrite(moduleType string, level string) (*os.File, error) {

	// 获取文件名
	CurrentDate := time.Now().Format("2006-01-02")
	fileName := daysLog.FilePath + "/" + moduleType + "-" + level + "-" + CurrentDate + ".log"

	// 获取文件句柄
	fileWrite, ok := daysLog.FileWriteMap[fileName]
	if !ok {

		// 创建目录
		dirName := filepath.Dir(fileName)
		if err := os.MkdirAll(dirName, 0755); err != nil {
			return nil, err
		}

		// 创建文件
		fw, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}

		// 保存句柄
		daysLog.FileWriteMap[fileName] = &FileWrite{
			date:  CurrentDate,
			write: fw,
		}

		// 关闭过期句柄
		for k, v := range daysLog.FileWriteMap {
			if v.date != CurrentDate {
				v.write.Close()
				delete(daysLog.FileWriteMap, k)
			}
		}

		return fw, nil
	}

	return fileWrite.write, nil
}

func (daysLog *DaysLog) getMsgFormat(channel string, level string, message string) string {
	return time.Now().Format("2006-01-02 15:04:05") + " " + channel + " " + level + " " + message + "\n"
}
