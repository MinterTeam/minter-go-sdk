package transaction

import (
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/rlp"
)

type DeepLink struct {
	Type    Type
	Data    []byte
	Payload []byte

	Nonce    *uint     // optional
	GasPrice *uint     // optional
	GasCoin  *[10]byte // optional
}

func (d *DeepLink) Encode() (string, error) {
	src, err := rlp.EncodeToBytes(d)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(src), err
}

func (d *DeepLink) setType(t Type) *DeepLink {
	d.Type = t
	return d
}

func (d *DeepLink) SetPayload(payload []byte) *DeepLink {
	d.Payload = payload
	return d
}

func (d *DeepLink) NewDeepLink(data DataInterface) (*DeepLink, error) {
	bytes, err := data.encode()
	if err != nil {
		return d, err
	}
	d.Data = bytes

	switch data.(type) {
	case *SendData:
		return d.setType(TypeSend), nil
	case *SellCoinData:
		return d.setType(TypeSellCoin), nil
	case *SellAllCoinData:
		return d.setType(TypeSellAllCoin), nil
	case *BuyCoinData:
		return d.setType(TypeBuyCoin), nil
	case *CreateCoinData:
		return d.setType(TypeCreateCoin), nil
	case *DeclareCandidacyData:
		return d.setType(TypeDeclareCandidacy), nil
	case *DelegateData:
		return d.setType(TypeDelegate), nil
	case *UnbondData:
		return d.setType(TypeUnbond), nil
	case *RedeemCheckData:
		return d.setType(TypeRedeemCheck), nil
	case *SetCandidateOnData:
		return d.setType(TypeSetCandidateOnline), nil
	case *SetCandidateOffData:
		return d.setType(TypeSetCandidateOffline), nil
	// case *CreateMultisigData:
	//	return transaction.setType(TypeCreateMultisig), nil
	case *MultisendData:
		return d.setType(TypeMultisend), nil
	case *EditCandidateData:
		return d.setType(TypeEditCandidate), nil

	default:
		return nil, errors.New("unknown transaction type")
	}
}
