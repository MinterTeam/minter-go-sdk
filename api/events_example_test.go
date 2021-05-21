package api_test

import (
	"fmt"

	"github.com/MinterTeam/minter-go-sdk/v2/api"
)

func ExampleConvertToEvent() {
	eventValueJSON := `{
        "address": "Mx7f0fc21d932f38ca9444f61703174569066cfa50",
        "amount": "3962700000000000000000",
        "role": "DAO",
        "validator_pub_key": "Mpd83e627510eea6aefa46d9914b0715dabf4a561ced78d34267b31d41d5f700b5"
      }`
	event, _ := api.ConvertToEvent("minter/RewardEvent", []byte(eventValueJSON))
	if event.Type() != api.TypeRewardEvent {
		return
	}
	fmt.Printf("%#v", event.(*api.RewardEvent))

	// Output:
	// &api.RewardEvent{Role:"DAO", Address:"Mx7f0fc21d932f38ca9444f61703174569066cfa50", Amount:"3962700000000000000000", ValidatorPubKey:"Mpd83e627510eea6aefa46d9914b0715dabf4a561ced78d34267b31d41d5f700b5"}

}
