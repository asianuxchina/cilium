<!-- This file was autogenerated via cilium-dbg cmdref, do not edit manually-->

## cilium-dbg bpf multicast subscriber add

Add a remote subscriber to the multicast group.

### Synopsis

Add a remote subscriber to the multicast group.
Only remote subscribers are added via command line. Local subscribers will be automatically populated in the map based on IGMP messages.
Remote subscribers are typically other Cilium nodes, identified by internal IP of the node.
To add remote subscriber, following information is required:
- group: multicast group address to which subscriber is added.
- subscriber-address: subscriber IP address.


```
cilium-dbg bpf multicast subscriber add <group> <subscriber-address> [flags]
```

### Examples

```
To add a remote node 10.100.0.1 to multicast group 229.0.0.1, use the following command:
cilium-dbg bpf multicast subscriber add 229.0.0.1 10.100.0.1

```

### Options

```
  -h, --help   help for add
```

### Options inherited from parent commands

```
      --config string        Config file (default is $HOME/.cilium.yaml)
  -D, --debug                Enable debug messages
  -H, --host string          URI to server-side API
      --log-driver strings   Logging endpoints to use (example: syslog)
      --log-opt map          Log driver options (example: format=json)
```

### SEE ALSO

* [cilium-dbg bpf multicast subscriber](cilium-dbg_bpf_multicast_subscriber.md)	 - Manage the multicast subscribers.

