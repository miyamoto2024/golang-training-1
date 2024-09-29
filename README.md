## Docker setup
- docker build

``` bash
docker build -t golang-training \
  --build-arg GITHUB_USER_NAME=user.name \
  --build-arg GITHUB_USER_EMAIL=user.email \
  --build-arg GITHUB_USER_TOKEN=user.password \
  --build-arg AWS_ACCESS_KEY_ID=user_access_key \
  --build-arg AWS_SECRET_ACCESS_KEY=user_secret_key \
  --no-cache .
```

- docker run

``` bash
docker run -itd --privileged --rm \
  --name golang-training \
  --hostname golang-training \
  golang-training
```

- docker exec
``` bash
docker exec -it golang-training /bin/bash
```

## AWS setup
- line bot access key configure
``` bash
aws secretsmanager create-secret \
--name linebot-access-key \
--description "for access line bot" \
--secret-string file://linebot-access-key.json
```

- lambda build

``` bash
sam build
```

- lambda deploy

``` bash
sam deploy --stack-name sam-app --resolve-s3
```

## Usage

### requirements
* Register LINE Developers and set LINE Messaging API.
* Register AWS Account.
* Input the line access key-value in kine-bot-access-key.json.
