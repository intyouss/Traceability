package utils

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/v4/disk"

	"github.com/shirou/gopsutil/v4/mem"

	"github.com/shirou/gopsutil/v4/cpu"
)

// GetTotalCpuUsage 获取CPU总使用率
func GetTotalCpuUsage(ctx context.Context) (float64, error) {
	totalUsage, err := cpu.PercentWithContext(ctx, time.Second, false)
	if err != nil {
		return 0, err
	}
	return totalUsage[0], nil
}

// GetCpuUsage 获取每个CPU使用率
func GetCpuUsage(ctx context.Context) ([]float64, error) {
	return cpu.PercentWithContext(ctx, time.Second, true)
}

// GetCpuInfo 获取各CPU信息
func GetCpuInfo(ctx context.Context) ([]cpu.InfoStat, error) {
	return cpu.InfoWithContext(ctx)
}

// GetMemoryInfo 获取内存信息
func GetMemoryInfo(ctx context.Context) (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemoryWithContext(ctx)
}

// GetMemoryUsage 获取内存使用率
func GetMemoryUsage(ctx context.Context) (float64, error) {
	memory, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return 0, err
	}
	return memory.UsedPercent, nil
}

// GetDiskInfo 获取指定磁盘信息
func GetDiskInfo(ctx context.Context, path string) (*disk.UsageStat, error) {
	return disk.UsageWithContext(ctx, path)
}

// GetDiskUsage 获取指定磁盘使用率
func GetDiskUsage(ctx context.Context, path string) (float64, error) {
	d, err := disk.UsageWithContext(ctx, path)
	if err != nil {
		return 0, err
	}
	return d.UsedPercent, nil
}
