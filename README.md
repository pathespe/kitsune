
# local development

Set up a local registry
```
docker run -d -p 5000:5000 --restart=always --name registry registry:2
```

Auto reloading for api
```
skaffold dev --default-repo=localhost:5000
```
