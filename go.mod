module github.com/desmos-labs/djuno

go 1.13

require (
	cloud.google.com/go v0.55.0 // indirect
	cloud.google.com/go/firestore v1.1.1 // indirect
	firebase.google.com/go v3.12.0+incompatible
	github.com/cosmos/cosmos-sdk v0.42.7
	github.com/desmos-labs/desmos v0.17.2
	github.com/desmos-labs/juno v0.0.0-20210715061943-1ad33ccc600c
	github.com/jmoiron/sqlx v1.2.0
	github.com/pelletier/go-toml v1.8.1
	github.com/proullon/ramsql v0.0.0-20181213202341-817cee58a244
	github.com/rs/zerolog v1.21.0
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.11
	github.com/ziutek/mymysql v1.5.4 // indirect
	google.golang.org/api v0.20.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/cosmos/cosmos-sdk => github.com/desmos-labs/cosmos-sdk v0.42.5-0.20210712073217-87acd62da7d7

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2
