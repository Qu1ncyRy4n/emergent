// Copyright (c) 2023, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package econfig

import "cogentcore.org/core/base/reflectx"

// SetFromDefaults sets Config values from field tag `def:` values.
// Parsing errors are automatically logged.
func SetFromDefaults(cfg any) error {
	return reflectx.SetFromDefaultTags(cfg)
}
