package core

import (
	"testing"

	tlsC "github.com/metacubex/mihomo/component/tls"
	utls "github.com/metacubex/utls"
)

func TestFirefoxFingerprintOffersX25519MLKEM768(t *testing.T) {
	fingerprint, ok := tlsC.GetFingerprint("firefox")
	if !ok {
		t.Fatal("firefox fingerprint is unavailable")
	}
	if fingerprint.Version != "148" {
		t.Fatalf("firefox fingerprint version = %q, want 148", fingerprint.Version)
	}

	spec, err := utls.UTLSIdToSpec(fingerprint)
	if err != nil {
		t.Fatalf("build firefox ClientHello spec: %v", err)
	}

	var curves []utls.CurveID
	var shares []utls.KeyShare
	for _, extension := range spec.Extensions {
		switch extension := extension.(type) {
		case *utls.SupportedCurvesExtension:
			curves = extension.Curves
		case *utls.KeyShareExtension:
			shares = extension.KeyShares
		}
	}

	if len(curves) == 0 || curves[0] != utls.X25519MLKEM768 {
		t.Fatalf("first supported group = %v, want X25519MLKEM768", curves)
	}
	if len(shares) < 2 || shares[0].Group != utls.X25519MLKEM768 || shares[1].Group != utls.X25519 {
		t.Fatalf("key shares = %v, want X25519MLKEM768 followed by X25519", shares)
	}
}
