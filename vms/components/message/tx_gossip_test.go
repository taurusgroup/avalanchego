// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/units"
)

func TestTxGossip(t *testing.T) {
	require := require.New(t)

	tx := utils.RandomBytes(256 * units.KiB)
	builtMsg := TxGossip{
		Tx: tx,
	}
	builtMsgBytes, err := Build(&builtMsg)
	require.NoError(err)
	require.Equal(builtMsgBytes, builtMsg.Bytes())

	parsedMsgIntf, err := Parse(builtMsgBytes)
	require.NoError(err)
	require.Equal(builtMsgBytes, parsedMsgIntf.Bytes())

	parsedMsg, ok := parsedMsgIntf.(*TxGossip)
	require.True(ok)

	require.Equal(tx, parsedMsg.Tx)
}