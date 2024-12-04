package cryptoz_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/ibrt/golang-utils/cryptoz"
	"github.com/ibrt/golang-utils/fixturez"
)

const (
	validPrivateKey = `-----BEGIN PRIVATE KEY-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDoXXH6CbQ1T+e1
y+Iex+R2CuITLTx1fkW+F7CnBzOrCEfiRuNcT3sNFqTnSeTRzxhBnGZzMC+X1L90
FU0GVCD95n2lvPLga2zMcKa5qU2BJFz83g59jmIBfHHVsyikci/VV1/0kOMYPEOm
7WttFgMsqgBiC1UaemGTmJXMhs4F7o1gonouFE9RVj1X27ZfnbhwLjDZEQJvp2k6
OzQNY8YRTTlPjBSs69FBJlWUhGjt3zmMJbVXYojDEm5mPzyPFVxiKBHNDBuNU8xj
79il3IrYv7YXfNldQzdBLBbK7ENfqu9t/PkJaOveBvgQLRPo4GX+3t+M9JuJJd2B
osZWNeWrAgMBAAECggEBAJxLjoiy0kYx0xeTZityJRfJRjvD57DYGK0+XhJbY8Od
NEzdhbznsUsiehUgvQrrE9O+EaNVPA4SihzY3xBssixWRxmeOHf/ihURiPPFD17Y
SLvF2VVW2lFJlYA6nBHQxJ/pv59PfZElqBO2CtY7QjNevhc0rC+9NbkDn28NFbMi
dE9woL+RAvoRWsPhEhZ/aLjI+KAZIqmpaRXODxO1MFAORnv6bvGKbrgMXHAO3hAQ
z4EU6uhZ/pVkaRaowmIvQ3Rv9+MSFzi9TbBjPLR8vYfSPeTgY2399v62XyOWWYOc
9TTE3A+q7iVxMqRSKkg06UBgCzwqXFUVqk4iw8cF1AECgYEA9xHLxE1dXGZID9T2
ROr1pSpTb/hAhcW5y/inhGSHE4BO5OwAFkv8PLHFTFhQhIi2mhd4/ToKhtwUbOZh
y1dJcAta2svrycG64L6Bguh7OkYvWljBUd7XAcZfZa6vXjRQlb9DbmmK71+q9MJ8
kYJ2490YmhPPYms2C64tMfxICQECgYEA8MOVm98OHTxnJz5qte1NMdR/I9RRf7PD
8qzj52ZmsR5yGh3IlRFyx1Y6dfVGkWQVxlG5OnKXL6INlwtb0OIGIGvT3kMNMDlU
hNiclwoQUEdlD3P9fMtipO469nw5kSYQxC3wvc63E/ztxvHP+FtiAiRSfIyLWTM0
D0Te7O0l4qsCgYEApgI0GxMsjwA+nTynuIjzQuYcqBhzKi8/9vh9fmyZghXtbM3S
BSlLM3DzM6gHefXuU70/004jcpf/tWha/2kH9Bv9ERSBus/MBGSc2tvgqLgt6xPF
2X/UkeG7ibQFK1QVbXjVEyQhcVOjp8/iKVczEUom1KhI6UVGTDTdMz/jGwECgYEA
hgn9Usf9zZ0BOMHxGtPANEu/dK0RqmzkXEiQoRVLerQehheqwgLyybNh3KXu4aa8
0KpS2w1MykIIGt1CAqqzCn29eHIP95cTTNpjY1tA9dCpnM7QgxegFX5j6TIDwqFU
mEOTUbiyCDi6EBYz2GrXx6V9HsYIFmMBSrbm/TSR8P0CgYB805+/t8g6xHuYTMBy
/8ySfiTCCQZZZF7rOguE2LuNuEzpL8XIGLRAZOYNbyZxpCmqespdrDNWc7xyELeQ
0eodfmeSdhpu6i0JINMa7lhsTY5FqiA8QhXwl4MVcna3RrX5VozyDxOy01zQcwyA
cUIQVkMJ07k2QLgp/QD13KwP+g==
-----END PRIVATE KEY-----`

	unexpectedPrivateKeyType = `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQg9atcufpRkP6IaGbk
DpcN1rUqd0CGVeBgBAGcxyoCv22hRANCAATuNGxyHfzbOhLFsiUUoALGxSSDdQfg
z3njwNAi0eSVi8kAXwwm3dYBDe5bK8LcwfuriTJk8uTpbaBEz3WFadiu
-----END PRIVATE KEY-----`

	invalidPrivateKey = `-----BEGIN PRIVATE KEY-----
aIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQg9atcufpRkP6IaGbk
DpcN1rUqd0CGVeBgBAGcxyoCv22hRANCAATuNGxyHfzbOhLFsiUUoALGxSSDdQfg
z3njwNAi0eSVi8kAXwwm3dYBDe5bK8LcwfuriTJk8uTpbaBEz3WFadiu
-----END PRIVATE KEY-----`

	validPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoXropu9yZG+FJS15/8pW
cU0SYkFETUFASTxzsxhKur+RquaayHQ7LFkIHcqlJowBNxEPx3bplSktwyzmjxXN
knUKh/GFLp1p8Whb2eY5YE0WiPLHmkfIoouXg1j2LcycBjyeSzO0TAFPMWgZlSaU
fPcCP7UHOVVJR/yj/ynYkswXazgS4/oNn4EJgJccb84pzKdR/nRR/FOKTlZZvryi
ek0Xn5Wjyaewl0gd0iZ0lLzvVDgATTyygr6ki3pMSgS4DteMn12tG/rW9eXmUfeV
eBLlhoVbZ51wzkgJCCH+9Q9aomopCyfpoS4teDWmnT4KcLwfwMvRsjxok+d1e2Rp
+QIDAQAB
-----END PUBLIC KEY-----`

	unexpectedPublicKeyType = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEoE/hID8xi255N5fJ8Lc0Pdh2gwJE
b4Lev/UG3mJuXZ5CaPvr7YTEDJ2nxjsGMiV8HalsAMwFxi7mOlZmjnuxPA==
-----END PUBLIC KEY-----`

	invalidPublicKey = `-----BEGIN PUBLIC KEY-----
aFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEoE/hID8xi255N5fJ8Lc0Pdh2gwJE
b4Lev/UG3mJuXZ5CaPvr7YTEDJ2nxjsGMiV8HalsAMwFxi7mOlZmjnuxPA==
-----END PUBLIC KEY-----`

	unexpectedBlockType = `-----BEGIN SOMETHING-----
MIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDoXXH6CbQ1T+e1
y+Iex+R2CuITLTx1fkW+F7CnBzOrCEfiRuNcT3sNFqTnSeTRzxhBnGZzMC+X1L90
FU0GVCD95n2lvPLga2zMcKa5qU2BJFz83g59jmIBfHHVsyikci/VV1/0kOMYPEOm
7WttFgMsqgBiC1UaemGTmJXMhs4F7o1gonouFE9RVj1X27ZfnbhwLjDZEQJvp2k6
OzQNY8YRTTlPjBSs69FBJlWUhGjt3zmMJbVXYojDEm5mPzyPFVxiKBHNDBuNU8xj
79il3IrYv7YXfNldQzdBLBbK7ENfqu9t/PkJaOveBvgQLRPo4GX+3t+M9JuJJd2B
osZWNeWrAgMBAAECggEBAJxLjoiy0kYx0xeTZityJRfJRjvD57DYGK0+XhJbY8Od
NEzdhbznsUsiehUgvQrrE9O+EaNVPA4SihzY3xBssixWRxmeOHf/ihURiPPFD17Y
SLvF2VVW2lFJlYA6nBHQxJ/pv59PfZElqBO2CtY7QjNevhc0rC+9NbkDn28NFbMi
dE9woL+RAvoRWsPhEhZ/aLjI+KAZIqmpaRXODxO1MFAORnv6bvGKbrgMXHAO3hAQ
z4EU6uhZ/pVkaRaowmIvQ3Rv9+MSFzi9TbBjPLR8vYfSPeTgY2399v62XyOWWYOc
9TTE3A+q7iVxMqRSKkg06UBgCzwqXFUVqk4iw8cF1AECgYEA9xHLxE1dXGZID9T2
ROr1pSpTb/hAhcW5y/inhGSHE4BO5OwAFkv8PLHFTFhQhIi2mhd4/ToKhtwUbOZh
y1dJcAta2svrycG64L6Bguh7OkYvWljBUd7XAcZfZa6vXjRQlb9DbmmK71+q9MJ8
kYJ2490YmhPPYms2C64tMfxICQECgYEA8MOVm98OHTxnJz5qte1NMdR/I9RRf7PD
8qzj52ZmsR5yGh3IlRFyx1Y6dfVGkWQVxlG5OnKXL6INlwtb0OIGIGvT3kMNMDlU
hNiclwoQUEdlD3P9fMtipO469nw5kSYQxC3wvc63E/ztxvHP+FtiAiRSfIyLWTM0
D0Te7O0l4qsCgYEApgI0GxMsjwA+nTynuIjzQuYcqBhzKi8/9vh9fmyZghXtbM3S
BSlLM3DzM6gHefXuU70/004jcpf/tWha/2kH9Bv9ERSBus/MBGSc2tvgqLgt6xPF
2X/UkeG7ibQFK1QVbXjVEyQhcVOjp8/iKVczEUom1KhI6UVGTDTdMz/jGwECgYEA
hgn9Usf9zZ0BOMHxGtPANEu/dK0RqmzkXEiQoRVLerQehheqwgLyybNh3KXu4aa8
0KpS2w1MykIIGt1CAqqzCn29eHIP95cTTNpjY1tA9dCpnM7QgxegFX5j6TIDwqFU
mEOTUbiyCDi6EBYz2GrXx6V9HsYIFmMBSrbm/TSR8P0CgYB805+/t8g6xHuYTMBy
/8ySfiTCCQZZZF7rOguE2LuNuEzpL8XIGLRAZOYNbyZxpCmqespdrDNWc7xyELeQ
0eodfmeSdhpu6i0JINMa7lhsTY5FqiA8QhXwl4MVcna3RrX5VozyDxOy01zQcwyA
cUIQVkMJ07k2QLgp/QD13KwP+g==
-----END SOMETHING-----`
)

