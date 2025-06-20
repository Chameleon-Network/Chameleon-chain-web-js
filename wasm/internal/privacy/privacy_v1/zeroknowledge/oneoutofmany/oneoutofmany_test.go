package oneoutofmany

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"chameleon-chain/common"
	"chameleon-chain/privacy/operation"
	"chameleon-chain/privacy/privacy_util"
	"chameleon-chain/privacy/privacy_v1/zeroknowledge/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	m.Run()
}

var _ = func() (_ struct{}) {
	fmt.Println("This runs before init()!")
	Logger.Init(common.NewBackend(nil).Logger("test", true))
	return
}()

//TestPKOneOfMany test protocol for one of many Commitment is Commitment to zero
func TestPKOneOfMany(t *testing.T) {
	// prepare witness for Out out of many protocol
	for i := 0; i < 10; i++ {
		witness := new(OneOutOfManyWitness)

		//indexIsZero := int(common.RandInt() % operation.CommitmentRingSize)
		indexIsZero := 0

		// list of commitments
		commitments := make([]*operation.Point, privacy_util.CommitmentRingSize)
		values := make([]*operation.Scalar, privacy_util.CommitmentRingSize)
		randoms := make([]*operation.Scalar, privacy_util.CommitmentRingSize)

		for i := 0; i < privacy_util.CommitmentRingSize; i++ {
			values[i] = operation.RandomScalar()
			randoms[i] = operation.RandomScalar()
			commitments[i] = operation.PedCom.CommitAtIndex(values[i], randoms[i], operation.PedersenSndIndex)
		}

		// create Commitment to zero at indexIsZero
		values[indexIsZero] = new(operation.Scalar).FromUint64(0)
		commitments[indexIsZero] = operation.PedCom.CommitAtIndex(values[indexIsZero], randoms[indexIsZero], operation.PedersenSndIndex)

		witness.Set(commitments, randoms[indexIsZero], uint64(indexIsZero))
		start := time.Now()
		proof, err := witness.Prove()
		assert.Equal(t, nil, err)
		end := time.Since(start)
		//fmt.Printf("One out of many proving time: %v\n", end)

		//fmt.Printf("Proof: %v\n", proof)

		// validate sanity for proof
		isValidSanity := proof.ValidateSanity()
		assert.Equal(t, true, isValidSanity)

		// verify the proof
		start = time.Now()
		res, err := proof.Verify()
		end = time.Since(start)
		fmt.Printf("One out of many verification time: %v\n", end)
		assert.Equal(t, true, res)
		assert.Equal(t, nil, err)

		//Convert proof to bytes array
		proofBytes := proof.Bytes()
		assert.Equal(t, utils.OneOfManyProofSize, len(proofBytes))

		// revert bytes array to proof
		proof2 := new(OneOutOfManyProof).Init()
		err = proof2.SetBytes(proofBytes)
		assert.Equal(t, nil, err)
		proof2.Statement.Commitments = commitments
		assert.Equal(t, proof, proof2)

		// verify the proof
		start = time.Now()
		res, err = proof2.Verify()
		end = time.Since(start)
		fmt.Printf("One out of many verification time: %v\n", end)
		assert.Equal(t, true, res)
		assert.Equal(t, nil, err)

	}
}
