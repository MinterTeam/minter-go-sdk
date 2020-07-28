package transaction

import (
	"testing"
)

func TestTransactionRedeemCheck_Sign(t *testing.T) {
	data := NewRedeemCheckData().
		MustSetProof("da021d4f84728e0d3d312a18ec84c21768e0caa12a53cb0a1452771f72b0d1a91770ae139fd6c23bcf8cec50f5f2e733eabb8482cf29ee540e56c6639aac469600").
		MustSetRawCheck("Mcf89b01830f423f8a4d4e5400000000000000843b9aca00b8419b3beac2c6ad88a8bd54d24912754bb820e58345731cb1b9bc0885ee74f9e50a58a80aa990a29c98b05541b266af99d3825bb1e5ed4e540c6e2f7c9b40af9ecc011ca00f7ba6d0aa47d74274b960fba02be03158d0374b978dcaa5f56fc7cf1754f821a019a829a3b7bba2fc290f5c96e469851a3876376d6a6a4df937327b3a5e9e8297")

	tx, err := NewBuilder(TestNetChainID).NewTransaction(data)
	if err != nil {
		t.Fatal(err)
	}

	nonce := uint64(1)
	gasPrice := uint8(1)

	transaction := tx.SetNonce(nonce).SetGasPrice(gasPrice).SetGasCoin(1)

	signedTx, err := transaction.Sign("05ddcd4e6f7d248ed1388f0091fe345bf9bf4fc2390384e26005e7675c98b3c1")
	if err != nil {
		t.Fatal(err)
	}

	validSignature := "0xf901350102010109b8e4f8e2b89df89b01830f423f8a4d4e5400000000000000843b9aca00b8419b3beac2c6ad88a8bd54d24912754bb820e58345731cb1b9bc0885ee74f9e50a58a80aa990a29c98b05541b266af99d3825bb1e5ed4e540c6e2f7c9b40af9ecc011ca00f7ba6d0aa47d74274b960fba02be03158d0374b978dcaa5f56fc7cf1754f821a019a829a3b7bba2fc290f5c96e469851a3876376d6a6a4df937327b3a5e9e8297b841da021d4f84728e0d3d312a18ec84c21768e0caa12a53cb0a1452771f72b0d1a91770ae139fd6c23bcf8cec50f5f2e733eabb8482cf29ee540e56c6639aac469600808001b845f8431ca0a761d80557ce73dae5eda3834d879e3585d3e7b08fc58015cd7ddfa1e8ca57f9a052d6844fcd66eef411e66e47425f0ee8a223616bfba63930831bcfad2fb84c52"
	bytes, err := signedTx.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != validSignature {
		t.Errorf("EncodeTx got %s, want %s", string(bytes), validSignature)
	}
}
