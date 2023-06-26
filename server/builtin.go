package server

import (
	"github.com/RavenspiritMedia/Indra-chain/consensus"
	consensusDev "github.com/RavenspiritMedia/Indra-chain/consensus/dev"
	consensusDummy "github.com/RavenspiritMedia/Indra-chain/consensus/dummy"
	consensusIBFT "github.com/RavenspiritMedia/Indra-chain/consensus/ibft"
	"github.com/RavenspiritMedia/Indra-chain/secrets"
	"github.com/RavenspiritMedia/Indra-chain/secrets/awsssm"
	"github.com/RavenspiritMedia/Indra-chain/secrets/gcpssm"
	"github.com/RavenspiritMedia/Indra-chain/secrets/hashicorpvault"
	"github.com/RavenspiritMedia/Indra-chain/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
