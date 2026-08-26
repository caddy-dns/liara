Liara module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Liara.

## Caddy module name

```
dns.providers.liara
```

## Config examples

> [!IMPORTANT]
> Obtain your api-key from [Liara's console](https://console.liara.ir/API).

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "liara",
				"api_token": "secret"
				"ttl": 120,
			}
		}
	}
}
```

or with the Caddyfile:

<!-- ``` -->
<!-- # globally -->
<!-- { -->
<!-- 	acme_dns liara ... -->

<!-- } -->
<!-- ``` -->
```
# one site
tls {
	dns liara <api-token>
	dns_ttl 120s
}
```

> [!WARNING]
> Always set the ttl parameter, as Liara's api only accept certain TTL
> values. Pick one of these [120, 180, 300, 600, 900].

> [!IMPORTANT]
> You can use the OS environment variables in the Caddyfile. [Here is how](https://caddyserver.com/docs/caddyfile/concepts#environment-variables)
