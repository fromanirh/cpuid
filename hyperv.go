package cpuid

const (
	HypervVendorSignature    string = "Microsoft Hv"
	HypervInterfaceSignature string = "Hv#1"
)

func HasHypervSupport() {
	return HypervVendorIDSignatureString == HypervVendorSignature && HypervInterfaceSignatureString == HypervInterfaceSignature
}

var HypervSupport bool

var HypervCPUIDMax uint32

var HypervVendorIDSignatureString string

var HypervInterfaceSignatureString string
