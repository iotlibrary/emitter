/**********************************************************************************
* Copyright (c) 2009-2017 Misakai Ltd.
* This program is free software: you can redistribute it and/or modify it under the
* terms of the GNU Affero General Public License as published by the  Free Software
* Foundation, either version 3 of the License, or(at your option) any later version.
*
* This program is distributed  in the hope that it  will be useful, but WITHOUT ANY
* WARRANTY;  without even  the implied warranty of MERCHANTABILITY or FITNESS FOR A
* PARTICULAR PURPOSE.  See the GNU Affero General Public License  for  more details.
*
* You should have  received a copy  of the  GNU Affero General Public License along
* with this program. If not, see<http://www.gnu.org/licenses/>.
************************************************************************************/

package security

import (
	"errors"
	//"github.com/emitter-io/emitter/collection"
	//"github.com/emitter-io/emitter/network/http"
)

// Contract represents an interface for a contract.
type Contract interface {
	Validate(key Key) bool // Validate checks the security key with the contract.
}

// contract represents a contract (user account).
type contract struct {
	ID        int32  `json:"id"`     // Gets or sets the contract id.
	MasterID  uint16 `json:"sign"`   // Gets or sets the master id.
	Signature int32  `json:"master"` // Gets or sets the signature of the contract.
}

// Validate validates the contract data against a key.
func (c *contract) Validate(key Key) bool {
	return c.MasterID == key.Master() &&
		c.Signature == key.Signature() &&
		c.ID == key.Contract()
}

// ContractProvider represents an interface for a contract provider.
type ContractProvider interface {
	// Creates a new instance of a Contract in the underlying data storage.
	Create() (Contract, error)
	Get(id int32) Contract
}

// SingleContractProvider provides contracts on premise.
type SingleContractProvider struct {
	owner *contract
}

// NewSingleContractProvider creates a new single contract provider.
func NewSingleContractProvider(license *License) *SingleContractProvider {
	p := new(SingleContractProvider)
	p.owner = new(contract)
	p.owner.MasterID = 1
	p.owner.ID = license.Contract
	p.owner.Signature = license.Signature
	return p
}

// Create creates a contract, the SingleContractProvider way.
func (p *SingleContractProvider) Create() (Contract, error) {
	return nil, errors.New("Single contract provider can not create contracts")
}

// Get returns a ContractData fetched by its id.
func (p *SingleContractProvider) Get(id int32) Contract {
	if p.owner == nil || p.owner.ID != id {
		return nil
	}
	return p.owner
}

/*
// HTTPContractProvider provides contracts over http.
type HTTPContractProvider struct {
	owner *contract
	cache *collection.ConcurrentMap
}

// NewHTTPContractProvider creates a new single contract provider.
func NewHTTPContractProvider(license *License) *HTTPContractProvider {
	p := new(HTTPContractProvider)
	p.owner = new(contract)
	p.owner.MasterID = 1
	p.owner.ID = license.Contract
	p.owner.Signature = license.Signature
	p.cache = collection.NewConcurrentMap()
	return p
}

// Create creates a contract, the HTTPContractProvider way.
func (p *HTTPContractProvider) Create() (Contract, error) {
	return nil, errors.New("HTTP contract provider can not create contracts")
}

// Get returns a ContractData fetched by its id.
// TODO : transform id in uint32 everywhere.
func (p *HTTPContractProvider) Get(id int32) Contract {
	//contract, ok := p.cache.Get(uint32(id))
	//if !ok {

	//}
	return nil
	//return Contract(contract)
}

func (p *HTTPContractProvider) fetchContract(id int32) *Contract {
	c := new(Contract)
	err := http.Get("http://meta.emitter.io/v1/contract/", c)
	if err != nil {
		return nil
	}
	return c
}
*/
