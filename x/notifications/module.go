package notifications

import (
	"github.com/desmos-labs/juno/modules"
	juno "github.com/desmos-labs/juno/types"
)

var (
	_ modules.Module            = &Module{}
	_ modules.TransactionModule = &Module{}
)

// Module represents the module to send notifications to the user
type Module struct{}

// NewModule builds a new Module instance
func NewModule() *Module {
	return &Module{}
}

// Name implements modules.Module
func (m Module) Name() string {
	return "notifications"
}

// HandleTx implements modules.TransactionModule
func (m *Module) HandleTx(tx *juno.Tx) error {
	return TxHandler(tx)
}
