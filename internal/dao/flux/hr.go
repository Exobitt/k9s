// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package dao

import (
	"fmt"
 "github.com/derailed/k9s/internal/client"
)
type HelmRelease struct {
	Resource
}


// Suspend a HelmRelease.
func (h *HelmRelease) Suspend(ctx context.Context, path string) error {
	ns, n := client.Namespaced(path)
	auth, err := h.Client().CanI(ns, "helm.toolkit.fluxcd.io/v2beta2:suspend", n, []string{client.GetVerb, client.UpdateVerb})
	if err != nil {
		return err
	}

	if !auth {
		return fmt.Errorf("user is not authorized to suspend a helmrelease")
	}

