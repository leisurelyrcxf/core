package complexscheduler

import (
	"testing"

	"github.com/docker/go-units"
	"github.com/stretchr/testify/assert"

	resourcetypes "github.com/projecteru2/core/resources/types"
	"github.com/projecteru2/core/types"
)

func TestCPUPriorPlan(t *testing.T) {
	// normal 分配
	scheduleInfos := resetscheduleInfos()
	_, resultCPUPlan, total, err := cpuPriorPlan(3.0, int64(units.MiB), scheduleInfos, -1, 100)
	assert.NoError(t, err)
	assert.Equal(t, len(resultCPUPlan), 1)
	assert.Equal(t, total, 1)
	// numa 分配
	scheduleInfos = resetscheduleInfos()
	_, resultCPUPlan, total, err = cpuPriorPlan(1.5, int64(units.MiB), scheduleInfos, -1, 100)
	assert.NoError(t, err)
	assert.Equal(t, len(resultCPUPlan), 1)
	assert.Equal(t, total, 2)
	r := resultCPUPlan["n1"]
	for _, p := range r {
		_, ok1 := p["1"]
		_, ok2 := p["2"]
		_, ok3 := p["3"]
		_, ok4 := p["4"]
		assert.True(t, (ok1 && ok3) || (ok2 && ok4))
	}
	// numa and normal 分配
	scheduleInfos = resetscheduleInfos()
	_, resultCPUPlan, total, err = cpuPriorPlan(1, int64(units.GiB), scheduleInfos, -1, 100)
	assert.NoError(t, err)
	assert.Equal(t, len(resultCPUPlan), 1)
	assert.Equal(t, total, 3)
}

func resetscheduleInfos() []resourcetypes.ScheduleInfo {
	return []resourcetypes.ScheduleInfo{
		{
			NodeMeta: types.NodeMeta{
				Name:   "n1",
				CPU:    types.CPUMap{"1": 100, "2": 100, "3": 100, "4": 100},
				MemCap: 3 * int64(units.GiB),
				NUMA: types.NUMA{
					"1": "node0",
					"2": "node1",
					"3": "node0",
					"4": "node1",
				},
				NUMAMemory: types.NUMAMemory{
					"node0": int64(units.GiB),
					"node1": int64(units.GiB),
				},
			},
		},
	}
}
