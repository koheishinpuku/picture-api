module github.com/unimal-jp/benkyoukai-go

go 1.14

replace local.packages/DB => ./DB

replace local.packages/Controllers => ./Controllers

replace local.packages/Models => ./Models

replace local.packages/Redd => ./goredis

replace local.packages/Midd => ./middleware

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/djimenez/iconv-go v0.0.0-20160305225143-8960e66bd3da // indirect
	github.com/go-redis/redis/v8 v8.1.3
	github.com/jinzhu/gorm v1.9.16
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.1.17
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/lib/pq v1.1.1
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/olahol/go-imageupload v0.0.0-20160503070439-09d2b92fa05e // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	local.packages/Controllers v0.0.0-00010101000000-000000000000
	local.packages/DB v0.0.0-00010101000000-000000000000
	local.packages/Midd v0.0.0-00010101000000-000000000000
	local.packages/Models v0.0.0-00010101000000-000000000000 // indirect
	local.packages/Redd v0.0.0-00010101000000-000000000000

)
