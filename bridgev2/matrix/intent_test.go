// Copyright (c) 2026 kryptonmourog
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package matrix

import (
	"testing"

	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/bridgev2/bridgeconfig"
)

func TestShouldEncryptCreatedRoom(t *testing.T) {
	tests := []struct {
		name     string
		all      bool
		personal bool
		portal   bool
		direct   bool
		want     bool
	}{
		{name: "disabled", portal: true, direct: true, want: false},
		{name: "all rooms", all: true, want: true},
		{name: "personal portal", personal: true, portal: true, direct: true, want: true},
		{name: "personal non-portal", personal: true, direct: true, want: false},
		{name: "personal group portal", personal: true, portal: true, direct: false, want: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cfg := &bridgeconfig.Config{}
			cfg.Encryption.Default = test.all
			cfg.Encryption.Personal = test.personal
			intent := &ASIntent{Connector: &Connector{Config: cfg}}
			req := &mautrix.ReqCreateRoom{BeeperPortal: test.portal, IsDirect: test.direct}
			if got := intent.shouldEncryptCreatedRoom(req); got != test.want {
				t.Fatalf("shouldEncryptCreatedRoom() = %t, want %t", got, test.want)
			}
		})
	}
}
