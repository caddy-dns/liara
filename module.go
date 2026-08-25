package template

import (
	"fmt"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	liara "github.com/libdns/liara"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct {
	*liara.Provider
}

func init() {
	caddy.RegisterModule(Provider{})
}

// All Caddy modules should implements the "caddy.Module" which is a interface
// which provides the module's name and id.
// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "dns.providers.liara",
		New: func() caddy.Module {
			return &Provider{
				new(liara.Provider),
			}
		},
	}
}

// Implements caddy.Provisioner.
// Provision sets up the module and does some additional provisioning steps.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.APIToken = caddy.NewReplacer().ReplaceAll(p.Provider.APIToken, "")

	if p.Provider.APIToken == "" {
		return fmt.Errorf("missing API token")
	}

	return nil
}

// UnmarshalCaddyfile sets up the Liara dns provider from Caddyfile tokens.
//
// It supports two kinds of syntax, Nested and one-liner.
//
// 1.
// liara <api_token>
//
// 2.
//
//	liara {
//	    api_token <api_token>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	// Consume "liara"
	for d.Next() {
		// In the one-line syntax, the next argument after liara should be the api_key
		if d.NextArg() {
			p.Provider.APIToken = d.Val()
		}
		// We don't expect a second argument for this line.
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_token":
				if p.Provider.APIToken != "" {
					return d.Err("API token already set")
				}
				if d.NextArg() {
					p.Provider.APIToken = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIToken == "" {
		return d.Err("missing API token")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
