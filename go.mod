module github.com/polynetwork/explorer

go 1.14

require (
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cmars/basen v0.0.0-20150613233007-fe3947df716e // indirect
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/ethereum/go-ethereum v1.9.15
	github.com/go-redis/redis v6.9.2+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/mock v1.4.0 // indirect
	github.com/joeqian10/neo-gogogo v0.0.0
	github.com/ontio/ontology v1.11.0
	github.com/ontio/ontology-go-sdk v1.11.1
	github.com/pkg/errors v0.9.1
	github.com/polynetwork/poly v0.0.0-20200818035458-8083385c9933
	github.com/polynetwork/poly-go-sdk v0.0.0-20200817120957-365691ad3493
	github.com/shopspring/decimal v1.2.0
	github.com/tendermint/tendermint v0.33.7
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/Switcheo/cosmos-sdk v0.39.2-0.20200814061308-474a0dbbe4ba
	github.com/ethereum/go-ethereum => github.com/ethereum/go-ethereum v1.9.13
	github.com/joeqian10/neo-gogogo => github.com/blockchain-develop/neo-gogogo v0.0.0-20200824102609-fddf06a45f66
)
