### credentials

you must set your guru credentials either by setting the enviroment variables

```
GURU_EMAIL={{ email }}
GURU_PASS={{ password }}
```

or by saving them in json format at $HOME/.guru/credentials

```
{
  "email": "your email",
  "password": "your password"
}
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
./guru-cli some tags
```

create a card

```
./guru-cli create-card < title > < everything else will be the content >
```

tag a card

```
./guru-cli add-tags < card id > some useful tags
```

create a tag

```
./guru-cli create-tag < tag >
```

get a card's contents

```
./guru-cli get-card < card id >
```
