// Copyright (c) 2026 kryptonmourog
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package matrix

import (
	"context"
	"errors"
	"testing"
	"time"

	"maunium.net/go/mautrix/bridgev2/bridgeconfig"
)

func TestWaitForMediaConfig(t *testing.T) {
	br := NewConnector(&bridgeconfig.Config{})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	if err := br.WaitForMediaConfig(ctx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("WaitForMediaConfig() before fetching = %v, want deadline exceeded", err)
	}

	close(br.mediaConfigFetched)
	if err := br.WaitForMediaConfig(context.Background()); err != nil {
		t.Fatalf("WaitForMediaConfig() after fetching = %v", err)
	}

	if err := (&Connector{}).WaitForMediaConfig(context.Background()); err != nil {
		t.Fatalf("WaitForMediaConfig() of a connector that never fetches = %v", err)
	}
}
