# CookieCutter template - SAM GO Lambda

## Install golang

<a href="https://golang.org/dl/">Official Installation guidelines</a>

## Install sam

<a href="https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html">Official Installation Guidelines</a>

or

```text
brew upgrade
brew update
brew tap aws/tap
brew install aws-sam-cli
```

## Create new project

```text
sam init -l  https://github.com/kdevar/sam-lambda-go.git -o $GOPATH/src/<repo-location>/
#prompt for project_name
#prompt for repo_location (e.g. github.com/kdevar)

```

For example

```text
sam init -l https://github.com/kdevar/sam-lambda-go.git -o $HOME/go/src/github.com/kdevar
# project_name: my-test
# repo_location: github.com/kdevar
```

You only need to specify the full git url the first time you use the template.  Subsequent times can be done simply as ```sam-lambda-go```.

```text
sam init -l sam-lambda-go -o $HOME/go/src/github.com/kdevar
...
```

Unless of course you want the latest changes to the template...in which case use the full git url

## Open project

```text
cd $HOME/go/src/github.com/kdevar/my-test
code .
```

## Running the project

Follow instructions in the readme of the newly created project