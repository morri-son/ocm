// Copyright 2020 Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"github.com/gardener/ocm/cmds/ocm/cmd/core"
	"github.com/gardener/ocm/pkg/common"
	"github.com/gardener/ocm/pkg/config"
	cfgcpi "github.com/gardener/ocm/pkg/config/cpi"
	"github.com/gardener/ocm/pkg/errors"
	ocicpi "github.com/gardener/ocm/pkg/oci/cpi"
	ocmcpi "github.com/gardener/ocm/pkg/ocm/cpi"
	"github.com/gardener/ocm/pkg/runtime"
)

const (
	OCMCmdConfigType   = "ocm.cmd.config" + common.TypeGroupSuffix
	OCMCmdConfigTypeV1 = OCMCmdConfigType + runtime.VersionSeparator + "v1"
)

func init() {
	cfgcpi.RegisterConfigType(OCMCmdConfigType, cfgcpi.NewConfigType(OCMCmdConfigType, &ConfigSpec{}))
	cfgcpi.RegisterConfigType(OCMCmdConfigTypeV1, cfgcpi.NewConfigType(OCMCmdConfigTypeV1, &ConfigSpec{}))
}

// ConfigSpec describes a memory based repository interface.
type ConfigSpec struct {
	runtime.ObjectVersionedType `json:",inline"`
	OCMRepositories             map[string]*ocmcpi.GenericRepositorySpec `json:"ocmRepositories,omitempty"`
	OCIRepositories             map[string]*ocicpi.GenericRepositorySpec `json:"ociRepositories,omitempty"`
}

// NewConfigSpec creates a new memory ConfigSpec
func NewConfigSpec() *ConfigSpec {
	return &ConfigSpec{
		ObjectVersionedType: runtime.NewVersionedObjectType(OCMCmdConfigType),
	}
}

func (a *ConfigSpec) GetType() string {
	return OCMCmdConfigType
}

func (a *ConfigSpec) AddOCIRepository(name string, spec ocicpi.RepositorySpec) error {
	g, err := ocicpi.ToGenericRepositorySpec(spec)
	if err != nil {
		return err
	}
	if a.OCIRepositories == nil {
		a.OCIRepositories = map[string]*ocicpi.GenericRepositorySpec{}
	}
	a.OCIRepositories[name] = g
	return nil
}

func (a *ConfigSpec) AddOCMRepository(name string, spec ocmcpi.RepositorySpec) error {
	g, err := ocmcpi.ToGenericRepositorySpec(spec)
	if err != nil {
		return err
	}
	if a.OCMRepositories == nil {
		a.OCMRepositories = map[string]*ocmcpi.GenericRepositorySpec{}
	}

	a.OCMRepositories[name] = g
	return nil
}

func (a *ConfigSpec) ApplyTo(ctx config.Context, target interface{}) error {
	list := errors.ErrListf("applying config")
	t, ok := target.(core.Context)
	if !ok {
		return config.ErrNoContext(OCMCmdConfigType)
	}
	for n, s := range a.OCIRepositories {
		list.Add(t.AddOCIRepository(n, s))
	}
	for n, s := range a.OCMRepositories {
		list.Add(t.AddOCMRepository(n, s))
	}
	return list.Result()
}