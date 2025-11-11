package server

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
    p.AddResourceConfigurator("hcloud_server", func(r *config.Resource) {
			r.ShortGroup = ""
    })

    p.AddResourceConfigurator("hcloud_server_network", func(r *config.Resource) {
        r.References["network"] = config.Reference{
            Type: "github.com/IonitaCatalin/provider-hcloud/apis/hcloud/v1alpha1.Network",
        }
    })
}
