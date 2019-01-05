# {{cookiecutter.project_name|replace("_", " ")|title|replace(" ", "")}}

Service generated from the sam-lambda-go template

## Start Local

1. First time only to setup go modules
```
make init
```
2. Run local server
```
make serve
```

## Deploy to sandbox

1. Paste your aws credentials in the terminal
2. Run command
```
make deploy_sandbox
```

## Deploy to environments

<a href="https://basket1.atlassian.net/wiki/spaces/EN/pages/82870283/How+to+Deploy+your+Project" target="_blank">Instructions to Deploy</a>