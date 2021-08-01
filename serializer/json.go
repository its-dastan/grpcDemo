package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)


func ProtobufToJSON(message proto.Message)([]byte, error) {
	marshaler:= protojson.MarshalOptions{
		Multiline: true,
		UseEnumNumbers: false,
		Indent: " ",
	}
	return marshaler.Marshal(message)
}