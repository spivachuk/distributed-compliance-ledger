package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateModel = "create_model"
	TypeMsgUpdateModel = "update_model"
	TypeMsgDeleteModel = "delete_model"
)

var _ sdk.Msg = &MsgCreateModel{}

func NewMsgCreateModel(
	creator string,
	vid int32,
	pid int32,
	deviceTypeId int32,
	productName string,
	productLabel string,
	partNumber string,
	commissioningCustomFlow int32,
	commissioningCustomFlowUrl string,
	commissioningModeInitialStepsHint uint64,
	commissioningModeInitialStepsInstruction string,
	commissioningModeSecondaryStepsHint uint64,
	commissioningModeSecondaryStepsInstruction string,
	userManualUrl string,
	supportUrl string,
	productUrl string,

) *MsgCreateModel {
	return &MsgCreateModel{
		Creator:                                  creator,
		Vid:                                      vid,
		Pid:                                      pid,
		DeviceTypeId:                             deviceTypeId,
		ProductName:                              productName,
		ProductLabel:                             productLabel,
		PartNumber:                               partNumber,
		CommissioningCustomFlow:                  commissioningCustomFlow,
		CommissioningCustomFlowUrl:               commissioningCustomFlowUrl,
		CommissioningModeInitialStepsHint:        commissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: commissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      commissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: commissioningModeSecondaryStepsInstruction,
		UserManualUrl: userManualUrl,
		SupportUrl:    supportUrl,
		ProductUrl:    productUrl,
	}
}

func (msg *MsgCreateModel) Route() string {
	return RouterKey
}

func (msg *MsgCreateModel) Type() string {
	return TypeMsgCreateModel
}

func (msg *MsgCreateModel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateModel{}

func NewMsgUpdateModel(
	creator string,
	vid int32,
	pid int32,
	deviceTypeId int32,
	productName string,
	productLabel string,
	partNumber string,
	commissioningCustomFlow int32,
	commissioningCustomFlowUrl string,
	commissioningModeInitialStepsHint uint64,
	commissioningModeInitialStepsInstruction string,
	commissioningModeSecondaryStepsHint uint64,
	commissioningModeSecondaryStepsInstruction string,
	userManualUrl string,
	supportUrl string,
	productUrl string,

) *MsgUpdateModel {
	return &MsgUpdateModel{
		Creator:                                  creator,
		Vid:                                      vid,
		Pid:                                      pid,
		DeviceTypeId:                             deviceTypeId,
		ProductName:                              productName,
		ProductLabel:                             productLabel,
		PartNumber:                               partNumber,
		CommissioningCustomFlow:                  commissioningCustomFlow,
		CommissioningCustomFlowUrl:               commissioningCustomFlowUrl,
		CommissioningModeInitialStepsHint:        commissioningModeInitialStepsHint,
		CommissioningModeInitialStepsInstruction: commissioningModeInitialStepsInstruction,
		CommissioningModeSecondaryStepsHint:      commissioningModeSecondaryStepsHint,
		CommissioningModeSecondaryStepsInstruction: commissioningModeSecondaryStepsInstruction,
		UserManualUrl: userManualUrl,
		SupportUrl:    supportUrl,
		ProductUrl:    productUrl,
	}
}

func (msg *MsgUpdateModel) Route() string {
	return RouterKey
}

func (msg *MsgUpdateModel) Type() string {
	return TypeMsgUpdateModel
}

func (msg *MsgUpdateModel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteModel{}

func NewMsgDeleteModel(
	creator string,
	vid int32,
	pid int32,

) *MsgDeleteModel {
	return &MsgDeleteModel{
		Creator: creator,
		Vid:     vid,
		Pid:     pid,
	}
}
func (msg *MsgDeleteModel) Route() string {
	return RouterKey
}

func (msg *MsgDeleteModel) Type() string {
	return TypeMsgDeleteModel
}

func (msg *MsgDeleteModel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}