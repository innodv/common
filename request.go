/**
 * Copyright 2020 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

import (
	"encoding/json"
	"time"
)

type Request struct {
	ID         string                 `json:"id,omitempty" cli:"id,omitempty"`
	Timestamp  time.Time              `json:"timestamp,omitempty" cli:"timestamp,omitempty"`
	URL        string                 `json:"url,omitempty" cli:"url,omitempty"`
	Version    string                 `json:"version,omitempty" cli:"version,omitempty"`
	Hash       string                 `json:"hash,omitempty" cli:"hash,omitempty"`
	VCS        string                 `json:"vcs,omitempty" cli:"vcs,omitempty"`
	Source     string                 `json:"-" cli:"-"`
	Name       string                 `json:"name,omitempty" cli:"name,omitempty"`
	NormalName string                 `json:"-" cli:"-"`
	Meta       map[string]interface{} `json:"meta,omitempty" cli:"meta,omitempty"`
	UID        string                 `json:"uid,omitempty" cli:"user,omitempty"`
	PR         PRContext
	AckFn      func() `json:"-"`
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
	return out, nil
}

func NewRequestFromModule(mod Module) Request {
	return Request{
		URL:        mod.GetURL(),
		Version:    mod.Version,
		Hash:       mod.Hash,
		VCS:        mod.GetVCS(),
		Source:     mod.Source,
		Name:       mod.Name,
		NormalName: mod.NormalName,
	}
}

func (req Request) ToModuleShell() Module {
	return Module{
		Name:       req.Name,
		NormalName: req.NormalName,
		URL:        req.URL,
		Version:    req.Version,
		Hash:       req.Hash,
		VCS:        req.VCS,
		Source:     req.Source,
	}
}

func (req Request) Ack() {
	if req.AckFn != nil {
		req.AckFn()
	}
}
