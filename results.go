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
	Error    string   `json:"error,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	Path     string   `json:"path,omitempty"`
}

type LicenseResult struct {
	Name     string                 `json:"name,omitempty"`
	Hash     string                 `json:"hash,omitempty"`
	Version  string                 `json:"version,omitempty"`
	DateTime string                 `json:"datetime,omitempty"`
	Meta     map[string]interface{} `json:"meta,omitempty"`
	Licenses []string               `json:"licenses,omitempty"`
	Counts   map[string]int64       `json:"counts,omitempty"`
	Indeces  []string               `json:"indeces,omitempty"`
}
