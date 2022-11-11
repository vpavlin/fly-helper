# fly-helper

A minimalistic tool helping with deployment to Fly.io

## Config

```
{
    "AppName": "some-app",
    "Secrets": {
        "Input": [
            {
                "Name": "my important secret file",
                "Path": "/tmp/my-file.txt"
            },
            {
                "Name": "client-cert",
                "Path": "./client.crt"
            }
        ],
        "Output": [
            {
                "Name": "my important secret file",
                "Path": "/mnt/something.txt"
            },
            {
                "Name": "client-cert",
                "Path": "/opt/client.crt"
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

