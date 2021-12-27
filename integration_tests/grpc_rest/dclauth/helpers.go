// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dclauth

import (
	"context"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/go-bip39"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	dclauthtypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/dclauth/types"

	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/utils"
)

//nolint:godox
/*
	To Run test you need:
		* Run LocalNet with: `make install && make localnet_init && make localnet_start`

	TODO: provide tests for error cases
*/

func GetAccount(suite *utils.TestSuite, address sdk.AccAddress) (*dclauthtypes.Account, error) {
	var res dclauthtypes.Account

	if suite.Rest {
		// TODO issue 99: explore the way how to get the endpoint from proto-
		//      instead of the hard coded value (the same for all rest queries)
		var resp dclauthtypes.QueryGetAccountResponse
		err := suite.QueryREST("/dcl/auth/account/"+address.String(), &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetAccount()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.Account(
			context.Background(),
			&dclauthtypes.QueryGetAccountRequest{Address: address.String()},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetAccount()
	}

	return &res, nil

}

func GetAccounts(suite *utils.TestSuite) (res []dclauthtypes.Account, err error) {
	if suite.Rest {
		// TODO issue 99: explore the way how to get the endpoint from proto-
		//      instead of the hard coded value (the same for all rest queries)
		var resp dclauthtypes.QueryAllAccountResponse
		err := suite.QueryREST("/dcl/auth/accounts", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetAccount()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.AccountAll(
			context.Background(),
			&dclauthtypes.QueryAllAccountRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetAccount()
	}

	return res, nil
}

func GetProposedAccounts(suite *utils.TestSuite) (res []dclauthtypes.PendingAccount, err error) {
	if suite.Rest {
		// TODO issue 99: explore the way how to get the endpoint from proto-
		//      instead of the hard coded value (the same for all rest queries)
		var resp dclauthtypes.QueryAllPendingAccountResponse
		err := suite.QueryREST("/dcl/auth/accounts/proposed", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetPendingAccount()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.PendingAccountAll(
			context.Background(),
			&dclauthtypes.QueryAllPendingAccountRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetPendingAccount()
	}

	return res, nil
}

func GetProposedAccountsToRevoke(suite *utils.TestSuite) (
	res []dclauthtypes.PendingAccountRevocation, err error,
) {
	if suite.Rest {
		// TODO issue 99: explore the way how to get the endpoint from proto-
		//      instead of the hard coded value (the same for all rest queries)
		var resp dclauthtypes.QueryAllPendingAccountRevocationResponse
		err := suite.QueryREST("/dcl/auth/accounts/proposed/revoked", &resp)
		if err != nil {
			return nil, err
		}
		res = resp.GetPendingAccountRevocation()
	} else {
		grpcConn := suite.GetGRPCConn()
		defer grpcConn.Close()

		// This creates a gRPC client to query the x/dclauth service.
		accClient := dclauthtypes.NewQueryClient(grpcConn)
		resp, err := accClient.PendingAccountRevocationAll(
			context.Background(),
			&dclauthtypes.QueryAllPendingAccountRevocationRequest{},
		)
		if err != nil {
			return nil, err
		}
		res = resp.GetPendingAccountRevocation()
	}

	return res, nil
}

// TODO issue 99: add support for query accounts stat

func ProposeAddAccount(
	suite *utils.TestSuite,
	accAddr sdk.AccAddress,
	accKey cryptotypes.PubKey,
	roles dclauthtypes.AccountRoles,
	vendorID uint64,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {

	msg, err := dclauthtypes.NewMsgProposeAddAccount(
		suite.GetAddress(signerName), accAddr, accKey, roles, vendorID)
	require.NoError(suite.T, err)
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func ApproveAddAccount(
	suite *utils.TestSuite,
	accAddr sdk.AccAddress,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {

	msg := dclauthtypes.NewMsgApproveAddAccount(suite.GetAddress(signerName), accAddr)
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func ProposeRevokeAccount(
	suite *utils.TestSuite,
	accAddr sdk.AccAddress,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {

	msg := dclauthtypes.NewMsgProposeRevokeAccount(suite.GetAddress(signerName), accAddr)
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func ApproveRevokeAccount(
	suite *utils.TestSuite,
	accAddr sdk.AccAddress,
	signerName string,
	signerAccount *dclauthtypes.Account,
) (*sdk.TxResponse, error) {

	msg := dclauthtypes.NewMsgApproveRevokeAccount(suite.GetAddress(signerName), accAddr)
	return suite.BuildAndBroadcastTx([]sdk.Msg{msg}, signerName, signerAccount)
}

func CreateAccountInfo(suite *utils.TestSuite, accountName string) keyring.Info {
	entropySeed, err := bip39.NewEntropy(256)
	require.NoError(suite.T, err)

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	require.NoError(suite.T, err)

	accountInfo, err := suite.Kr.NewAccount(accountName, mnemonic, testconstants.Passphrase, sdk.FullFundraiserPath, hd.Secp256k1)
	require.NoError(suite.T, err)

	return accountInfo
}

func CreateAccount(
	suite *utils.TestSuite,
	accountName string,
	roles dclauthtypes.AccountRoles,
	vendorId uint64,
	proposerName string,
	proposerAccount *dclauthtypes.Account,
	approverName string,
	approverAccount *dclauthtypes.Account,
) *dclauthtypes.Account {

	accountInfo := CreateAccountInfo(suite, accountName)

	_, err := ProposeAddAccount(
		suite,
		accountInfo.GetAddress(),
		accountInfo.GetPubKey(),
		dclauthtypes.AccountRoles{dclauthtypes.Vendor},
		vendorId,
		proposerName,
		proposerAccount,
	)
	require.NoError(suite.T, err)

	_, err = ApproveAddAccount(
		suite,
		accountInfo.GetAddress(),
		approverName,
		approverAccount,
	)
	require.NoError(suite.T, err)

	account, err := GetAccount(suite, accountInfo.GetAddress())
	require.NoError(suite.T, err)

	return account
}

// Common Test Logic

//nolint:funlen
func AuthDemo(suite *utils.TestSuite) {
	// Jack, Alice and Bob are predefined Trustees
	jackName := testconstants.JackAccount
	jackKeyInfo, err := suite.Kr.Key(jackName)
	require.NoError(suite.T, err)
	jackAccount, err := GetAccount(suite, jackKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	aliceName := testconstants.AliceAccount
	aliceKeyInfo, err := suite.Kr.Key(aliceName)
	require.NoError(suite.T, err)
	aliceAccount, err := GetAccount(suite, aliceKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	bobName := testconstants.BobAccount
	bobKeyInfo, err := suite.Kr.Key(bobName)
	require.NoError(suite.T, err)
	bobAccount, err := GetAccount(suite, bobKeyInfo.GetAddress())
	require.NoError(suite.T, err)

	// Query all active accounts
	inputAccounts, err := GetAccounts(suite)
	require.NoError(suite.T, err)

	// Query all proposed accounts
	inputProposedAccounts, err := GetProposedAccounts(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 0, len(inputProposedAccounts))

	// Query all proposed accounts to revoke
	inputProposedAccountsToRevoke, err := GetProposedAccountsToRevoke(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, 0, len(inputProposedAccountsToRevoke))

	_, testAccPubKey, testAccAddr := testdata.KeyTestPubAddr()

	// Jack proposes new account
	_, err = ProposeAddAccount(
		suite,
		testAccAddr, testAccPubKey,
		dclauthtypes.AccountRoles{dclauthtypes.Vendor}, uint64(testconstants.Vid),
		jackName, jackAccount,
	)
	require.NoError(suite.T, err)

	// Query all active accounts
	receivedAccounts, _ := GetAccounts(suite)
	require.Equal(suite.T, len(inputAccounts), len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ := GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts)+1, len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ := GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Alice approves new account
	_, err = ApproveAddAccount(suite, testAccAddr, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Query all active accounts
	receivedAccounts, _ = GetAccounts(suite)
	require.Equal(suite.T, len(inputAccounts)+1, len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Get new account
	testAccount, err := GetAccount(suite, testAccAddr)
	require.NoError(suite.T, err)
	require.Equal(suite.T, testAccAddr, testAccount.GetAddress())
	require.Equal(suite.T, []dclauthtypes.AccountRole{dclauthtypes.Vendor}, testAccount.GetRoles())

	// FIXME issue 99: enable once implemented
	/*
		// Publish model info by test account
		model := NewMsgAddModel(suite, testAccountKeyInfo.Address, testconstants.VID)
		_, _ = AddModel(suite, model, testAccountKeyInfo)

		// Check model is created
		receivedModel, _ := GetModel(suite, model.VID, model.PID)
		require.Equal(suite.T, receivedModel.VID, model.VID)
		require.Equal(suite.T, receivedModel.PID, model.PID)
		require.Equal(suite.T, receivedModel.ProductName, model.ProductName)
	*/

	// Alice proposes to revoke new account
	_, err = ProposeRevokeAccount(suite, testAccAddr, aliceName, aliceAccount)
	require.NoError(suite.T, err)

	// Query all active accounts
	receivedAccounts, err = GetAccounts(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, len(inputAccounts)+1, len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke)+1, len(receivedProposedAccountsToRevoke))

	// Bob approves to revoke new account
	_, err = ApproveRevokeAccount(suite, testAccAddr, bobName, bobAccount)
	require.NoError(suite.T, err)

	// Query all active accounts
	receivedAccounts, err = GetAccounts(suite)
	require.NoError(suite.T, err)
	require.Equal(suite.T, len(inputAccounts), len(receivedAccounts))

	// Query all proposed accounts
	receivedProposedAccounts, _ = GetProposedAccounts(suite)
	require.Equal(suite.T, len(inputProposedAccounts), len(receivedProposedAccounts))

	// Query all accounts proposed to be revoked
	receivedProposedAccountsToRevoke, _ = GetProposedAccountsToRevoke(suite)
	require.Equal(suite.T, len(inputProposedAccountsToRevoke), len(receivedProposedAccountsToRevoke))

	// Ensure that new account is not available anymore
	_, err = GetAccount(suite, testAccAddr)
	require.Error(suite.T, err)
	require.Contains(suite.T, err.Error(), "rpc error: code = InvalidArgument desc = not found: invalid request")

	// FIXME issue 99: enable once implemented
	/*
		// Try to publish another model info by test account.
		// Ensure that the request is responded with not OK status code.
		model = NewMsgAddModel(suite, testAccountKeyInfo.Address, testconstants.VID)
		_, code = AddModel(suite, model, testAccountKeyInfo)
		require.NotEqual(suite.T, http.StatusOK, code)
	*/
}
