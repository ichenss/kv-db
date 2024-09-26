package utils

import (
	"golang.org/x/sys/windows"
	"io/fs"
	"path/filepath"
	"syscall"
)

// DirSize 获取一个目录大小
func DirSize(dirPath string) (int64, error) {
	var size int64
	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}

func AvailableDiskSize() (uint64, error) {
	wd, err := syscall.Getwd()
	if err != nil {
		return 0, err
	}
	// 以下为 Linux 环境中写法
	// var stat syscall.Statfs_t
	// if err = syscall.Statfs_t(wd, &stat); err != nil {
	// 	return 0, nil
	// }
	// return stat.Bavail * uint64(stat.Bsize), nil

	// 在 Windows 环境中获取磁盘可用空间
	var freeBytes uint64
	var totalBytes uint64
	var totalFreeBytes uint64

	// 获取驱动器的状态信息
	err = windows.GetDiskFreeSpaceEx(windows.StringToUTF16Ptr(wd), &freeBytes, &totalBytes, &totalFreeBytes)
	if err != nil {
		return 0, err
	}

	return freeBytes, nil
}
