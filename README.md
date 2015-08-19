# etcd-leader

Simple way to find the etcd leader.

## Usage

Set the ETCDCTL_PEERS env var:

```
export ETCDCTL_PEERS=http://etcd0.example.com:2379,http://etcd1.example.com:2379,http://etcd2.example.com:2379
```

Run the etcd-leader command:

```
./etcd-leader
http://etcd1.example.com:2379
```
