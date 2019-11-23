package main

import (
	"encoding/json"
	"encoding/xml"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"math/rand"
	"time"
)

var _ = proto.Marshal

type NumPairs struct {
	Pair []*NumPair `protobuf:"bytes,1,rep,name=pair" json:"pair,omitempty"`
}

func (m *NumPairs) Reset()         { *m = NumPairs{} }
func (m *NumPairs) String() string { return proto.CompactTextString(m) }
func (*NumPairs) ProtoMessage()    {}
func (m *NumPairs) GetPair() []*NumPair {
	if m != nil {
		return m.Pair
	}
	return nil
}

type NumPair struct {
	Odd  int32 `protobuf:"varint,1,opt,name=odd" json:"odd,omitempty"`
	Even int32 `protobuf:"varint,2,opt,name=even" json:"even,omitempty"`
}

func (m *NumPair) Reset()         { *m = NumPair{} }
func (m *NumPair) String() string { return proto.CompactTextString(m) }
func (*NumPair) ProtoMessage()    {}
func init()                       {}

var numPairsStruct NumPairs
var numPairs = &numPairsStruct

func encodeAndserialize() {
	// XML encoding
	filename := "./pairs.xml"
	bytes, _ := xml.MarshalIndent(numPairs, "", " ")
	ioutil.WriteFile(filename, bytes, 0644)

	// JSON encoding
	filename = "./pairs.json"
	bytes, _ = json.MarshalIndent(numPairs, "", " ")
	ioutil.WriteFile(filename, bytes, 0644)

	// ProtoBuf encoding
	filename = "./pairs.pbuf"
	bytes, _ = proto.Marshal(numPairs)
	ioutil.WriteFile(filename, bytes, 0644)
}

const HowMany = 200 * 100 * 100 // two million

func main() {
	rand.Seed(time.Now().UnixNano())

	// comment out the modulus operations to get the original output
	for i := 0; i < HowMany; i++ {
		n1 := rand.Int31() % 2047
		if (n1 & 1) == 0 {
			n1++
		} // ensure it's odd
		n2 := rand.Int31() % 2047
		if (n2 & 1) == 1 {
			n2++
		} // ensure it's even

		next := &NumPair{
			Odd:  n1,
			Even: n2,
		}
		numPairs.Pair = append(numPairs.Pair, next)
	}
	encodeAndserialize()
}
