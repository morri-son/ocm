// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package oci_test

import (
	"github.com/gardener/ocm/pkg/oci"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontainers/go-digest"
)

func CheckRef(ref string, exp *oci.RefSpec) {
	spec, err := oci.ParseRef(ref)
	if exp == nil {
		Expect(err).To(HaveOccurred())
	} else {
		Expect(err).To(Succeed())
		Expect(spec).To(Equal(*exp))
	}
}

var _ = Describe("ref parsing", func() {
	digest := digest.Digest("sha256:3d05e105e350edf5be64fe356f4906dd3f9bf442a279e4142db9879bba8e677a")
	tag := "v1"

	It("succeeds", func() {
		CheckRef("https://ubuntu", &oci.RefSpec{Scheme: "https", Host: "docker.io", Repository: "library/ubuntu"})
		CheckRef("ubuntu", &oci.RefSpec{Host: "docker.io", Repository: "library/ubuntu"})
		CheckRef("ubuntu:v1", &oci.RefSpec{Host: "docker.io", Repository: "library/ubuntu", Tag: &tag})
		CheckRef("test/ubuntu", &oci.RefSpec{Host: "docker.io", Repository: "test/ubuntu"})
		CheckRef("test/ubuntu:v1", &oci.RefSpec{Host: "docker.io", Repository: "test/ubuntu", Tag: &tag})
		CheckRef("ghcr.io/test/ubuntu", &oci.RefSpec{Host: "ghcr.io", Repository: "test/ubuntu"})
		CheckRef("ghcr.io:8080/test/ubuntu", &oci.RefSpec{Host: "ghcr.io:8080", Repository: "test/ubuntu"})
		CheckRef("ghcr.io/test/ubuntu:v1", &oci.RefSpec{Host: "ghcr.io", Repository: "test/ubuntu", Tag: &tag})
		CheckRef("ghcr.io/test/ubuntu@sha256:3d05e105e350edf5be64fe356f4906dd3f9bf442a279e4142db9879bba8e677a", &oci.RefSpec{Host: "ghcr.io", Repository: "test/ubuntu", Digest: &digest})
		CheckRef("ghcr.io/test/ubuntu:v1@sha256:3d05e105e350edf5be64fe356f4906dd3f9bf442a279e4142db9879bba8e677a", &oci.RefSpec{Host: "ghcr.io", Repository: "test/ubuntu", Tag: &tag, Digest: &digest})
	})

	It("fails", func() {
		CheckRef("ubuntu@4711", nil)
		CheckRef("test/ubuntu@4711", nil)
		CheckRef("test/ubuntu:v1@4711", nil)
		CheckRef("ghcr.io/test/ubuntu:v1@4711", nil)

	})
})