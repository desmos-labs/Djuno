package x

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	desmosdb "github.com/desmos-labs/djuno/database"
	"github.com/desmos-labs/djuno/x/bank"
	"github.com/desmos-labs/djuno/x/notifications"
	"github.com/desmos-labs/djuno/x/posts"
	"github.com/desmos-labs/djuno/x/profiles"
	"github.com/desmos-labs/djuno/x/relationships"
	"github.com/desmos-labs/djuno/x/reports"
	"github.com/desmos-labs/juno/client"
	"github.com/desmos-labs/juno/db"
	"github.com/desmos-labs/juno/modules"
	"github.com/desmos-labs/juno/modules/registrar"
	juno "github.com/desmos-labs/juno/types"
)

var (
	_ registrar.Registrar = &ModulesRegistrar{}
)

// ModulesRegistrar represents the registrar.Registrar that allows to register all custom DJuno modules
type ModulesRegistrar struct{}

// NewModulesRegistrar allows to build a new ModulesRegistrar instance
func NewModulesRegistrar() *ModulesRegistrar {
	return &ModulesRegistrar{}
}

// BuildModules implements registrar.Registrar
func (r *ModulesRegistrar) BuildModules(
	_ *juno.Config, encodingConfig *params.EncodingConfig, _ *sdk.Config, db db.Database, _ *client.Proxy,
) modules.Modules {
	desmosDb := desmosdb.Cast(db)
	return []modules.Module{
		bank.NewModule(desmosDb),
		notifications.NewModule(),
		posts.NewModule(encodingConfig, desmosDb),
		profiles.NewModule(encodingConfig, desmosDb),
		relationships.NewModule(desmosDb),
		reports.NewModule(desmosDb),
	}
}
