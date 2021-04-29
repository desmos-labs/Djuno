module github.com/desmos-labs/djuno

go 1.13

require (
	cloud.google.com/go v0.55.0 // indirect
	cloud.google.com/go/firestore v1.1.1 // indirect
	firebase.google.com/go v3.12.0+incompatible
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/desmos-labs/desmos v0.15.5-0.20210326121423-2f2e8049808e
	github.com/desmos-labs/juno v0.0.0-20210429060016-6548afd78e72
	github.com/jmoiron/sqlx v1.2.0
	github.com/proullon/ramsql v0.0.0-20181213202341-817cee58a244
	github.com/rs/zerolog v1.20.0
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.9
	github.com/ziutek/mymysql v1.5.4 // indirect
	golang.org/x/tools v0.0.0-20200321224714-0d839f3cf2ed // indirect
	google.golang.org/api v0.20.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
