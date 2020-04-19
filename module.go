/**
 * Copyright 2020 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

import (
	"encoding/json"
	"fmt"
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

func (mod Module) depMap() map[string]map[string]interface{} {
	deps := map[string]map[string]interface{}{}
	for _, dep := range mod.Dependencies {
		deps[dep.GetUniqueName()] = dep.GetTrunc()
		moarDeps := dep.depMap()
		for name, val := range moarDeps {
			deps[name] = val
		}
	}
	return deps
}

func (mod Module) FlatDeps() []map[string]interface{} {
	deps := mod.depMap()
	out := []map[string]interface{}{}
	for _, vals := range deps {
		out = append(out, vals)
	}
	return out
}

func (mod Module) ToText() string {
	return fmt.Sprintf("name=%s url=%s ", mod.GetName(), mod.GetURL()) +
		fmt.Sprintf(" index=%s licenses=%v", mod.GetSubIndex(), mod.Licenses)
}

func (mod Module) IsFresh() bool {
	return !mod.NotFresh
}

func (mod *Module) Add(key string, val interface{}) {
	if mod.Meta == nil {
		mod.Meta = map[string]interface{}{}
	}
	mod.Meta[key] = val
}

func (mod Module) getURL() string {
	if mod.URL != "" {
		return mod.URL
	}
	return "https://" + mod.Name
}

func (mod Module) GetURL() string {
	out := mod.getURL()
	if !strings.HasPrefix(out, "http") {
		return "https://" + out
	}
	return out
}

func (mod *Module) UpdateVersion(vers Version) {
	if vers.IsHashSet() {
		mod.Hash = vers.Hash
	} else {
		mod.Version = vers.Branch
	}
}

func (mod Module) GetVersion() Version {
	return Version{
		Hash:   mod.Hash,
		Branch: mod.Version,
	}
}

func (mod Module) IsSame(name string) bool {
	if len(mod.Name) == len(name) && mod.Name == name {
		return true
	}
	if len(mod.Name) > len(name) && strings.HasPrefix(mod.Name, name) {
		return true
	}
	if len(mod.Name) < len(name) && strings.HasPrefix(name, mod.Name) {
		return true
	}

	if mod.NormalName == "" {
		return false
	}

	if len(mod.NormalName) == len(name) && mod.NormalName == name {
		return true
	}
	if len(mod.NormalName) > len(name) && strings.HasPrefix(mod.NormalName, name) {
		return true
	}
	if len(mod.NormalName) < len(name) && strings.HasPrefix(name, mod.NormalName) {
		return true
	}
	return false
}
