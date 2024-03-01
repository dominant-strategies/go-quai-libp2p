package types

import (
	"math/big"

	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/common/hexutil"
	"github.com/dominant-strategies/go-quai/log"
	"google.golang.org/protobuf/proto"
	"lukechampine.com/blake3"
)

type WorkObject struct {
	woHeader WorkObjectHeader
	woBlock  *WorkObjectBody
	tx       Transaction
}

type WorkObjectHeader struct {
	headerHash common.Hash
	parentHash common.Hash
	number     *big.Int
	difficulty *big.Int
	txHash     common.Hash
	location   common.Location
	nonce      BlockNonce
}

type WorkObjectBody struct {
	header          *Header
	transactions    Transactions
	extTransactions Transactions
	uncles          []*Header
	manifest        BlockManifest
}

func (wo *WorkObject) Header() *Header {
	return wo.woBlock.header
}

func (wo *WorkObject) Transactions() Transactions {
	return wo.woBlock.transactions
}

func (wo *WorkObject) ExtTransactions() Transactions {
	return wo.woBlock.extTransactions
}

func (wo *WorkObject) Uncles() []*Header {
	return wo.woBlock.uncles
}

func (wo *WorkObject) Manifest() BlockManifest {
	return wo.woBlock.manifest
}

func (wo *WorkObject) ParentHash() common.Hash {
	return wo.woHeader.parentHash
}

func (wo *WorkObject) Number() *big.Int {
	return wo.woHeader.number
}

func (wo *WorkObject) Difficulty() *big.Int {
	return wo.woHeader.difficulty
}

func (wo *WorkObject) TxHash() common.Hash {
	return wo.woHeader.txHash
}

func (wo *WorkObject) Nonce() BlockNonce {
	return wo.woHeader.nonce
}

func (wo *WorkObject) HeaderHash() common.Hash {
	return wo.woHeader.headerHash
}

func (wo *WorkObject) Tx() Transaction {
	return wo.tx
}

func (wo *WorkObject) SetTx(tx Transaction) {
	wo.tx = tx
}

func (wo *WorkObject) SetHeader(header *Header) {
	wo.woBlock.header = header
}

func (wo *WorkObject) SetTransactions(transactions Transactions) {
	wo.woBlock.transactions = transactions
}

func (wo *WorkObject) SetExtTransactions(transactions Transactions) {
	wo.woBlock.extTransactions = transactions
}

func (wo *WorkObject) SetUncles(uncles []*Header) {
	wo.woBlock.uncles = uncles
}

func (wo *WorkObject) SetManifest(manifest BlockManifest) {
	wo.woBlock.manifest = manifest
}

func (wo *WorkObject) SetParentHash(parentHash common.Hash) {
	wo.woHeader.parentHash = parentHash
}

func (wo *WorkObject) SetNumber(number *big.Int) {
	wo.woHeader.number = number
}

func (wo *WorkObject) SetDifficulty(difficulty *big.Int) {
	wo.woHeader.difficulty = difficulty
}

func (wo *WorkObject) SetTxHash(txHash common.Hash) {
	wo.woHeader.txHash = txHash
}

func (wo *WorkObject) SetNonce(nonce BlockNonce) {
	wo.woHeader.nonce = nonce
}

func (wo *WorkObject) SetHeaderHash(headerHash common.Hash) {
	wo.woHeader.headerHash = headerHash
}

func NewWorkObject(woHeader *WorkObjectHeader, woBody *WorkObjectBody, tx Transaction) *WorkObject {
	return &WorkObject{
		woHeader: *woHeader,
		woBlock:  woBody,
		tx:       tx,
	}
}

func (wh *WorkObjectHeader) SetHeaderHash(headerHash common.Hash) {
	wh.headerHash = headerHash
}

func (wh *WorkObjectHeader) SetParentHash(parentHash common.Hash) {
	wh.parentHash = parentHash
}

func (wh *WorkObjectHeader) SetNumber(number *big.Int) {
	wh.number = number
}

func (wh *WorkObjectHeader) SetDifficulty(difficulty *big.Int) {
	wh.difficulty = difficulty
}

