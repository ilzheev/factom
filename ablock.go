// Copyright 2016 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factom

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

type AdminID byte

const (
	AIDMinuteNumber             AdminID = iota // 0
	AIDDBSignature                             // 1
	AIDRevealHash                              // 2
	AIDAddHash                                 // 3
	AIDIncreaseServerCount                     // 4
	AIDAddFederatedServer                      // 5
	AIDAddAuditServer                          // 6
	AIDRemoveFederatedServer                   // 7
	AIDAddFederatedServerKey                   // 8
	AIDAddFederatedServerBTCKey                // 9
	AIDServerFault                             // 10
	AIDCoinbaseDescriptor                      // 11
	AIDCoinbaseDescriptorCancel                // 12
	AIDAddAuthorityAddress                     // 13
	AIDAddAuthorityEfficiency                  // 14
)

var (
	ErrAIDUnknown = errors.New("unknown ABlock Entry type")
)

type ABlock struct {
	PrevBackreferenceHash string    `json:"prevbackrefhash"`
	DBHeight              int64     `json:"dbheight"`
	BackReverenceHash     string    `json:"backreferencehash"`
	LookupHash            string    `json:"lookuphash"`
	ABEntries             []ABEntry `json:"abentries"`
}

func (a *ABlock) String() string {
	var s string

	s += fmt.Sprintln("BackReverenceHash:", a.BackReverenceHash)
	s += fmt.Sprintln("LookupHash:", a.LookupHash)
	s += fmt.Sprintln("PrevBackreferenceHash:", a.PrevBackreferenceHash)
	s += fmt.Sprintln("DBHeight:", a.DBHeight)

	s += fmt.Sprintln("ABEntries {")
	for _, v := range a.ABEntries {
		s += fmt.Sprintln(v)
	}
	s += fmt.Sprintln("}")

	return s
}

func (a *ABlock) UnmarshalJSON(js []byte) error {
	tmp := new(struct {
		Header struct {
			PrevBackreferenceHash string `json:"prevbackrefhash"`
			DBHeight              int64  `json:"dbheight"`
		}
		BackReverenceHash string            `json:"backreferencehash"`
		LookupHash        string            `json:"lookuphash"`
		ABEntries         []json.RawMessage `json:"abentries"`
	})

	err := json.Unmarshal(js, tmp)
	if err != nil {
		return err
	}

	a.PrevBackreferenceHash = tmp.Header.PrevBackreferenceHash
	a.DBHeight = tmp.Header.DBHeight
	a.BackReverenceHash = tmp.BackReverenceHash
	a.LookupHash = tmp.LookupHash

	// Use a regular expression to match the "adminidtype" field from the json
	// and unmarshal the ABEntry into its correct type
	for _, v := range tmp.ABEntries {
		switch {
		case regexp.MustCompile(`"adminidtype":0`).MatchString(string(v)):
			e := new(AdminMinuteNumber)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":1,`).MatchString(string(v)):
			e := new(AdminDBSignature)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":2,`).MatchString(string(v)):
			e := new(AdminRevealHash)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":3,`).MatchString(string(v)):
			e := new(AdminAddHash)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":4,`).MatchString(string(v)):
			e := new(AdminIncreaseServerCount)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":5,`).MatchString(string(v)):
			e := new(AdminAddFederatedServer)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":6,`).MatchString(string(v)):
			e := new(AdminAddAuditServer)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":7,`).MatchString(string(v)):
			e := new(AdminRemoveFederatedServer)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":8,`).MatchString(string(v)):
			e := new(AdminAddFederatedServerKey)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":9,`).MatchString(string(v)):
			e := new(AdminAddFederatedServerBTCKey)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":10,`).MatchString(string(v)):
			e := new(AdminServerFault)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":11,`).MatchString(string(v)):
			e := new(AdminCoinbaseDescriptor)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":12,`).MatchString(string(v)):
			e := new(AdminCoinbaseDescriptorCancel)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":13,`).MatchString(string(v)):
			e := new(AdminAddAuthorityAddress)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		case regexp.MustCompile(`"adminidtype":14,`).MatchString(string(v)):
			e := new(AdminAddAuthorityEfficiency)
			err := json.Unmarshal(v, e)
			if err != nil {
				return err
			}
			a.ABEntries = append(a.ABEntries, e)
		default:
			return ErrAIDUnknown
		}
	}

	return nil
}

