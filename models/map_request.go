package models

import (
	"errors"
)

type MapRequest struct {
	X         uint32
	Y         uint32
	Z         uint32
	LayerName string
}

func MapRequetDecode(params any) (*MapRequest, error) {
	args, ok := params.(map[string]any)
	if !ok {
		return nil, errors.New("failed to parse request")
	}

	var X, Y, Z uint32
	if X, ok = args["_x"].(uint32); !ok {
		return nil, errors.New("failed to parse request")
	}

	if Y, ok = args["_y"].(uint32); !ok {
		return nil, errors.New("failed to parse request")
	}

	if Z, ok = args["_z"].(uint32); !ok {
		return nil, errors.New("failed to parse request")
	}

	var TileName string
	if TileName, ok = args["layer"].(string); !ok {
		return nil, errors.New("failed to parse request")
	}

	return &MapRequest{
		X,
		Y,
		Z,
		TileName,
	}, nil
}
