/**
 * Copyright 2020 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

import (
	"encoding/json"
	"strings"
)

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

func (mod Module) GetDBForm() (map[string]interface{}, error) {
	out := map[string]interface{}{}
	tmp, err := json.Marshal(mod)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(tmp, &out)
	if err != nil {
		return nil, err
	}
	delete(out, "indeces")
	out["counts"], _ = mod.DepLicenseCounts()
	out["indeces"] = mod.DepIndeces()
	return out, nil
}

func (mod Module) DepLicenseCounts() (map[string]int64, []string) {
	if len(mod.Dependencies) == 0 {
		return mod.Counts, nil
	}
	out := map[string]int64{}
	noLicense := map[string]bool{}
	for _, dep := range mod.Dependencies {
		for _, license := range dep.Licenses {
			out[string(license)] += 1
		}
		if len(dep.Licenses) == 0 {
			out["Unknown"] += 1

			noLicense[mod.Index()] = true
		}
	}

	for _, dep := range mod.Dependencies {
		newLicenses, newUnknown := dep.DepLicenseCounts()
		for license, count := range newLicenses {
			out[string(license)] += count
		}
		for _, depName := range newUnknown {
			noLicense[depName] = true
		}
	}
	unknown := []string{}
	for dep := range noLicense {
		unknown = append(unknown, dep)
	}

	return out, unknown
}

func (mod Module) depIndeces(depth int) []string {
	out := []string{}
	if depth != 0 {
		out = append(out, mod.Index())
	}
	for i := range mod.Dependencies {
		out = append(out, mod.Dependencies[i].depIndeces(depth+1)...)
	}
	return out
}

func (mod Module) DepIndeces() []string {
	if mod.Indeces != nil {
		return mod.Indeces
	}
	return mod.depIndeces(0)
}

func (mod Module) Index() string {
	return mod.GetName() + "@" + mod.GetSubIndex()
}

func (mod Module) getSubIndex() string {
	if mod.Hash != "" {
		return mod.Hash
	}
	if mod.Version != "" {
		return mod.Version
	}
	return ""
}

func (mod Module) GetSubIndex() string {
	return strings.Replace(mod.getSubIndex(), "/", "-", -1)
}

func (mod Module) GetName() string {
	if mod.NormalName != "" {
		return mod.NormalName
	}
	return mod.Name
}

func (mod Module) GetUniqueName() string {
	return mod.GetName() + mod.GetSubIndex()
}

func (mod Module) GetVCS() string {
	if mod.VCS == "" {
		return "git"
	}
	return mod.VCS
}

func (mod Module) GetTrunc() map[string]interface{} {
	return map[string]interface{}{
		"name":     mod.GetName(),
		"version":  mod.Version,
		"hash":     mod.Hash,
		"licenses": mod.Licenses,
		"vcs":      mod.GetVCS(),
	}
}
