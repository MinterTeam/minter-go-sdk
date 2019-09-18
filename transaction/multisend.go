package transaction

type MultisendData struct {
	List []MultisendDataItem
}

type MultisendDataItem SendData
