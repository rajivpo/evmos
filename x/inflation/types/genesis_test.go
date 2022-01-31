package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type GenesisTestSuite struct {
	suite.Suite
}

func (suite *GenesisTestSuite) SetupTest() {
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

func (suite *GenesisTestSuite) TestValidateGenesis() {
	// Team Address needs to be set manually at Genesis
	validParams := DefaultParams()

	newGen := NewGenesisState(validParams, uint64(0), "day", 365, sdk.OneDec())

	testCases := []struct {
		name     string
		genState *GenesisState
		expPass  bool
	}{
		{
			"empty genesis",
			&GenesisState{},
			false,
		},
		{
			"invalid default genesis",
			DefaultGenesisState(),
			true,
		},
		{
			"valid genesis constructor",
			&newGen,
			true,
		},
		{
			"valid genesis",
			&GenesisState{
				Params:          validParams,
				Period:          uint64(5),
				EpochIdentifier: "day",
				EpochsPerPeriod: 365,
				BondedRatio:     sdk.OneDec(),
			},
			true,
		},
		{
			"invalid genesis",
			&GenesisState{
				Params: validParams,
			},
			false,
		},
		{
			"invalid genesis - empty eporchIdentifier",
			&GenesisState{
				Params:          validParams,
				Period:          uint64(5),
				EpochIdentifier: "",
				EpochsPerPeriod: 365,
				BondedRatio:     sdk.OneDec(),
			},
			false,
		},
		{
			"invalid genesis - zero epochsperPerid",
			&GenesisState{
				Params:          validParams,
				Period:          uint64(5),
				EpochIdentifier: "day",
				EpochsPerPeriod: 0,
				BondedRatio:     sdk.OneDec(),
			},
			false,
		},
		{
			"invalid genesis - negative bondedRatio",
			&GenesisState{
				Params:          validParams,
				Period:          uint64(5),
				EpochIdentifier: "day",
				EpochsPerPeriod: 365,
				BondedRatio:     sdk.OneDec().Neg(),
			},
			false,
		},
		{
			"invalid genesis - greater than 1 bondedRatio",
			&GenesisState{
				Params:          validParams,
				Period:          uint64(5),
				EpochIdentifier: "day",
				EpochsPerPeriod: 365,
				BondedRatio:     sdk.NewDecWithPrec(12, 1),
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		err := tc.genState.Validate()
		if tc.expPass {
			suite.Require().NoError(err, tc.name)
		} else {
			suite.Require().Error(err, tc.name)
		}
	}
}