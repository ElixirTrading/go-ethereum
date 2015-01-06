/*
	This file is part of go-ethereum

	go-ethereum is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	go-ethereum is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with go-ethereum.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"crypto/elliptic"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/logger"
	"github.com/ethereum/go-ethereum/p2p"
)

func main() {
	logger.AddLogSystem(logger.NewStdLogSystem(os.Stdout, log.LstdFlags, logger.InfoLevel))
	key, _ := crypto.GenerateKey()
	marshaled := elliptic.Marshal(crypto.S256(), key.PublicKey.X, key.PublicKey.Y)

	srv := p2p.Server{
		MaxPeers:   100,
		Identity:   p2p.NewSimpleClientIdentity("Ethereum(G)", "0.1", "Peer Server Two", marshaled),
		ListenAddr: ":30301",
		NAT:        p2p.UPNP(),
		NoDial:     true,
	}
	if err := srv.Start(); err != nil {
		log.Fatal("could not start server:", err)
	}
	select {}
}
