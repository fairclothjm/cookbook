# nmap

## usage examples

Flag to set timing template (higher is faster):
```
-T<0-5>
```

### OS scan

Determine OS of hosts:
```
nmap -O 192.168.0.0/24
```

### ping scan

Determine hosts reachable via ping:
```
nmap -sn 192.168.0.0/24
```
