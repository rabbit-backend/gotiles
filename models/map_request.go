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

	return &MapRequest{
		X:         uint32((args["_x"].(int))),
		Y:         uint32(args["_y"].(int)),
		Z:         uint32((args["_z"].(int))),
		LayerName: args["_layer"].(string),
	}, nil
}
