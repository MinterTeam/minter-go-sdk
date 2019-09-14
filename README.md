## About

This is a pure Go SDK for working with <b>Minter</b> blockchain

* [Minter Api]()
    - Methods:
	    - [GetBalance]()
	    - [GetNonce]()
	    - [Send]()
	
* [Minter SDK]()
	- [Sign transaction]()
		- [Send]()
		
* [Tests](#tests)

## Using MinterAPI

You can get all valid responses and full documentation at [Minter Node Api](https://minter-go-node.readthedocs.io/en/latest/api.html)

Create MinterAPI instance

```
import "github.com/MinterTeam/minter-go-sdk/api"

nodeUrl = "https://minter-node-1.testnet.minter.network:8841"
api := api.NewApi(node)
```

### GetAddress

Returns coins list, balance and transaction count (for nonce) of an address.

```
func (a *Api) GetAddress(address []byte) (*AddressResponse, error) {...}
````

### Example

```
response, err := api.GetAddress("Mxfe60014a6e9ac91618f5d1cab3fd58cded61ee99")

// &api.AddressResponse{Result:struct { Balance map[string]string "json:\"balance\""; TransactionCount string "json:\"transaction_count\"" }{Balance:map[string]string{"CAPITAL":"57010462073783319332082", "KLM0VCOIN":"16619033694080914686", "MNT":"42229514740835940564509"}, TransactionCount:"37"}}
```

### getNonce

Returns next transaction number (nonce) of an address.

```
func (a *Api) GetAddressNonce(address []byte) (uint64, error) {...}
```

###### Example

```
nonce, err := api.GetAddressNonce([]byte("Mxeeee1973381ab793719fff497b9a516719fcd5a2"))
```

### Send

Returns the result of sending <b>signed</b> tx.

```
func (a *Api) Send(transaction transaction.SignedTransaction) (*SendResponse, error) {...}
```

###### Example

```
res, err := api.Send(signedTransaction)

// &api.SendResponse{Result:struct { Code int "json:\"code\""; Data string "json:\"data\""; Log string "json:\"log\""; Hash string "json:\"hash\"" }{Code:0, Data:"", Log:"", Hash:"E0C49AAAC9D32108CF1328D8E398CBFFB6230BFEBFA9A4CBF43E9748F3D0D088"}, Error:struct { Code int "json:\"code\""; Message string "json:\"message\""; TxResult struct { Code int "json:\"code\""; Log string "json:\"log\"" } "json:\"tx_result\"" }{Code:0, Message:"", TxResult:struct { Code int "json:\"code\""; Log string "json:\"log\"" }{Code:0, Log:""}}}
```

## Tests

To run unit tests: 

```shell script
go test ./...
go test ./... -tags=integration
```