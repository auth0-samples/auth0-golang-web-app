module 01-Login/routes/callback

go 1.12

require (
	app v0.0.0
	auth v0.0.0
	github.com/coreos/go-oidc v2.1.0+incompatible
)

replace app => ../../app

replace auth => ../../auth