func (wh *WorkObjectHeader) SetTxHash(txHash common.Hash) {
	wh.txHash = txHash
}

func (wh *WorkObjectHeader) SetNonce(nonce BlockNonce) {
	wh.nonce = nonce
}

func (wh *WorkObjectHeader) SetLocation(location common.Location) {
	wh.location = location
}

func (wh *WorkObjectHeader) HeaderHash() common.Hash {
	return wh.headerHash
}

func (wh *WorkObjectHeader) ParentHash() common.Hash {
	return wh.parentHash
}

func (wh *WorkObjectHeader) Number() *big.Int {
	return wh.number
}

func (wh *WorkObjectHeader) Difficulty() *big.Int {
	return wh.difficulty
}

func (wh *WorkObjectHeader) TxHash() common.Hash {
	return wh.txHash
}

func (wh *WorkObjectHeader) Nonce() BlockNonce {
	return wh.nonce
}

func (wh *WorkObjectHeader) Location() common.Location {
	return wh.location
}

func NewWorkObjectHeader(headerHash common.Hash, parentHash common.Hash, number *big.Int, difficulty *big.Int, txHash common.Hash, nonce BlockNonce, location common.Location) *WorkObjectHeader {
	return &WorkObjectHeader{
		headerHash: headerHash,
		parentHash: parentHash,
		number:     number,
		difficulty: difficulty,
		txHash:     txHash,
		nonce:      nonce,
		location:   location,
	}
}

func (wh *WorkObjectHeader) RPCMarshalWorkObjectHeader() map[string]interface{} {
	result := map[string]interface{}{
		"headerHash": wh.HeaderHash(),
		"parentHash": wh.ParentHash(),
		"number":     (*hexutil.Big)(wh.Number()),
		"difficulty": (*hexutil.Big)(wh.Difficulty()),
		"nonce":      wh.Nonce(),
		"location":   wh.Location(),
		"txHash":     wh.TxHash(),
	}
	return result
}

func (wh *WorkObjectHeader) Hash() (hash common.Hash) {
	sealHash := wh.SealHash().Bytes()
	hasherMu.Lock()
	defer hasherMu.Unlock()
	hasher.Reset()
	var hData [40]byte
	copy(hData[:], wh.Nonce().Bytes())
	copy(hData[len(wh.nonce):], sealHash)
	sum := blake3.Sum256(hData[:])
	hash.SetBytes(sum[:])
	return hash
}

func (wh *WorkObjectHeader) SealHash() (hash common.Hash) {
	hasherMu.Lock()
	defer hasherMu.Unlock()
	hasher.Reset()
	protoSealData := wh.SealEncode()
	data, err := proto.Marshal(protoSealData)
	if err != nil {
		log.Global.Error("Failed to marshal seal data ", "err", err)
	}
	sum := blake3.Sum256(data[:])
	hash.SetBytes(sum[:])
	return hash
}

func (wh *WorkObjectHeader) SealEncode() *ProtoWorkObjectHeader {
	hash := common.ProtoHash{Value: wh.HeaderHash().Bytes()}
	parentHash := common.ProtoHash{Value: wh.ParentHash().Bytes()}
	txHash := common.ProtoHash{Value: wh.TxHash().Bytes()}
	number := wh.Number().Bytes()
	difficulty := wh.Difficulty().Bytes()
	location := wh.Location().ProtoEncode()

	return &ProtoWorkObjectHeader{
		HeaderHash: &hash,
		ParentHash: &parentHash,
		Number:     number,
		Difficulty: difficulty,
		TxHash:     &txHash,
		Location:   location,
	}
}

func (wh *WorkObjectHeader) ProtoEncode() *ProtoWorkObjectHeader {
	hash := common.ProtoHash{Value: wh.HeaderHash().Bytes()}
	parentHash := common.ProtoHash{Value: wh.ParentHash().Bytes()}
	txHash := common.ProtoHash{Value: wh.TxHash().Bytes()}
	number := wh.Number().Bytes()
	difficulty := wh.Difficulty().Bytes()
	location := wh.Location().ProtoEncode()
	nonce := wh.Nonce().Uint64()

	return &ProtoWorkObjectHeader{
		HeaderHash: &hash,
		ParentHash: &parentHash,
		Number:     number,
		Difficulty: difficulty,
		TxHash:     &txHash,
		Location:   location,
		Nonce:      &nonce,
	}
}

