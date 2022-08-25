package v1beta1

func SetDefaultShmScoringArgs(args *ShmScoringArgs) {
	if args.AddrPorts == nil {
		addrPorts := make([]string, 0)
		args.AddrPorts = &addrPorts
	}
}