type ABEntry interface {
	Type() AdminID
	String() string
}

type AdminMinuteNumber struct {
	MinuteNumber int `json:"minutenumber"`
}

func (a *AdminMinuteNumber) Type() AdminID {
	return AIDMinuteNumber
}

func (a *AdminMinuteNumber) String() string {
	return fmt.Sprintln("MinuteNumber:", a.MinuteNumber)
}

type AdminDBSignature struct {
	IdentityChainID   string `json:"identityadminchainid"`
	PreviousSignature struct {
		Pub string `json:"pub"`
		Sig string `json:"sig"`
	} `json:"prevdbsig"`
}

func (a *AdminDBSignature) Type() AdminID {
	return AIDDBSignature
}

func (a *AdminDBSignature) String() string {
	var s string

	s += fmt.Sprintln("DBSignature {")
	s += fmt.Sprintln("	IdentityChainID:", a.IdentityChainID)
	s += fmt.Sprintln("	PreviousSignature {")
	s += fmt.Sprintln("		Pub:", a.PreviousSignature.Pub)
	s += fmt.Sprintln("		Sig:", a.PreviousSignature.Sig)
	s += fmt.Sprintln("	}")
	s += fmt.Sprintln("}")

	return s
}

type AdminRevealHash struct {
	IdentityChainID string `json:"identitychainid"`
	MatryoshkaHash  string `json:"mhash"`
}

func (a *AdminRevealHash) Type() AdminID {
	return AIDRevealHash
}

func (a *AdminRevealHash) String() string {
	var s string

	s += fmt.Sprintln("RevealHash {")
	s += fmt.Sprintln("	IdentityChainID:", a.IdentityChainID)
	s += fmt.Sprintln("	MatryoshkaHash:", a.MatryoshkaHash)
	s += fmt.Sprintln("}")

	return s
}

type AdminAddHash struct {
	IdentityChainID string `json:"identitychainid"`
	MatryoshkaHash  string `json:"mhash"`
}

func (a *AdminAddHash) Type() AdminID {
	return AIDAddHash
}

func (a *AdminAddHash) String() string {
	var s string

	s += fmt.Sprintln("AddHash {")
	s += fmt.Sprintln("	IdentityChainID:", a.IdentityChainID)
	s += fmt.Sprintln("	MatryoshkaHash:", a.MatryoshkaHash)
	s += fmt.Sprintln("}")

	return s
}

type AdminIncreaseServerCount struct {
	Amount int `json:"amount"`
}

func (a *AdminIncreaseServerCount) Type() AdminID {
	return AIDIncreaseServerCount
}

func (a *AdminIncreaseServerCount) String() string {
	return fmt.Sprintln("IncreaseServerCount:", a.Amount)
}

type AdminAddFederatedServer struct {
	IdentityChainID string `json:"identitychainid"`
	DBHeight        int64  `json:"dbheight"`
}

func (a *AdminAddFederatedServer) Type() AdminID {
	return AIDAddFederatedServer
}

func (a *AdminAddFederatedServer) String() string {
	var s string

	s += fmt.Sprintln("AddFederatedServer {")
	s += fmt.Sprintln("	IdentityChainID:", a.IdentityChainID)
	s += fmt.Sprintln("	DBHeight:", a.DBHeight)
	s += fmt.Sprintln("}")

	return s
}

