package conf

import (
	. "github.com/mageddo/dns-proxy-server/log"
	"fmt"
)

type setMachineDNSServerHandler struct {
	serverIP string
}

func (hd *setMachineDNSServerHandler) process(line string, entryType DnsEntry) *string {

	switch entryType {
	case PROXY:
		LOGGER.Infof("status=found-dns-proxy-entry")
		v := getDNSLine(hd.serverIP)
		return &v
	case SERVER:
		v := fmt.Sprintf("# %s # dps-comment", line)
		return &v
	default:
		return &line
	}
}

func (hd *setMachineDNSServerHandler) afterProcess(hasContent bool, foundDnsProxy bool) *string {
	if !hasContent || !foundDnsProxy {
		v := getDNSLine(hd.serverIP)
		return &v
	}
	return nil
}

func newSetMachineDnsServerHandler(serverIP string) DnsHandler {
	hd := setMachineDNSServerHandler{}
	hd.serverIP = serverIP;
	return &hd
}
