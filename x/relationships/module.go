package relationships

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	desmosdb "github.com/desmos-labs/djuno/database"
	"github.com/desmos-labs/juno/modules"
	juno "github.com/desmos-labs/juno/types"
)

var (
	_ modules.Module        = &Module{}
	_ modules.MessageModule = &Module{}
)

type Module struct {
	db *desmosdb.DesmosDb
}

func NewModule(db *desmosdb.DesmosDb) *Module {
	return &Module{
		db: db,
	}
}

// Name implements Module
func (m *Module) Name() string {
	return "relationships"
}

func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	return HandleMsg(tx, msg, m.db)
}