type AdminAddAuditServer struct {
	IdentityChainID string `json:"identitychainid"`
	DBHeight        int64  `json:"dbheight"`
}

func (a *AdminAddAuditServer) Type() AdminID {
	return AIDAddAuditServer
}

func (a *AdminAddAuditServer) String() string {
	var s string

	s += fmt.Sprintln("AdminAddAuditServer {")
	s += fmt.Sprintln("	IdentityChainID:", a.IdentityChainID)
	s += fmt.Sprintln("	DBHeight:", a.DBHeight)
	s += fmt.Sprintln("}")

	return s
}

type AdminRemoveFederatedServer struct {
	IdentityChainID string `json:"identitychainid"`
	DBHeight        int64  `json:"dbheight"`
}

func (a *AdminRemoveFederatedServer) Type() AdminID {
	return AIDRemoveFederatedServer
}

func (a *AdminRemoveFederatedServer) String() string {
	var s string

	s += fmt.Sprintln("RemoveFederatedServer {")
	s += fmt.Sprintln("	IdentityChainID:", a.IdentityChainID)
	s += fmt.Sprintln("	DBHeight:", a.DBHeight)
	s += fmt.Sprintln("}")

	return s
}

type AdminAddFederatedServerKey struct {
	IdentityChainID string `json:"identitychainid"`
	KeyPriority     int    `json:"keypriority"`
	PublicKey       string `json:"publickey"`
	DBHeight        int    `json:"dbheight"`
}

func (a *AdminAddFederatedServerKey) Type() AdminID {
	return AIDAddFederatedServerKey
}

func (a *AdminAddFederatedServerKey) String() string {
	var s string

	s += fmt.Sprintln("AddFederatedServerKey {")
	s += fmt.Sprintln("	IdentityChainID:", a.IdentityChainID)
	s += fmt.Sprintln("	KeyPriority:", a.KeyPriority)
	s += fmt.Sprintln("	PublicKey:", a.PublicKey)
	s += fmt.Sprintln("	DBHeight:", a.DBHeight)
	s += fmt.Sprintln("}")

	return s
}

type AdminAddFederatedServerBTCKey struct {
	IdentityChainID string `json:"identitychainid"`
	KeyPriority     int    `json:"keypriority"`
	KeyType         int    `json:"keytype"`
	ECDSAPublicKey  string `json:"ecdsapublickey"`
}

func (a *AdminAddFederatedServerBTCKey) Type() AdminID {
	return AIDAddFederatedServerBTCKey
}

func (a *AdminAddFederatedServerBTCKey) String() string {
	var s string

	s += fmt.Sprintln("AddFederatedServerBTCKey {")
	s += fmt.Sprintln("	IdentityChainID:", a.IdentityChainID)
	s += fmt.Sprintln("	KeyPriority:", a.KeyPriority)
	s += fmt.Sprintln("	KeyType:", a.KeyType)
	s += fmt.Sprintln("	ECDSAPublicKey:", a.ECDSAPublicKey)
	s += fmt.Sprintln("}")

	return s
}

type AdminServerFault struct {
	Timestamp     string `json:"timestamp"`
	ServerID      string `json:"serverid"`
	AuditServerID string `json:"auditserverid"`
	VMIndex       int    `json:"vmindex"`
	DBHeight      int    `json:"dbheight"`
	Height        int    `json:"height"`
	// TODO: change SignatureList type to match json return
	SignatureList json.RawMessage `json:"signaturelist"`
}

func (a *AdminServerFault) Type() AdminID {
	return AIDServerFault
}

func (a *AdminServerFault) String() string {
	var s string

	s += fmt.Sprintln("ServerFault {")
	s += fmt.Sprintln("	Timestamp:", a.Timestamp)
	s += fmt.Sprintln("	ServerID:", a.ServerID)
	s += fmt.Sprintln("	AuditServerID:", a.AuditServerID)
	s += fmt.Sprintln("	VMIndex:", a.VMIndex)
	s += fmt.Sprintln("	DBHeight:", a.DBHeight)
	s += fmt.Sprintln("	Height:", a.Height)
	s += fmt.Sprintln("	SignatureList:", a.SignatureList)
	s += fmt.Sprintln("}")

	return s
}

