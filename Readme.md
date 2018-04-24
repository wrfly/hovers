# hovers

A http proxy over a socks5 proxy.

Just a poor tool like [privoxy](https://www.privoxy.org/)

## usage

```bash
➜ ./hovers -h
Usage of ./hovers:
  -p 		proxy listen port (default "8080")
  -socks5	socks5 proxy (default "127.0.0.1:1080")
  -v		verbose
```

## otherwise

```bash
http_proxy=http://localhost:8080 curl ip.kfd.me

http_proxy=socks5://localhost:1080 curl ip.kfd.me

https_proxy=http://localhost:8080 curl https://ip.kfd.me
```