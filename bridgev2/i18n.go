// Copyright (c) 2026 kryptonmourog
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package bridgev2

import (
	"context"

	"maunium.net/go/mautrix/id"
)

// Translate translates user-facing text of the bridge (command replies, help, notices) for the given
// user. msg is the English text as written in the code, before any formatting arguments are applied,
// so it can be used as the key of a translation catalog. It returns msg unchanged by default; network
// connectors can replace it to localize the bridge.
var Translate = func(ctx context.Context, userID id.UserID, msg string) string {
	return msg
}
