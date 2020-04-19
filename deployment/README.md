# Authenticate Docker on the server
https://cloud.google.com/container-registry/docs/advanced-authentication#json-key

Service account required with *Cloud Storage/admin* permissions.

Authenticate Docker on production: 

```
cat keyfile.json | docker login -u _json_key --password-stdin https://gcr.io
```