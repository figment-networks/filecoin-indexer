package pipeline

import (
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"

	"github.com/figment-networks/filecoin-indexer/model"
	"github.com/figment-networks/indexing-engine/pipeline"
)

var (
	_ pipeline.PayloadFactory = (*PayloadFactory)(nil)
	_ pipeline.Payload        = (*payload)(nil)
)

// NewPayloadFactory creates a payload factory
func NewPayloadFactory() *PayloadFactory {
	return &PayloadFactory{}
}

// PayloadFactory creates payloads
type PayloadFactory struct{}

// GetPayload returns a payload for a given height
func (pf *PayloadFactory) GetPayload(height int64) pipeline.Payload {
	return &payload{
		startedAt:     time.Now(),
		currentHeight: height,
	}
}

type payload struct {
	startedAt     time.Time
	currentHeight int64
	processed     bool

	// Fetcher stage
	EpochTipset          *types.TipSet
	DealsData            map[string]api.MarketDeal
	DealsCount           map[address.Address]uint32
	DealsSlashedCount    map[address.Address]uint32
	DealsSlashedIDs      []string
	MinersAddresses      []address.Address
	MinersInfo           []*miner.MinerInfo
	MinersPower          []*api.MinerPower
	MinersFaults         []*bitfield.BitField
	TransactionsCIDs     []cid.Cid
	TransactionsMessages []*types.Message

	// Parser stage
	Epoch        *model.Epoch
	Miners       []*model.Miner
	Transactions []*model.Transaction
	Events       []*model.Event

	// Sequencer stage
	StoredMiners         map[string]model.Miner
	StoredSlashedDealIDs []string
}

func (p *payload) SetCurrentHeight(height int64) {
	p.currentHeight = height
}

func (p *payload) GetCurrentHeight() int64 {
	return p.currentHeight
}

func (p *payload) MarkAsProcessed() {
	p.processed = true
}

func (p *payload) IsProcessed() bool {
	return p.processed
}

func (p *payload) Duration() float64 {
	return time.Since(p.startedAt).Seconds()
}