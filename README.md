# go.nc0.fr

A Cloudflare Workers server to provide vanity URLs for my Go modules
on `go.nc0.fr`.

If you want a version made in Go (CF does not support Go directly), check
Google Cloud's [Vanity URL service](https://github.com/GoogleCloudPlatform/govanityurl).

Configuration is done inside the [VANITIES.ts](VANITIES.ts) file.

Update the worker using Wrangler:

```bash
$ npm run deploy
```

## Contributing

I do not expect any contributions but we never know /shrug/.

## License

Th project is governed under a BSD-style license that can be found in the
[LICENSE](LICENSE) file.
