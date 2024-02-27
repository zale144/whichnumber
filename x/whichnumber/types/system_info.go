package types

func DefaultSystemInfo() SystemInfo {
	return SystemInfo{
		NextId:     DefaultId,
		FifoHeadId: NoFifoId,
		FifoTailId: NoFifoId,
	}
}
