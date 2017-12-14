package wavingocean

import (
	"bytes"
	"context"

	"github.com/xiaokangwang/waVingOcean/configure"
	"github.com/yinghuocho/gotun2socks"
	"github.com/yinghuocho/gotun2socks/tun"
	"v2ray.com/core"
)

/*Ignite Start Tap server from configure
 */
func Ignite(cfg configure.WaVingOceanConfigure) {
	//Start V2Ray
	configure, err := core.LoadConfig(core.ConfigFormat_Protobuf, bytes.NewBuffer(cfg.V2RayConfigure))
	if err != nil {
		panic(err)
	}
	ns, err := newSimpleServer(configure)
	if err != nil {
		panic(err)
	}
	err = ns.Start()
	if err != nil {
		panic(err)
	}

	//Start Tap
	f, e := tun.OpenTunDevice(cfg.Tun.Name, cfg.Tun.Address, cfg.Tun.Gateway, cfg.Tun.Mask, cfg.DNSServers)
	if e != nil {
		panic(e)
	}
	//Start Tun2Socks
	ctx := context.Background()
	tunc := gotun2socks.New(f, &V2Dialer{ser: ns}, cfg.DNSServers, cfg.PublicOnly, cfg.EnableDnsCache, ctx)
	tunc.Run()
}
