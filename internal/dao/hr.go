// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/derailed/k9s/internal/client"
	helmv2 "github.com/fluxcd/helm-controller/api/v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
)

type HelmRelease struct {
	Resource
}

func (h *HelmRelease) Suspend(ctx context.Context, path string) error {
	o, err := h.getFactory().Get("helm.toolkit.fluxcd.io/v2", path, true, labels.Everything())
	if err != nil {
		return err
	}

	var hr helmv2.HelmRelease
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(o.(*unstructured.Unstructured).Object, &hr)
	if err != nil {
		return err
	}

	auth, err := h.Client().CanI(hr.Namespace, "helm.toolkit.fluxcd.io/v2/helmreleases:update", hr.Name, client.PatchAccess)
	if err != nil {
		return err
	}
	if !auth {
		return fmt.Errorf("user is not authorized to suspend a HelmRelease")
	}

	dial, err := h.Client().Dial()
	if err != nil {
		return err
	}

	//hr.Spec.Suspend = true
}

// GetInstance returns a HelmRelease instance.
func (h *HelmRelease) GetInstance(fqn string) (*helmv2.HelmRelease, error) {
	o, err := h.Factory.Get(h.GVR(), fqn, true, labels.Everything())
	if err != nil {
		return nil, err
	}

	var hr helmv2.HelmRelease
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(o.(*unstructured.Unstructured).Object, &hr)

	if err != nil {
		return nil, errors.New("expecting HelmRelease resource")
	}

	return &hr, nil
}
