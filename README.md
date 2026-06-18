# Config Reloader Agent

Kubernetes ConfigMap/Secret change detector with auto-reload.

## How It Works
1. Watches ConfigMap/Secret changes via K8s API
2. Triggers rolling restart on affected Deployments
3. Supports annotation-based opt-in/opt-out
4. Cooldown to prevent restart storms

## Annotations
```yaml
metadata:
  annotations:
    config-reloader/reload: "true"
    config-reloader/cooldown: "30s"
```

## License
MIT
