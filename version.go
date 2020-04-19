/**
 * Copyright 2020 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

type Version struct {
	Branch string
	Hash   string
}

func (vers Version) IsHashSet() bool {
	return vers.Hash != ""
}

func (vers Version) String() string {
	if vers.IsHashSet() {
		return vers.Hash
	}
	return vers.Branch
}
