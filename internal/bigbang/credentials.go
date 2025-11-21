// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2024-Present Defense Unicorns

// Package bigbang contains the logic for installing Big Bang and Flux
package bigbang

import (
	"strings"

	"github.com/Masterminds/semver/v3"
)

const bbV1ZarfCredentialsValues = `apiVersion: v1
kind: Secret
metadata:
  name: zarf-credentials
  namespace: bigbang
stringData:
  values.yaml: |
    registryCredentials:
      registry: "###ZARF_REGISTRY###"
      username: "###REGISTRY_USERNAME###"
      password: "###ZARF_REGISTRY_AUTH_PULL###"
    git:
      existingSecret: "private-git-server"	# -- Chart created secrets with user defined values
      credentials:
        username: "###ZARF_GIT_PUSH###" # -- HTTP git credentials, both username and password must be provided
        password: "###ZARF_GIT_AUTH_PUSH###"
		# -- Big Bang v1 Kyverno Support
	kyvernopolicies:
	  values:
	    exclude:
		  any:
		  - resources:
			namespaces:
			  - zarf # don't have Kyverno prevent Zarf from doing zarf things
	`

const bbV2ZarfCredentialsValues = `apiVersion: v1
kind: Secret
metadata:
  name: zarf-credentials
  namespace: bigbang
stringData:
  values.yaml: |
    registryCredentials:
      registry: "###ZARF_REGISTRY###"
      username: "###REGISTRY_USERNAME###"
      password: "###ZARF_REGISTRY_AUTH_PULL###"
    git:
      existingSecret: "private-git-server"	# -- Chart created secrets with user defined values
      credentials:
        username: "###ZARF_GIT_PUSH###" # -- HTTP git credentials, both username and password must be provided
        password: "###ZARF_GIT_AUTH_PUSH###"
    kyvernoPolicies:
      values:
        exclude:
          any:
          - resources:
            namespaces:
            - zarf # don't have Kyverno prevent Zarf from doing zarf things
`

func manifestZarfCredentials(version string, registryUsername string) (string, error) {
	semverVersion, err := semver.NewVersion(version)
	if err != nil {
		return "", err
	}
	var template string
	if semverVersion.Major() == 2 {
		template = bbV2ZarfCredentialsValues
	} else {
		template = bbV1ZarfCredentialsValues
	}
	// Replace the registry username placeholder with the actual value
	return strings.ReplaceAll(template, "###REGISTRY_USERNAME###", registryUsername), nil
}
