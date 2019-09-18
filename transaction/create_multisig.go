package transaction

type CreateMultisigData struct {
	Threshold uint
	Weights   []uint
	Addresses [][20]byte
}
