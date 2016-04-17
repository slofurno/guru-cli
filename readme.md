In order to work with guru's api, we need to get a relogin token from the guru webapp,  which is then used to request temporary tokens.

open your browsers dev tools, and then load the guru app at https://app.getguru.com/

look for the request to /auth, checkout the request headers, and find the cookie

you will need at least the following part of the cookie:

```
email=slofurno%40gmail.com; slofurno%40gmail.com_reloginTok={{ some uuid }};
```

make a directory .guru in your $HOME, and paste the above into a file called relogin_token, eg:

```
echo "{{ cookie }}" > $HOME/.guru/relogin_token
```

### install
```
git clone https://github.com/slofurno/guru-cli.git
cd guru-cli
go get
go build
```

### testing

```
go test
```

### cli
search by tag

```
./guru-cli find <comma,delimited,tags>
```

create a card

```
./guru-cli create-card <title> <everything else will be the content>
```
