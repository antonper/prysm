package blockchain

import (
	"context"
	"testing"

	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	blockchainTesting "github.com/prysmaticlabs/prysm/beacon-chain/blockchain/testing"
	testDB "github.com/prysmaticlabs/prysm/beacon-chain/db/testing"
	"github.com/prysmaticlabs/prysm/beacon-chain/forkchoice/protoarray"
	"github.com/prysmaticlabs/prysm/beacon-chain/operations/attestations"
	"github.com/prysmaticlabs/prysm/beacon-chain/state/stateutil"
	"github.com/prysmaticlabs/prysm/shared/testutil"
)

func TestService_ReceiveBlockNoPubsub(t *testing.T) {
	ctx := context.Background()

	genesis, keys := testutil.DeterministicGenesisState(t, 64)
	genFullBlock := func(t *testing.T, conf *testutil.BlockGenConfig, slot uint64) *ethpb.SignedBeaconBlock {
		blk, err := testutil.GenerateFullBlock(genesis, keys, conf, slot)
		if err != nil {
			t.Error(err)
		}
		return blk
	}

	type args struct {
		block *ethpb.SignedBeaconBlock
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		check   func(*testing.T, *Service)
	}{
		{
			name: "applies block with state transition",
			args: args{
				block: genFullBlock(t, testutil.DefaultBlockGenConfig(), 2 /*slot*/),
			},
			check: func(t *testing.T, s *Service) {
				if hs := s.head.state.Slot(); hs != 2 {
					t.Errorf("Unexpected state slot. Got %d but wanted %d", hs, 2)
				}
				if bs := s.head.block.Block.Slot; bs != 2 {
					t.Errorf("Unexpected head block slot. Got %d but wanted %d", bs, 2)
				}
			},
		},
		{
			name: "saves attestations to pool",
			args: args{
				block: genFullBlock(t,
					&testutil.BlockGenConfig{
						NumProposerSlashings: 0,
						NumAttesterSlashings: 0,
						NumAttestations:      2,
						NumDeposits:          0,
						NumVoluntaryExits:    0,
					},
					1, /*slot*/
				),
			},
			check: func(t *testing.T, s *Service) {
				if baCount := len(s.attPool.BlockAttestations()); baCount != 2 {
					t.Errorf("Did not get the correct number of block attestations saved to the pool. "+
						"Got %d but wanted %d", baCount, 2)
				}
			},
		},
		/*
			{
				name: "updates exit pool",
			},
			{
				name: "sets epochParticipation",
			},
			{
				name: "notifies block processed on state feed",
			},
		*/
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := testDB.SetupDB(t)
			defer testDB.Cleanup(t, db)
			gBlkRoot := [32]byte{'g', 'e', 'n', 'e', 's', 'i', 's'}
			if err := db.SaveState(ctx, genesis, gBlkRoot); err != nil {
				t.Fatal(err)
			}

			cfg := &Config{
				BeaconDB: db,
				ForkChoiceStore: protoarray.New(
					0, // justifiedEpoch
					0, // finalizedEpoch
					gBlkRoot,
				),
				AttPool:       attestations.NewPool(),
				StateNotifier: &blockchainTesting.MockStateNotifier{},
			}
			s, err := NewService(ctx, cfg)
			if err != nil {
				t.Fatal(err)
			}
			if err := s.saveGenesisData(ctx, genesis); err != nil {
				t.Fatal(err)
			}
			root, err := stateutil.BlockRoot(tt.args.block.Block)
			if err != nil {
				t.Error(err)
			}
			if err := s.ReceiveBlockNoPubsub(ctx, tt.args.block, root); (err != nil) != tt.wantErr {
				t.Errorf("ReceiveBlockNoPubsub() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				tt.check(t, s)
			}
		})
	}
}
