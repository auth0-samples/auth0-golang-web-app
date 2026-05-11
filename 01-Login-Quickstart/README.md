# Auth0 + Go Web App Quickstart

This sample demonstrates how to add authentication to a Go web app using Auth0 and the [go-auth0](https://github.com/auth0/go-auth0) SDK.

Check the [Go Quickstart](https://auth0.com/docs/quickstart/webapp/golang) for a full step-by-step guide.

## Running the App

Make sure you have **Go 1.25+** installed.

Copy the `.env.example` file to `.env` and provide your Auth0 credentials:

```bash
AUTH0_DOMAIN=YOUR_AUTH0_DOMAIN
AUTH0_CLIENT_ID=YOUR_AUTH0_CLIENT_ID
AUTH0_CLIENT_SECRET=YOUR_AUTH0_CLIENT_SECRET
AUTH0_CALLBACK_URL=http://localhost:3000/callback
SESSION_SECRET=a-long-random-secret-key-for-cookie-encryption
```

Then run:

```bash
go run .
```

Navigate to [http://localhost:3000/](http://localhost:3000/).

## Running with Docker

```bash
cp .env.example .env
# Edit .env with your Auth0 credentials
./exec.sh
```

## What is Auth0?

Auth0 helps you to:

* Add authentication with [multiple authentication sources](https://auth0.com/docs/authenticate/identity-providers), either social like Google, Facebook, Microsoft Account, LinkedIn, GitHub, Twitter, and others, or enterprise identity systems like Windows Azure AD, Google Apps, Active Directory, ADFS, or any SAML Identity Provider.
* Add authentication through more traditional [username/password databases](https://auth0.com/docs/authenticate/database-connections/custom-db/create-db-connection).
* Add support for [linking different user accounts](https://auth0.com/docs/manage-users/user-accounts/user-account-linking/link-user-accounts) with the same user.
* Support for generating signed [JSON Web Tokens](https://auth0.com/docs/secure/tokens/json-web-tokens) to call your APIs and flow the user identity securely.
* Analytics of how, when, and where users are logging in.
* Pull data from other sources and add it to the user profile through [Actions](https://auth0.com/docs/customize/actions).

## Create a free Auth0 account

1. Go to [Auth0](https://auth0.com/signup) and click Sign Up.
2. Use Google, GitHub, or Microsoft Account to login.

## Issue Reporting

If you have found a bug or if you have a feature request, please report them at this repository issues section. Please do not report security vulnerabilities on the public GitHub issue tracker. The [Responsible Disclosure Program](https://auth0.com/whitehat) details the procedure for disclosing security issues.

## Author

[Auth0](https://auth0.com)

## License

This project is licensed under the MIT license. See the [LICENSE](../LICENSE) file for more info.
