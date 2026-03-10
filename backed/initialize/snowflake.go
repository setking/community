package initialize

import (
	"fmt"
	"math"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// 需传入当前的机器ID
func InitSnowflake(machineId uint16) {
	sonyMachineID = machineId
	t, _ := time.Parse("2006-01-02", "2020-01-01")
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
}
func genSonyFlakeId() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")
		return
	}
	id, err = sonyFlake.NextID()
	return
}

// GetID 返回生成的id值
func GetID() (id int64, err error) {
	uint64id, err := genSonyFlakeId()
	if err != nil {
		err = fmt.Errorf("generate id error:%w", err)
	}
	if uint64id > math.MaxInt64 {
		err = fmt.Errorf("id out of range, max value is %d", math.MaxInt64)
	}
	id = int64(uint64id)
	return
}
