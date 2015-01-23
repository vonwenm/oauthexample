# OAuth 2.0 example

## Install and run

To install this example application, run:

`go get go get github.com/apita1973/oauthexample`

This app requires `cert.pem` and `key.pem` files (for TLS/https). Generate the files by running the following from the app root:

`go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"`

You must obtain obtain OAuth 2.0 credentials from the [Google Developers Console][google]. 
Make sure to select "web application" as the client type when creating the client id.
Also make sure to enter the following for the redirect url:

`https://localhost:8001/oauth2callback`

Your app's client id and secret must be entered in the conf.gcfg file. See the conf.gcfg.example file for format.

Simply access `https://localhost:8001` to launch.

## OAuth 2.0 Resources

* [http://tutorials.jenkov.com/oauth2/index.html](http://tutorials.jenkov.com/oauth2/index.html)
* [http://tools.ietf.org/html/draft-ietf-oauth-v2-23](http://tools.ietf.org/html/draft-ietf-oauth-v2-23)
* [https://developers.google.com/accounts/docs/OAuth2WebServer](https://developers.google.com/accounts/docs/OAuth2WebServer)
* [https://developers.google.com/oauthplayground/](https://developers.google.com/oauthplayground/)	

## OAuth 2.0 Security Concerns

* [http://homakov.blogspot.com/2013/03/oauth1-oauth2-oauth.html](http://homakov.blogspot.com/2013/03/oauth1-oauth2-oauth.html)

## License

The [MIT license][mit].

[google]: https://developers.google.com/accounts/docs/OAuth2
[mit]: http://opensource.org/licenses/MIT