type RSASuite struct {
	// intentionally empty
}

func TestRSASuite(t *testing.T) {
	fixturez.RunSuite(t, &RSASuite{})
}

func (*RSASuite) TestPEMToRSAPrivateKey(g *WithT) {
	g.Expect(cryptoz.PEMToRSAPrivateKey([]byte(validPrivateKey))).ToNot(BeNil())
	g.Expect(cryptoz.PEMToRSAPrivateKey([]byte("\n\n" + validPrivateKey + "\n\n\n" + unexpectedBlockType + "\n\n"))).ToNot(BeNil())
	g.Expect(cryptoz.PEMToRSAPrivateKey([]byte(unexpectedBlockType + "\n" + validPrivateKey))).ToNot(BeNil())

	g.Expect(cryptoz.PEMToRSAPrivateKey([]byte("bad"))).Error().
		To(MatchError("invalid PEM file or RSA private key not found"))

	g.Expect(cryptoz.PEMToRSAPrivateKey([]byte(unexpectedBlockType))).Error().
		To(MatchError("invalid PEM file or RSA private key not found"))

	g.Expect(cryptoz.PEMToRSAPrivateKey([]byte(invalidPrivateKey))).Error().
		To(MatchError("invalid PEM file or RSA private key not found"))

	g.Expect(cryptoz.PEMToRSAPrivateKey([]byte(unexpectedPrivateKeyType))).Error().
		To(MatchError("invalid PEM file or RSA private key not found"))
}

