package process

import (
	"context"

	"github.com/shirou/gopsutil/v3/internal/common"
)

type FreeStat struct {
	Path string `json:"path"`
	Rss  uint64 `json:"rss"`
	Size uint64 `json:"size"`
	Pss  uint64 `json:"pss"`
}

func (f *FreeStat) ProcessesWithContext(ctx context.Context) ([]*Process, error) {
	return nil, common.ErrNotImplementedError
}

func (f *FreeStat) PidExistsWithContext(ctx context.Context, pid int32) (bool, error) {
	return false, common.ErrNotImplementedError
}