type AdminCoinbaseDescriptor struct {
	Outputs []struct {
		Amount  int    `json:"amount"`
		Address string `json:"address"`
	} `json:"outputs"`
}

func (a *AdminCoinbaseDescriptor) Type() AdminID {
	return AIDCoinbaseDescriptor
}

func (a *AdminCoinbaseDescriptor) String() string {
	var s string

	s += fmt.Sprintln("CoinbaseDescriptor {")
	for _, v := range a.Outputs {
		s += fmt.Sprintln("	Output {")
		s += fmt.Sprintln("		Amount:", v.Amount)
		s += fmt.Sprintln("		Address:", v.Address)
		s += fmt.Sprintln("	}")
	}
	s += fmt.Sprintln("}")

	return s
}

type AdminCoinbaseDescriptorCancel struct {
	DescriptorHeight int `json:"descriptor_height"`
	DescriptorIndex  int `json:descriptor_index`
}

func (a *AdminCoinbaseDescriptorCancel) Type() AdminID {
	return AIDCoinbaseDescriptorCancel
}

func (a *AdminCoinbaseDescriptorCancel) String() string {
	var s string

	s += fmt.Sprintln("CoinbaseDescriptorCancel {")
	s += fmt.Sprintln("	DescriptorHeight:", a.DescriptorHeight)
	s += fmt.Sprintln("	DescriptorIndex:", a.DescriptorIndex)
	s += fmt.Sprintln("}")

	return s
}

type AdminAddAuthorityAddress struct {
	IdentityChainID string `json:"identitychainid"`
	FactoidAddress  string `json:"factoidaddress"`
}

func (a *AdminAddAuthorityAddress) Type() AdminID {
	return AIDAddAuthorityAddress
}

func (a *AdminAddAuthorityAddress) String() string {
	var s string

	s += fmt.Sprintln("AddAuthorityAddress {")
	s += fmt.Sprintln("	IdentityChainID:", a.IdentityChainID)
	s += fmt.Sprintln("	FactoidAddress:", a.FactoidAddress)
	s += fmt.Sprintln("}")

	return s
}

type AdminAddAuthorityEfficiency struct {
	IdentityChainID string `json:"identitychainid"`
	Efficiency      int    `json:"efficiency"`
}

func (a *AdminAddAuthorityEfficiency) Type() AdminID {
	return AIDAddAuthorityEfficiency
}

func (a *AdminAddAuthorityEfficiency) String() string {
	var s string

	s += fmt.Sprintln("AddAuthorityEfficiency {")
	s += fmt.Sprintln("	IdentityChainID:", a.IdentityChainID)
	s += fmt.Sprintln("	Efficiency:", a.Efficiency)
	s += fmt.Sprintln("}")

	return s
}

// GetABlock requests a specified ABlock from the factomd API.
func GetABlock(keymr string) (ablock *ABlock, raw []byte, err error) {
	params := keyMRRequest{KeyMR: keymr}
	req := NewJSON2Request("admin-block", APICounter(), params)
	resp, err := factomdRequest(req)
	if err != nil {
		return
	}
	if resp.Error != nil {
		return nil, nil, resp.Error
	}

	// create a wraper construct for the ECBlock API return
	wrap := new(struct {
		ABlock  *ABlock `json:"ablock"`
		RawData string  `json:"rawdata"`
	})

	err = json.Unmarshal(resp.JSONResult(), wrap)
	if err != nil {
		return
	}

	raw, err = hex.DecodeString(wrap.RawData)
	if err != nil {
		return
	}

	return wrap.ABlock, raw, nil
}