func (*RSASuite) TestMustPEMToRSAPrivateKey(g *WithT) {
	g.Expect(func() {
		g.Expect(cryptoz.MustPEMToRSAPrivateKey([]byte(validPrivateKey))).ToNot(BeNil())
	}).ToNot(Panic())

	g.Expect(func() {
		cryptoz.MustPEMToRSAPrivateKey([]byte("bad"))
	}).To(PanicWith(MatchError("invalid PEM file or RSA private key not found")))
}

func (*RSASuite) TestPEMToRSAPublicKey(g *WithT) {
	g.Expect(cryptoz.PEMToRSAPublicKey([]byte(validPublicKey))).NotTo(BeNil())
	g.Expect(cryptoz.PEMToRSAPublicKey([]byte("\n\n" + validPublicKey + "\n\n\n" + unexpectedBlockType + "\n\n"))).NotTo(BeNil())
	g.Expect(cryptoz.PEMToRSAPublicKey([]byte(unexpectedBlockType + "\n" + validPublicKey))).NotTo(BeNil())

	g.Expect(cryptoz.PEMToRSAPublicKey([]byte("bad"))).Error().
		To(MatchError("invalid PEM file or RSA public key not found"))

	g.Expect(cryptoz.PEMToRSAPublicKey([]byte(unexpectedBlockType))).Error().
		To(MatchError("invalid PEM file or RSA public key not found"))

	g.Expect(cryptoz.PEMToRSAPublicKey([]byte(invalidPublicKey))).Error().
		To(MatchError("invalid PEM file or RSA public key not found"))

	g.Expect(cryptoz.PEMToRSAPublicKey([]byte(unexpectedPublicKeyType))).Error().
		To(MatchError("invalid PEM file or RSA public key not found"))
}

func (*RSASuite) TestMustPEMToRSAPublicKey(g *WithT) {
	g.Expect(func() {
		g.Expect(cryptoz.MustPEMToRSAPublicKey([]byte(validPublicKey))).ToNot(BeNil())
	}).ToNot(Panic())

	g.Expect(func() {
		cryptoz.MustPEMToRSAPublicKey([]byte("bad"))
	}).To(PanicWith(MatchError("invalid PEM file or RSA public key not found")))
}

func (*RSASuite) TestRSAPrivateKeyToPEM(g *WithT) {
	key, err := cryptoz.PEMToRSAPrivateKey([]byte(validPrivateKey))
	g.Expect(err).To(Succeed())
	g.Expect(key).ToNot(BeNil())
	g.Expect(cryptoz.RSAPrivateKeyToPEM(key)).To(Equal([]byte(validPrivateKey + "\n")))
}

func (*RSASuite) TestRSAPublicKeyToPEM(g *WithT) {
	key, err := cryptoz.PEMToRSAPublicKey([]byte(validPublicKey))
	g.Expect(err).To(Succeed())
	g.Expect(key).ToNot(BeNil())
	g.Expect(cryptoz.RSAPublicKeyToPEM(key)).To(Equal([]byte(validPublicKey + "\n")))
}
