# Return Client IP

config example

```yml
# Static configuration

experimental:
  plugins:
    ReturnClientIP:
        moduleName: github.com/moonlightwatch/ReturnClientIP
        version: v0.1.0

```

```yml
# Dynamic configuration

http:
  routers:
    my-router:
      rule: host(`demo.localhost`)
      service: service-foo
      entryPoints:
        - web
      middlewares:
        - my-plugin

  services:
   service-foo:
      loadBalancer:
        servers:
          - url: http://127.0.0.1:5000
  
  middlewares:
    my-plugin:
      plugin:
        ReturnClientIP:
          Active: true
```