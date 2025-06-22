# Netron
It is a system for self-organizing a decentralized virtual network that uses blockchain technology, and solves the problem of EdgeVPN, which causes exponential memory and storage pressure as time passes and the number of participants in the network increases. By solving this problem, even devices with extremely limited resources can stably participate in a decentralized network, making it possible to explosively expand the sustainable network size. This means that a variety of layered, bubble-like network spaces can self-organize on top of the network infrastructure known as the Internet, with countless devices of all sizes acting like neurons in this complex network. So I call it Netron, from the words network and neuron.

# Usage

Netron works by generating tokens (or a configuration file) that can be shared between different machines, hosts or peers to access to a decentralized secured network between them.

Every token is unique and identifies the network,  no central server setup, or specifying hosts ip is required.

To generate a config run:

```bash
# Generate a new config file and use it later as NETRONCONFIG
$ netron -g > config.yaml
```

OR to generate a portable token:

```bash
$ NETRONTOKEN=$(netron -g -b)
```

Note, tokens are config merely encoded in base64, so this is equivalent:

```bash
$ NETRONTOKEN=$(netron -g | tee config.yaml | base64 -w0)
```

All netron commands implies that you either specify a `NETRONTOKEN` (or `--token` as parameter) or a `NETRONCONFIG` as this is the way for `netron` to establish a network between the nodes.

The configuration file is the network definition and allows you to connect over to your peers securely.

**Warning** Exposing this file or passing-it by is equivalent to give full control to the network.

## As a VPN

To start the VPN, simply run `netron` without any argument.

An example of running netron on multiple hosts:

```bash
# on Node A
$ NETRONTOKEN=.. netron --address 10.1.0.11/24
# on Node B
$ NETRONTOKEN=.. netron --address 10.1.0.12/24
# on Node C ...
$ NETRONTOKEN=.. netron --address 10.1.0.13/24
...
```

... and that's it! the `--address` is a _virtual_ unique IP for each node, and it is actually the ip where the node will be reachable to from the vpn. You can assign IPs freely to the nodes of the network, while you can override the default `netron0` interface with `IFACE` (or `--interface`)

*Note*: It might take up time to build the connection between nodes. Wait at least 5 mins, it depends on the network behind the hosts.

## Example use case: network-decentralized [k3s](https://github.com/k3s-io/k3s) test cluster

Let's see a practical example, you are developing something for kubernetes and you want to try a multi-node setup, but you have machines available that are only behind NAT (pity!) and you would really like to leverage HW.

If you are not really interested in network performance (again, that's for development purposes only!) then you could use `netron` + [k3s](https://github.com/k3s-io/k3s) in this way:

1) Generate netron config: `netron -g > vpn.yaml`
2) Start the vpn:

   on node A: `sudo IFACE=netron0 ADDRESS=10.1.0.3/24 NETRONCONFIG=vpn.yml netron`

   on node B: `sudo IFACE=netron0 ADDRESS=10.1.0.4/24 NETRONCONFIG=vpm.yml netron`
3) Start k3s:

   on node A: `k3s server --flannel-iface=netron0`

   on node B: `K3S_URL=https://10.1.0.3:6443 K3S_TOKEN=xx k3s agent --flannel-iface=netron0 --node-ip 10.1.0.4`

We have used flannel here, but other CNI should work as well.


## As a library

Netron can be used as a library. It is very portable and offers a functional interface.

To join a node in a network from a token, without starting the vpn:

```golang

import (
    node "github.com/mudler/netron/pkg/node"
)

e := node.New(
    node.Logger(l),
    node.LogLevel(log.LevelInfo),
    node.MaxMessageSize(2 << 20),
    node.FromBase64( mDNSEnabled, DHTEnabled, token ),
    // ....
  )

e.Start(ctx)

```

or to start a VPN:

```golang

import (
    vpn "github.com/mudler/netron/pkg/vpn"
    node "github.com/mudler/netron/pkg/node"
)

opts, err := vpn.Register(vpnOpts...)
if err != nil {
	return err
}

e := netron.New(append(o, opts...)...)

e.Start(ctx)
```

# Credits

- Revised [EdgeVPN](https://github.com/mudler/edgevpn/tree/master).
- The awesome [libp2p](https://github.com/libp2p) library
- [https://github.com/songgao/water](https://github.com/songgao/water) for tun/tap devices in golang
- [Room example](https://github.com/libp2p/go-libp2p/tree/master/examples/chat-with-rendezvous) (shamelessly parts are copied by)
- Logo originally made by [Uniconlabs](https://www.flaticon.com/authors/uniconlabs) from [www.flaticon.com](https://www.flaticon.com/).

# Troubleshooting

If during bootstrap you see messages like:

```
netron[3679]:             * [/ip4/104.131.131.82/tcp/4001] failed to negotiate stream multiplexer: context deadline exceeded
```

or

```
netron[9971]: 2021/12/16 20:56:34 failed to sufficiently increase receive buffer size (was: 208 kiB, wanted: 2048 kiB, got: 416 kiB). See https://github.com/lucas-clemente/quic-go/wiki/UDP-Receive-Buffer-Size for details.
```

or generally experiencing poor network performance, it is recommended to increase the maximum buffer size by running:

```
sysctl -w net.core.rmem_max=2500000
```
