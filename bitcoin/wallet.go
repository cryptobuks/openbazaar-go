package bitcoin

import (
	"github.com/OpenBazaar/spvwallet"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	btc "github.com/btcsuite/btcutil"
	hd "github.com/btcsuite/btcutil/hdkeychain"
)

type BitcoinWallet interface {

	// Start the wallet
	Start()

	// Return the network parameters
	Params() *chaincfg.Params

	// Returns the type of crytocurrency this wallet implements
	CurrencyCode() string

	// Get the master private key
	MasterPrivateKey() *hd.ExtendedKey

	// Get the master public key
	MasterPublicKey() *hd.ExtendedKey

	// Get the current address for the given purpose
	CurrentAddress(purpose spvwallet.KeyPurpose) btc.Address

	// Returns if the wallet has the key for the given address
	HasKey(addr btc.Address) bool

	// Get the confirmed and unconfirmed balances
	Balance() (confirmed, unconfirmed int64)

	// Returns a list of transactions for this wallet
	Transactions() ([]spvwallet.Txn, error)

	// Get the height of the blockchain
	ChainTip() uint32

	// Get the current fee per byte
	GetFeePerByte(feeLevel spvwallet.FeeLevel) uint64

	// Send bitcoins to an external wallet
	Spend(amount int64, addr btc.Address, feeLevel spvwallet.FeeLevel) (*chainhash.Hash, error)

	// Calculates the estimated size of the transaction and returns the total fee for the given feePerByte
	EstimateFee(ins []spvwallet.TransactionInput, outs []spvwallet.TransactionOutput, feePerByte uint64) uint64

	// Build and broadcast a transaction that sweeps all coins from a 1 of 2 multisig to an internal address (optionally provided)
	SweepMultisig(utxos []spvwallet.Utxo, address *btc.Address, key *hd.ExtendedKey, reddemScript []byte, feeLevel spvwallet.FeeLevel) error

	// Create a signature for a multisig transaction
	CreateMultisigSignature(ins []spvwallet.TransactionInput, outs []spvwallet.TransactionOutput, key *hd.ExtendedKey, redeemScript []byte, feePerByte uint64) ([]spvwallet.Signature, error)

	// Combine signatures and broadcast
	Multisign(ins []spvwallet.TransactionInput, outs []spvwallet.TransactionOutput, sigs1 []spvwallet.Signature, sigs2 []spvwallet.Signature, redeemScript []byte, feePerByte uint64) error

	// Generate a multisig script from public keys
	GenerateMultisigScript(keys []hd.ExtendedKey, threshold int) (addr btc.Address, redeemScript []byte, err error)

	// Add a script to the wallet and get notifications back when coins are received or spent from it
	AddWatchedScript(script []byte) error

	// Add a callback for incoming transactions
	AddTransactionListener(func(spvwallet.TransactionCallback))

	// Use this to re-download merkle blocks in case of missed transactions
	ReSyncBlockchain(fromHeight int32)

	// Cleanly disconnect from the wallet
	Close()
}
