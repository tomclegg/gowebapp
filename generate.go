//go:generate sh -c "which go-bindata 2>&1 >/dev/null || go get github.com/jteeuwen/go-bindata/..."
//go:generate sh -c "which go-bindata-assetfs 2>&1 >/dev/null || go get github.com/elazarl/go-bindata-assetfs/..."
//go:generate sh -c "[ -d node_modules ] || npm install"
//go:generate sh -c "rm -r bindata.tmp && mkdir bindata.tmp"
//go:generate npm run webpack
//go:generate sh -c "cp -rp static/* bindata.tmp/"
//go:generate go-bindata-assetfs bindata.tmp/...

package main
