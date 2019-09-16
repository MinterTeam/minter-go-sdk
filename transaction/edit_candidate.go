package transaction

import (
	"github.com/ethereum/go-ethereum/rlp"
)

type EditCandidateData struct {
	PubKey        []byte
	RewardAddress [20]byte
	OwnerAddress  [20]byte
}

func NewEditCandidateData() *EditCandidateData {
	return &EditCandidateData{}
}

func (d *EditCandidateData) SetPubKey(key string) *EditCandidateData {
	d.PubKey = []byte(key)
	return d
}

func (d *EditCandidateData) MustSetRewardAddress(address string) *EditCandidateData {
	_, err := d.SetRewardAddress(address)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *EditCandidateData) SetRewardAddress(address string) (*EditCandidateData, error) {
	bytes, err := AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.RewardAddress[:], bytes)
	return d, nil
}

func (d *EditCandidateData) MustSetOwnerAddress(address string) *EditCandidateData {
	_, err := d.SetOwnerAddress(address)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *EditCandidateData) SetOwnerAddress(address string) (*EditCandidateData, error) {
	bytes, err := AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.OwnerAddress[:], bytes)
	return d, nil
}

func (d *EditCandidateData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
