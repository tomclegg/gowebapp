# gowebapp

A basic skeleton web application server. Just add HTML, client-side JavaScript code, and server-side APIs.

Strategy:
* In development, use npm to install JavaScript libraries.
* At build time, use webpack on nodejs to compile JavaScript assets.
* Deploy with a single Go binary -- no nodejs, no asset files.

## dev/build dependencies

Go:

```
curl https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz \
     | sudo tar -C /usr/local -xzf - \
     && (cd /usr/local/bin && sudo ln -s ../go/bin/* .)
```

nodejs:

```sh
curl -sL https://deb.nodesource.com/setup_6.x | sudo bash -
sudo apt-get install nodejs
```

## run dev-mode server

In dev mode, source maps are served, and JS is not minified.

```sh
npm run dev
```

## build production-mode server

```sh
npm build
```

The server binary will be installed to `$GOPATH/bin/`.

## build & run production-mode server

```sh
npm start
```
