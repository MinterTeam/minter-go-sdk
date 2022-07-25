package grpc_client

import (
	"github.com/MinterTeam/minter-go-sdk/v2/api"
	_struct "google.golang.org/protobuf/types/known/structpb"
)

// ConvertStructToEvent returns Event model
// Deprecated
func ConvertStructToEvent(str *_struct.Struct) (api.Event, error) {
	value, err := str.Fields["value"].GetStructValue().MarshalJSON()
	if err != nil {
		return nil, err
	}

	event, err := api.ConvertToEvent(api.EventType(str.Fields["type"].GetStringValue()), value)
	if err != nil {
		return nil, err
	}

	return event, nil
}
