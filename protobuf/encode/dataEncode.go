package main

import (
	"encoding/json" // JSON encoding
	"encoding/xml"  // XML encoding
	"fmt"
	"github.com/golang/protobuf/jsonpb" // protobuf/json conversions
	"github.com/golang/protobuf/proto"  // golang/protobuf interaction
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

//### Code from: protoc --go_out=. data.proto
var _ = proto.Marshal

type DataItem struct {
	OddA  int64   `protobuf:"varint,1,opt,name=oddA" json:"oddA,omitempty"`
	EvenA int64   `protobuf:"varint,2,opt,name=evenA" json:"evenA,omitempty"`
	OddB  int32   `protobuf:"varint,3,opt,name=oddB" json:"oddB,omitempty"`
	EvenB int32   `protobuf:"varint,4,opt,name=evenB" json:"evenB,omitempty"`
	Small float32 `protobuf:"fixed32,5,opt,name=small" json:"small,omitempty"`
	Big   float32 `protobuf:"fixed32,6,opt,name=big" json:"big,omitempty"`
	Short string  `protobuf:"bytes,7,opt,name=short" json:"short,omitempty"`
	Long  string  `protobuf:"bytes,8,opt,name=long" json:"long,omitempty"`
}

func (m *DataItem) Reset()         { *m = DataItem{} }
func (m *DataItem) String() string { return proto.CompactTextString(m) }
func (*DataItem) ProtoMessage()    {}
func init()                        {}

//### end of protoc code

const PbufFile = "./dataitem.pbuf"
const XmlFile = "./dataitem.xml"
const JsonFile = "./dataitem.json"

func encodeAndserialize(dataItem *DataItem) {
	// XML encoding (text)
	bytes, _ := xml.MarshalIndent(dataItem, "", " ")
	ioutil.WriteFile(XmlFile, bytes, 0644)

	// JSON encoding (text)
	bytes, _ = json.MarshalIndent(dataItem, "", " ")
	ioutil.WriteFile(JsonFile, bytes, 0644)

	// ProtoBuf encoding (binary)
	bytes, _ = proto.Marshal(dataItem)
	ioutil.WriteFile(PbufFile, bytes, 0644)
}

func test(orig *DataItem) {
	// test the jsonpb package for converting between Json/Protobuf
	marshal := jsonpb.Marshaler{}
	jsonStr, _ := marshal.MarshalToString(orig) // pbuf to json
	fmt.Println(jsonStr)                        // {"oddA":"8665166788083696923","evenA":"1781541684984057652",...}

	testItem := &DataItem{}
	_ = jsonpb.Unmarshal(strings.NewReader(jsonStr), testItem) // json to pbuf
	fmt.Println(testItem)                                      // oddA:5838657556748278843 evenA:6503071194550460952...

	// test the standard serializing by deserializing with the proto package
	filebytes, err := ioutil.ReadFile(PbufFile)
	if err != nil {
		fmt.Println("ReadFile")
	}

	testItem.Reset() // clear the structure
	err = proto.Unmarshal(filebytes, testItem)
	if err != nil {
		fmt.Println("Unmarshal")
	}

	fmt.Printf("\nOriginal:\n%d %d %d %d\n%f %f\n%s %s\n",
		orig.EvenA,
		orig.OddA,
		orig.EvenB,
		orig.OddB,
		orig.Big,
		orig.Small,
		orig.Long,
		orig.Short)

	fmt.Printf("\nDeserialized:\n%d %d %d %d\n%f %f\n%s %s\n",
		testItem.EvenA,
		testItem.OddA,
		testItem.EvenB,
		testItem.OddB,
		testItem.Big,
		testItem.Small,
		testItem.Long,
		testItem.Short)
}

const StrShort = 16
const StrLong = 32
const Chars = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ#$%^&"
const HowMany = 1
const UpperBound = 2047

func randString(n int) string {
	bytes := make([]byte, n)
	for i := range bytes {
		bytes[i] = Chars[rand.Intn(len(Chars))]
	}
	return string(bytes)
}

func main() {
	// create a ProtoBuf message with random data
	rand.Seed(time.Now().UnixNano())

	// variable-length integers
	n1 := rand.Int63()
	if (n1 & 1) == 0 {
		n1++
	} // ensure it's odd
	n2 := rand.Int63()
	if (n2 & 1) == 1 {
		n2++
	} // ensure it's even
	n3 := rand.Int31() % UpperBound
	if (n3 & 1) == 0 {
		n3++
	} // ensure it's odd
	n4 := rand.Int31() % UpperBound
	if (n4 & 1) == 1 {
		n4++
	} // ensure it's even

	// fixed-length floats
	var f1, f2 float32
	t1 := rand.Float32()
	t2 := rand.Float32()
	if t1 > t2 {
		f1 = t1
		f2 = t2
	} else {
		f2 = t1
		f1 = t2
	}

	// strings
	str1 := randString(StrShort)
	str2 := randString(StrLong)

	// the message
	dataItem := &DataItem{
		OddA:  n1,
		EvenA: n2,
		OddB:  n3,
		EvenB: n4,
		Big:   f1,
		Small: f2,
		Short: str1,
		Long:  str2,
	}

	// encode, serialize, and deserialize to test
	encodeAndserialize(dataItem)
	test(dataItem)
}
