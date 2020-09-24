package dagpb

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/ipfs/go-cid"
	merkledag_pb "github.com/ipfs/go-merkledag/pb"
	ipld "github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"
)

// DecodeDagProto is a fast path decoding to protobuf
// from PBNode__NodeBuilders
func (nb _PBNode__NodeBuilder) DecodeDagProto(r io.Reader) (ipld.Node, error) {
	var pbn merkledag_pb.PBNode
	encoded, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("io error during unmarshal. %v", err)
	}
	if err := pbn.Unmarshal(encoded); err != nil {
		return nil, fmt.Errorf("unmarshal failed. %v", err)
	}
	pbLinks := make([]PBLink, 0, len(pbn.Links))
	for _, link := range pbn.Links {
		hash, err := cid.Cast(link.GetHash())

		if err != nil {
			return nil, fmt.Errorf("unmarshal failed. %v", err)
		}
		pbLinks = append(pbLinks, PBLink{
			&Link{cidlink.Link{Cid: hash}},
			&String{link.GetName()},
			&Int{int(link.GetTsize())},
		})
	}
	pbData := Bytes{pbn.GetData()}
	return PBNode{PBLinks{pbLinks}, pbData}, nil
}

// EncodeDagProto is a fast path encoding to protobuf
// for PBNode types
func (nd PBNode) EncodeDagProto(w io.Writer) error {
	pbn := new(merkledag_pb.PBNode)
	pbn.Links = make([]*merkledag_pb.PBLink, 0, len(nd.Links.x))
	for _, link := range nd.Links.x {
		var hash []byte
		if link.Hash != nil {
			cid := link.Hash.x.(cidlink.Link).Cid
			if cid.Defined() {
				hash = cid.Bytes()
			}
		}
		var name *string
		if link.Name != nil {
			name = &link.Name.x
		}
		var tsize *uint64
		if link.Tsize != nil {
			tmp := uint64(link.Tsize.x)
			tsize = &tmp
		}
		pbn.Links = append(pbn.Links, &merkledag_pb.PBLink{
			Hash:  hash,
			Name:  name,
			Tsize: tsize})
	}
	pbn.Data = nd.Data.x
	data, err := pbn.Marshal()
	if err != nil {
		return fmt.Errorf("marshal failed. %v", err)
	}
	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf(" error during marshal. %v", err)
	}
	return nil
}

// DecodeDagRaw is a fast path decoding to protobuf
// from RawNode__NodeBuilders
func (nb _RawNode__NodeBuilder) DecodeDagRaw(r io.Reader) (ipld.Node, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("io error during unmarshal. %v", err)
	}
	return RawNode{data}, nil
}

// EncodeDagRaw is a fast path encoding to protobuf
// for RawNode types
func (nd RawNode) EncodeDagRaw(w io.Writer) error {
	_, err := w.Write(nd.x)
	if err != nil {
		return fmt.Errorf(" error during marshal. %v", err)
	}
	return nil
}
