package services

import (
	"context"
	"encoding/hex"
	"github.com/gho1b/pantrack/proto-gen/ptypes"
	"github.com/gho1b/pantrack/proto-gen/services"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/micro/go-micro/v3/client"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LocationType struct {
	services.LocationService
}

func (l *LocationType) GetLocationType(ctx context.Context, in *wrappers.BytesValue, opts ...client.CallOption) (*ptypes.LocationType, error) {
	id, _ := hex.DecodeString(primitive.NewObjectID().Hex())
	return &ptypes.LocationType{
		Id:     id,
		Parent: nil,
		Level:  1,
		Name:   "Country",
	}, nil
}
