package jomtx

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jomluz/jomtx/testutil/sample"
	jomtxsimulation "github.com/jomluz/jomtx/x/jomtx/simulation"
	"github.com/jomluz/jomtx/x/jomtx/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = jomtxsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateTxn = "op_weight_msg_txn"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTxn int = 100

	opWeightMsgUpdateTxn = "op_weight_msg_txn"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTxn int = 100

	opWeightMsgDeleteTxn = "op_weight_msg_txn"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTxn int = 100

	opWeightMsgClaimTxn = "op_weight_msg_claim_txn"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaimTxn int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	jomtxGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		TxnList: []types.Txn{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		TxnCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&jomtxGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTxn int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTxn, &weightMsgCreateTxn, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTxn = defaultWeightMsgCreateTxn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTxn,
		jomtxsimulation.SimulateMsgCreateTxn(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTxn int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTxn, &weightMsgUpdateTxn, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTxn = defaultWeightMsgUpdateTxn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTxn,
		jomtxsimulation.SimulateMsgUpdateTxn(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTxn int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteTxn, &weightMsgDeleteTxn, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTxn = defaultWeightMsgDeleteTxn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTxn,
		jomtxsimulation.SimulateMsgDeleteTxn(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgClaimTxn int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgClaimTxn, &weightMsgClaimTxn, nil,
		func(_ *rand.Rand) {
			weightMsgClaimTxn = defaultWeightMsgClaimTxn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimTxn,
		jomtxsimulation.SimulateMsgClaimTxn(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateTxn,
			defaultWeightMsgCreateTxn,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				jomtxsimulation.SimulateMsgCreateTxn(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateTxn,
			defaultWeightMsgUpdateTxn,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				jomtxsimulation.SimulateMsgUpdateTxn(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteTxn,
			defaultWeightMsgDeleteTxn,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				jomtxsimulation.SimulateMsgDeleteTxn(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgClaimTxn,
			defaultWeightMsgClaimTxn,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				jomtxsimulation.SimulateMsgClaimTxn(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
