package testutil

import (
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/stretchr/testify/suite"

	"github.com/0glabs/0g-chain/app"
	"github.com/0glabs/0g-chain/chaincfg"
	"github.com/0glabs/0g-chain/x/dasigners/v1/keeper"
	"github.com/0glabs/0g-chain/x/dasigners/v1/types"
)

// Suite implements a test suite for the module integration tests
type Suite struct {
	suite.Suite

	Keeper        keeper.Keeper
	StakingKeeper *stakingkeeper.Keeper
	App           app.TestApp
	Ctx           sdk.Context
	QueryClient   types.QueryClient
	Addresses     []sdk.AccAddress
}

// SetupTest instantiates a new app, keepers, and sets suite state
func (suite *Suite) SetupTest() {
	chaincfg.SetSDKConfig()
	suite.App = app.NewTestApp()
	suite.Keeper = suite.App.GetDASignersKeeper()
	suite.StakingKeeper = suite.App.GetStakingKeeper()
	suite.Ctx = suite.App.NewContext(true, tmproto.Header{})
	_, accAddresses := app.GeneratePrivKeyAddressPairs(10)
	suite.Addresses = accAddresses

	// Set query client
	queryHelper := suite.App.NewQueryServerTestHelper(suite.Ctx)
	queryHandler := suite.Keeper
	types.RegisterQueryServer(queryHelper, queryHandler)
	suite.QueryClient = types.NewQueryClient(queryHelper)
}
