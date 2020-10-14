package models

// ValueAsMap return map from Value
func (m *EventsResponseEvent) ValueAsMap() map[string]string {
	result := map[string]string{}
	for k, v := range m.Value.(map[string]interface{}) {
		result[k] = v.(string)
	}
	return result
}

type event struct {
	address         string
	validatorPubKey string
}

func (e *event) Address() string {
	return e.address
}
func (e *event) ValidatorPubKey() string {
	return e.validatorPubKey
}

// ValueAsEventInterface return map from Value
func (m *EventsResponseEvent) ValueAsEventInterface() interface {
	Address() string
	ValidatorPubKey() string
} {
	return &event{
		address:         m.ValueAsMap()["address"],
		validatorPubKey: m.ValueAsMap()["validator_pub_key"],
	}
}

// Event type names
const (
	TypeRewardEvent    = "minter/RewardEvent"
	TypeSlashEvent     = "minter/SlashEvent"
	TypeUnbondEvent    = "minter/UnbondEvent"
	TypeStakeKickEvent = "minter/StakeKickEvent"
)

// Events
type (
	RewardEvent struct {
		Role            string `json:"role"`
		Address         string `json:"address"`
		Amount          string `json:"amount"`
		ValidatorPubKey string `json:"validator_pub_key"`
	}
	SlashEvent struct {
		Address         string `json:"address"`
		Amount          string `json:"amount"`
		Coin            string `json:"coin"`
		ValidatorPubKey string `json:"validator_pub_key"`
	}
	UnbondEvent struct {
		Address         string `json:"address"`
		Amount          string `json:"amount"`
		Coin            string `json:"coin"`
		ValidatorPubKey string `json:"validator_pub_key"`
	}
	StakeKickEvent struct {
		Address         string `json:"address"`
		Amount          string `json:"amount"`
		Coin            string `json:"coin"`
		ValidatorPubKey string `json:"validator_pub_key"`
	}
)
