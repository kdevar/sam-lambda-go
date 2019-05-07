# {{cookiecutter.project_name|replace("_", " ")|title|replace(" ", "")}}

Service generated from the sam-lambda-go template

## Start Local

```
make serve
```

You should find your application at <a href="http://127.0.0.1:3000/healthcheck">http://127.0.0.1:3000/healthcheck</a>

## Deploy to aws sandbox

1. Paste your aws credentials in the terminal
2. Run command

```
make deploy_sandbox
```
