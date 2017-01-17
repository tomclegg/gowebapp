# gowebapp

A basic skeleton web application server. Just add HTML, client-side JavaScript code, and server-side APIs.

Strategy:
* In development, use npm to install JavaScript libraries.
* At build time, use webpack on nodejs to compile JavaScript assets.
* Deploy with a single Go binary -- no nodejs, no asset files.

## dev/build dependencies

Install Go and nodejs.

## run dev-mode server

In dev mode, source maps are served, and JS is not minified.

```sh
npm run dev
```

## build production-mode server

```sh
npm build
```

## build & run production-mode server

```sh
npm start
```
