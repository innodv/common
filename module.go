/**
 * Copyright 2020 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

type License string

type Module struct {
	Name       string                 `json:"name,omitempty"`
	NormalName string                 `json:"normalName,omitempty"`
	URL        string                 `json:"url,omitempty"`
	Hash       string                 `json:"hash,omitempty"`
	Version    string                 `json:"version,omitempty"`
	DateTime   string                 `json:"datetime,omitempty"`
	Meta       map[string]interface{} `json:"meta,omitempty"`
	Licenses   []License              `json:"licenses,omitempty"`
	NotFresh   bool                   `json:"-"`
	VCS        string                 `json:"vcs,omitempty"`
	Source     string                 `json:"source,omitempty"`
	// Dirty flags that the data from this module is untrusted and should be
	// stored in the cache
	Dirty bool `json:"dirty,omitempty"`

	Dependencies []Module         `json:"-"`
	Counts       map[string]int64 `json:"counts"`
	CountIsNull  bool             `json:"countIsNull"`
	Indeces      []string         `json:"indeces"`
	Rev          string           `json:"_rev,omitempty"`
}
