# twitterReader
twitter reader assistant written in Go lang

uses festival for the text to speech features


to execute
```
./twitterReader
```

install festival
```
sudo apt-get install festival festvox-kallpc16k
```
install spanish
```
sudo apt-get install festvox-ellpc11k
```
install catalan voice
```
sudo add-apt-repository ppa:zeehio/festcat
sudo apt-get update
sudo apt-get install festival-ca festvox-ca-ona-hts festcat-utils
```

needs the twitter account tokens in the file twitterConfig.json
```json
{
    "username": "account_name",
    "consumer_key": "xxxxxxxxxxxxxx",
    "consumer_secret": "xxxxxxxxxxx",
    "access_token_key": "xxxxxxxxxx",
    "access_token_secret": "xxxxxxx"
}

```
