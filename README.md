
# single-url-proxy

Proxy all requests to a single url (supplied as a config). Written to run on Heroku and act as a simple way to wrap an IP camera stream in HTTPS.

Based on the Heroku sample Go app.

Written by Thomas Parslow of [Buzzshot](https://buzzshot.co). Create your online Escape Rooms using [Telescape Live](https://telescapelive.com)

Deploy to herokku now by hitting the button below!

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) version 1.12 or newer and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ git clone https://github.com/almost/single-url-proxy.git
$ cd single-url-proxy
$ go build -o bin/single-url-proxy -v .
$ export URL=http://google.com
$ heroku local
```

Your app should now be running on [localhost:5000](http://localhost:5000/).

## Deploying to Heroku

```sh
$ heroku create
$ heroku config:set URL=http://google.com
$ git push heroku main
$ heroku open
```

or use the deploy button above.