// func (wh *WorkObjectHeader) ProtoDecode(data *ProtoWorkObjectHeader) {
// 	wh.SetHeaderHash(common.BytesToHash(data.HeaderHash.Value))
// 	wh.SetParentHash(common.BytesToHash(data.ParentHash.Value))
// 	wh.SetNumber(new(big.Int).SetBytes(data.Number))
// 	wh.SetDifficulty(new(big.Int).SetBytes(data.Difficulty))
// 	wh.SetTxHash(common.BytesToHash(data.TxHash.Value))
// 	wh.SetNonce(new(big.Int).SetBytes(data.Nonce))
// 	wh.SetLocation(common.ProtoDecodeLocation(data.Location))
// }

func (wb *WorkObjectBody) Header() *Header {
	return wb.header
}

func (wb *WorkObjectBody) Transactions() Transactions {
	return wb.transactions
}

func (wb *WorkObjectBody) ExtTransactions() Transactions {
	return wb.extTransactions
}

func (wb *WorkObjectBody) Uncles() []*Header {
	return wb.uncles
}

func (wb *WorkObjectBody) Manifest() BlockManifest {
	return wb.manifest
}

func (wb *WorkObjectBody) SetHeader(header *Header) {
	wb.header = header
}

func (wb *WorkObjectBody) SetTransactions(transactions Transactions) {
	wb.transactions = transactions
}

func (wb *WorkObjectBody) SetExtTransactions(transactions Transactions) {
	wb.extTransactions = transactions
}

func (wb *WorkObjectBody) SetUncles(uncles []*Header) {
	wb.uncles = uncles
}

func (wb *WorkObjectBody) SetManifest(manifest BlockManifest) {
	wb.manifest = manifest
}

func NewWorkObjectBody(header *Header, txs []*Transaction, etxs []*Transaction, uncles []*Header, subManifest BlockManifest, receipts []*Receipt, hasher TrieHasher, nodeCtx int) *WorkObjectBody {
	wb := &WorkObjectBody{header: CopyHeader(header)}

	// TODO: panic if len(txs) != len(receipts)
	if len(txs) == 0 {
		wb.header.SetTxHash(EmptyRootHash)
	} else {
		wb.header.SetTxHash(DeriveSha(Transactions(txs), hasher))
		wb.transactions = make(Transactions, len(txs))
		copy(wb.transactions, txs)
	}

	if len(receipts) == 0 {
		wb.header.SetReceiptHash(EmptyRootHash)
	} else {
		wb.header.SetReceiptHash(DeriveSha(Receipts(receipts), hasher))
	}

	if len(uncles) == 0 {
		wb.header.SetUncleHash(EmptyUncleHash)
	} else {
		wb.header.SetUncleHash(CalcUncleHash(uncles))
		wb.uncles = make([]*Header, len(uncles))
		for i := range uncles {
			wb.uncles[i] = CopyHeader(uncles[i])
		}
	}

	if len(etxs) == 0 {
		wb.header.SetEtxHash(EmptyRootHash)
	} else {
		wb.header.SetEtxHash(DeriveSha(Transactions(etxs), hasher))
		wb.extTransactions = make(Transactions, len(etxs))
		copy(wb.extTransactions, etxs)
	}

	// Since the subordinate's manifest lives in our body, we still need to check
	// that the manifest matches the subordinate's manifest hash, but we do not set
	// the subordinate's manifest hash.
	subManifestHash := EmptyRootHash
	if len(subManifest) != 0 {
		subManifestHash = DeriveSha(subManifest, hasher)
		wb.manifest = make(BlockManifest, len(subManifest))
		copy(wb.manifest, subManifest)
	}
	if nodeCtx < common.ZONE_CTX && subManifestHash != wb.Header().ManifestHash(nodeCtx+1) {
		log.Global.Error("attempted to build block with invalid subordinate manifest")
		return nil
	}

	return wb
}
