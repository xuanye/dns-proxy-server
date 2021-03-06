package flags

import (
	"flag"
	"os"
	"fmt"
	"github.com/mageddo/dns-proxy-server/cache/store"
)

const TEST_MODE = "TEST_MODE"

var (
	version = "dev" // will be populated by the compiler when generate the release or by this program reading VERSION file
	Cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	Compress = flag.Bool("compress", false, "compress replies")
	Tsig = flag.String("tsig", "", "use MD5 hmac tsig: keyname:base64")
	WebServerPort = flag.Int("web-server-port", 5380, "The web server port")
	DnsServerPort = flag.Int("server-port", 53, "The DNS server to start into")
	SetupResolvconf = flag.Bool("default-dns", true, "This DNS server will be the default server for this machine")
	ConfPath = flag.String("conf-path", "conf/config.json", "The config file path ")
	SetupService = flag.String("service", "", `Setup as service, starting with machine at boot
		docker = start as docker service,
		normal = start as normal service,
		uninstall = uninstall the service from machine `)
	publishServicePort = flag.Bool("service-publish-web-port", true, "Publish web port when running as service in docker mode")
	Version = flag.Bool("version", false, "Current version")
	Help = flag.Bool("help", false, "This message")
)

func init(){

	flag.Parse()
	if *Help {
		flag.PrintDefaults()
		os.Exit(0)
	} else if *Version {
		fmt.Println(GetRawCurrentVersion())
		os.Exit(0)
	}

}

func PublishServicePort() bool {
	return *publishServicePort
}

func GetRawCurrentVersion() string {

	//if len(version) == 0 {
		//b, err := ioutil.ReadFile(utils.GetPath("VERSION")) // just pass the file name
		//if err == nil {
		//	return string(b)
		//}
		//log.Logger.Warningf("status=could-not-recover-version, err=%v", err)
	//}
	return version
}

func IsTestVersion() bool {
	cache := store.GetInstance()
	if !cache.ContainsKey(TEST_MODE){
		cache.PutIfAbsent(TEST_MODE, flag.Lookup("test.v") != nil)
	}
	return cache.Get(TEST_MODE).(bool)
}
