package profiles

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/profile"
	desmosdb "github.com/desmos-labs/djuno/database"
	"github.com/desmos-labs/juno/parse/worker"
	tmtypes "github.com/tendermint/tendermint/types"
)

// GenesisHandler allows to properly handle the genesis state for the posts module
func GenesisHandler(
	codec *codec.Codec, _ *tmtypes.GenesisDoc, appState map[string]json.RawMessage, w worker.Worker,
) error {
	db, ok := w.Db.(desmosdb.DesmosDb)
	if !ok {
		return fmt.Errorf("database is not a DesmosDB instance")
	}

	var genState profile.GenesisState
	codec.MustUnmarshalJSON(appState[profile.ModuleName], &genState)

	// Save the profiles
	for _, prof := range genState.Profiles {
		if _, err := db.UpsertProfile(prof); err != nil {
			return err
		}
	}

	return nil
}