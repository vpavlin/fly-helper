# fly-helper

A minimalistic tool helping with deployment to Fly.io

## Config

This tool supports configuration in YAML or JSON


```
---
secrets:
  input:
  - name: my important secret file
    path: "/tmp/my-file.txt"
  - name: client-cert
    path: "./client.crt"
  output:
  - name: my important secret file
    path: "/mnt/something.txt"
  - name: client-cert
    path: "/opt/client.crt"
```

```
{
    "secrets": {
        "input": [
            {
                "name": "my important secret file",
                "path": "/tmp/my-file.txt"
            },
            {
                "name": "client-cert",
                "path": "./client.crt"
            }
        ],
        "output": [
            {
                "name": "my important secret file",
                "path": "/mnt/something.txt"
            },
            {
                "name": "client-cert",
                "path": "/opt/client.crt"
            }
        ]
    }
}
```

Configuration can be provided as a file (`--config`) or as base64 encoded environment variable (`--config-env`).

## Secrets

Allows you to provide list of files which are uploaded as `fly secrets` using the following command

```
flyhelper secrets push
```

You can tehn run `flyhelper` in the Fly App container to turn the secrets provided in environment variables into files again.

```
flyhelper secrets pull
```

The `Name` values are used to produce secrets names e.g. `my important secret file` will turn into `FLY_SECRET_MY_IMPORTANT_SECRET_FILE`. The content of the file on `Path` will be base64 encoded an used as the value of the secret.

The tool will also upload the `config.json` into the secrets as `FLY_HELPER_CONFG_ENV`. This is then used by the `flyhelper` inside the app image to export the secrets to filesystem.