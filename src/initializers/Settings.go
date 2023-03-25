package initializers

import "os"

var SecretKey = []byte(os.Getenv("secret"))

// JWTExpirationTime How long the expiration time is in minutes
var JWTExpirationTime = 1
