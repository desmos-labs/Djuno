package profiles

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/simapp/params"

	sdk "github.com/cosmos/cosmos-sdk/types"
	desmosdb "github.com/desmos-labs/djuno/database"
	"github.com/desmos-labs/juno/modules"
	juno "github.com/desmos-labs/juno/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.MessageModule = &Module{}
)

type Module struct {
	encodingConfig *params.EncodingConfig
	db             *desmosdb.DesmosDb
}

func NewModule(encodingConfig *params.EncodingConfig, db *desmosdb.DesmosDb) *Module {
	return &Module{
		encodingConfig: encodingConfig,
		db:             db,
	}
}

// Name implements Module
func (m *Module) Name() string {
	return "profiles"
}

func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	return HandleGenesis(m.encodingConfig.Amino, appState, m.db)
}

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	return HandleMsg(tx, index, msg, m.db)
}
