/**
 * Copyright 2020 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

type Result struct {
	ExecutionResult
	LicenseResult
}

type ExecutionResult struct {
	Success  bool     `json:"success" cli:"success"`
	Error    string   `json:"error,omitempty" cli:"error,omitempty"`
	Warnings []string `json:"warnings,omitempty" cli:"warnings,omitempty"`
	Path     string   `json:"path,omitempty" cli:"path,omitempty"`
}

type LicenseResult struct {
	Name     string                 `json:"name,omitempty" cli:"name,omitempty"`
	Hash     string                 `json:"hash,omitempty" cli:"hash,omitempty"`
	Version  string                 `json:"version,omitempty" cli:"version,omitempty"`
	DateTime string                 `json:"datetime,omitempty" cli:"datetime,omitempty"`
	Meta     map[string]interface{} `json:"meta,omitempty" cli:"meta,omitempty"`
	Licenses []string               `json:"licenses,omitempty" cli:"licenses,omitempty"`
	Counts   map[string]int64       `json:"counts,omitempty" cli:"counts,omitempty"`
	Indeces  []string               `json:"indeces,omitempty" cli:"indeces,omitempty"`
}

type PRContext struct {
	Provider string //e.g. github
	Owner    string
	Repo     string
	Number   int
}

func (prc PRContext) Empty() bool {
	return prc.Provider == "" && prc.Owner == "" && prc.Repo == "" && prc.Number == 0
}

type RawResult struct {
	Module Module
	Error  error
	UID    string
	ReqID  string
	PR     PRContext
}
