/**
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package dhcplb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRREmpty(t *testing.T) {
	subject := new(roundRobin)
	_, err := subject.SelectRatioBasedDhcpServer(&DHCPMessage{
		ClientID: []byte{0},
	})
	require.Error(t, err)
}

func TestRRBalance(t *testing.T) {
	subject := new(roundRobin)
	servers := make([]*DHCPServer, 4)
	for i := 0; i < 4; i++ {
		servers[i] = &DHCPServer{
			Port: i,
		}
	}
	subject.UpdateStableServerList(servers)
	msg := DHCPMessage{
		ClientID: []byte{0},
	}
	for i := 0; i < 4; i++ {
		server, err := subject.SelectRatioBasedDhcpServer(&msg)
		require.NoError(t, err)
		require.Equal(t, i, server.Port)
	}
}
