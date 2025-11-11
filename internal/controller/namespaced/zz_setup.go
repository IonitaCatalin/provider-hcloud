// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	network "github.com/IonitaCatalin/provider-hcloud/internal/controller/namespaced/hcloud/network"
	server "github.com/IonitaCatalin/provider-hcloud/internal/controller/namespaced/hcloud/server"
	route "github.com/IonitaCatalin/provider-hcloud/internal/controller/namespaced/network/route"
	subnet "github.com/IonitaCatalin/provider-hcloud/internal/controller/namespaced/network/subnet"
	providerconfig "github.com/IonitaCatalin/provider-hcloud/internal/controller/namespaced/providerconfig"
	networkserver "github.com/IonitaCatalin/provider-hcloud/internal/controller/namespaced/server/network"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		network.Setup,
		server.Setup,
		route.Setup,
		subnet.Setup,
		providerconfig.Setup,
		networkserver.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		network.SetupGated,
		server.SetupGated,
		route.SetupGated,
		subnet.SetupGated,
		providerconfig.SetupGated,
		networkserver.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
