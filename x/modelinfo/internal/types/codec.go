package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module.
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete type on the Amino codec.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAddModelInfo{}, ModuleName+"/AddModelInfo", nil)
	cdc.RegisterConcrete(MsgUpdateModelInfo{}, ModuleName+"/UpdateModelInfo", nil)
	cdc.RegisterConcrete(MsgDeleteModelInfo{}, ModuleName+"/DeleteModelInfo", nil)
}