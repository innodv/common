/**
 * Copyright 2020 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

import (
	"encoding/json"
	"time"

	"firebase.google.com/go/auth"
)

type Request struct {
	ID         string    `json:"id,omitempty" cli:"id,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty" cli:"timestamp,omitempty"`
	URL        string    `json:"url,omitempty" cli:"url,omitempty"`
	Version    string    `json:"version,omitempty" cli:"version,omitempty"`
	Hash       string    `json:"hash,omitempty" cli:"hash,omitempty"`
	VCS        string    `json:"vcs,omitempty" cli:"vcs,omitempty"`
	source     string
	Name       string                 `json:"name,omitempty" cli:"name,omitempty"`
	NormalName string                 `json:"-,omitempty" cli:"-,omitempty"`
	Meta       map[string]interface{} `json:"meta,omitempty" cli:"meta,omitempty"`
	Token      auth.Token             `json:"token,omitempty" cli:"token,omitempty"`
}

func (req Request) MarshalFirebase() (map[string]interface{}, error) {
	tmp, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	out := map[string]interface{}{}
	err = json.Unmarshal(tmp, &out)
	if err != nil {
		return nil, err
	}
	delete(out, "token")
	return out, nil
}
