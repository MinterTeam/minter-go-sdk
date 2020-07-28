package transaction

import (
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/rlp"
	"net/url"
)

type DeepLink struct {
	Type    Type
	Data    []byte
	Payload []byte

	Nonce    *uint   `rlp:"nil"` // optional
	GasPrice *uint   `rlp:"nil"` // optional
	GasCoin  *CoinID `rlp:"nil"` // optional
}

// Returns deep link.
func (d *DeepLink) CreateLink(pass string) (string, error) {
	tx, err := d.Encode()
	if err != nil {
		return "", err
	}

	rawQuery := ""
	if pass != "" {
		rawQuery = fmt.Sprintf("p=%s", base64.RawStdEncoding.EncodeToString([]byte(pass)))
	}

	u := &url.URL{
		Scheme:   "https",
		Host:     "bip.to",
		Path:     fmt.Sprintf("/tx/%s", tx),
		RawQuery: rawQuery,
	}
	return u.String(), nil
}

func (d *DeepLink) Encode() (string, error) {
	src, err := rlp.EncodeToBytes(d)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(src), nil
}

func (d *DeepLink) setType(t Type) *DeepLink {
	d.Type = t
	return d
}

func (d *DeepLink) SetPayload(payload []byte) *DeepLink {
	d.Payload = payload
	return d
}

func (d *DeepLink) SetNonce(nonce uint) *DeepLink {
	d.Nonce = &nonce
	return d
}

func (d *DeepLink) SetGasPrice(gasPrice uint) *DeepLink {
	d.GasPrice = &gasPrice
	return d
}

func (d *DeepLink) SetGasCoin(id CoinID) *DeepLink {
	d.GasCoin = &id
	return d
}

func NewDeepLink(data Data) (*DeepLink, error) {
	d := new(DeepLink)

	bytes, err := data.encode()
	if err != nil {
		return d, err
	}
	d.Data = bytes

	return d.setType(data.Type()), nil
}
