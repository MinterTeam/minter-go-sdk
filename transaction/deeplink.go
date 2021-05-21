package transaction

import (
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/ethereum/go-ethereum/rlp"
)

type DeepLink struct {
	Type     Type    // type of transaction
	Data     []byte  // data of transaction (depends on transaction type)
	Payload  []byte  // optional, arbitrary user-defined bytes
	Nonce    *uint32 `rlp:"nilList"` // optional, used for prevent transaction reply
	GasPrice *uint32 `rlp:"nilList"` // optional, fee multiplier, should be equal or greater than current mempool min gas price
	GasCoin  *CoinID `rlp:"nilList"` // optional, ID of a coin to pay fee, right padded with zeros

	url *url.URL
}

// Returns url link.
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
		Scheme:   d.url.Scheme,
		Host:     d.url.Host,
		Path:     fmt.Sprintf("/tx/%s", tx),
		RawQuery: rawQuery,
	}
	return u.String(), nil
}

// Returns tx-like data. RLP-encoded structure in base64url format.
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

// Set arbitrary user-defined bytes
func (d *DeepLink) SetPayload(payload []byte) *DeepLink {
	d.Payload = payload
	return d
}

// Set nonce of transaction
func (d *DeepLink) SetNonce(nonce uint32) *DeepLink {
	d.Nonce = &nonce
	return d
}

// SetGasPrice sets fee multiplier.
func (d *DeepLink) SetGasPrice(gasPrice uint32) *DeepLink {
	d.GasPrice = &gasPrice
	return d
}

// SetGasCoin sets ID of a coin to pay fee
func (d *DeepLink) SetGasCoin(id uint64) *DeepLink {
	coinID := CoinID(id)
	d.GasCoin = &coinID
	return d
}

func (d *DeepLink) SetUrl(value string) (*DeepLink, error) {
	u, err := url.Parse(value)

	if err != nil {
		return d, err
	}

	d.url = u

	return d, nil
}

func (d *DeepLink) MustSetUrl(value string) *DeepLink {
	dl, err := d.SetUrl(value)

	if err != nil {
		panic(err)
	}

	return dl
}

func NewDeepLink(data Data) (*DeepLink, error) {
	d := new(DeepLink)

	bytes, err := data.Encode()
	if err != nil {
		return d, err
	}
	d.Data = bytes

	d.url = &url.URL{
		Scheme: "https",
		Host:   "bip.to",
	}

	return d.setType(data.Type()), nil
}
