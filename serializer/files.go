package serializer

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON: %w", err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}

	return nil
}

func ReadProtobufFromJsonFile(message proto.Message, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}

	err = protojson.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshall: %w", err)
	}
	return nil
}


func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err!= nil{
		return fmt.Errorf("cannot marshal proto message to binary: %w",err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err!= nil{
		return fmt.Errorf("cannot write binary data to a file: %w",err)
	}
	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err!= nil{
		return fmt.Errorf("cannot read form binary file: %w",err)
	}
	err = proto.Unmarshal(data, message)

	fmt.Println(message)
	if err!= nil{
		return fmt.Errorf("cannot unmarshal proto message from binary: %w",err)
	}
	return nil
}