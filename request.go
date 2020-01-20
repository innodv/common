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
	ID         string    `json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	URL        string    `json:"url"`
	Version    string    `json:"version"`
	Hash       string    `json:"hash"`
	VCS        string    `json:"vcs"`
	source     string
	Name       string                 `json:"name"`
	NormalName string                 `json:"-"`
	Meta       map[string]interface{} `json:"meta"`
	Token      auth.Token             `json:"token"`
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
