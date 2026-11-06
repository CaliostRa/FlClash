package core

import (
	"testing"

	"github.com/metacubex/mihomo/config"
)

func TestExistingProxyProtocolsStillParse(t *testing.T) {
	const profile = `
mixed-port: 7890
mode: rule
proxies:
  - name: reality-vision
    type: vless
    server: 192.0.2.1
    port: 443
    uuid: 00000000-0000-0000-0000-000000000001
    network: tcp
    tls: true
    flow: xtls-rprx-vision
    client-fingerprint: firefox
    reality-opts:
      public-key: AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
      short-id: "0123456789abcdef"
      support-x25519mlkem768: false
  - name: reality-xhttp
    type: vless
    server: 192.0.2.2
    port: 443
    uuid: 00000000-0000-0000-0000-000000000002
    network: xhttp
    tls: true
    client-fingerprint: firefox
    reality-opts:
      public-key: AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
      short-id: "0123456789abcdef"
    xhttp-opts:
      path: /xhttp
  - name: vless-grpc
    type: vless
    server: 192.0.2.3
    port: 443
    uuid: 00000000-0000-0000-0000-000000000003
    network: grpc
    tls: true
    grpc-opts:
      grpc-service-name: vless
  - name: hysteria2
    type: hysteria2
    server: 192.0.2.4
    port: 443
    password: test-password
    sni: example.com
  - name: trojan-grpc
    type: trojan
    server: 192.0.2.5
    port: 443
    password: test-password
    network: grpc
    sni: example.com
    grpc-opts:
      grpc-service-name: trojan
proxy-groups:
  - name: PROXY
    type: select
    proxies: [reality-vision, reality-xhttp, vless-grpc, hysteria2, trojan-grpc]
rules:
  - MATCH,PROXY
`

	parsed, err := config.Parse([]byte(profile))
	if err != nil {
		t.Fatalf("parse compatibility profile: %v", err)
	}
	for _, name := range []string{
		"reality-vision",
		"reality-xhttp",
		"vless-grpc",
		"hysteria2",
		"trojan-grpc",
	} {
		if _, ok := parsed.Proxies[name]; !ok {
			t.Errorf("proxy %q was not parsed", name)
		}
	}
}
