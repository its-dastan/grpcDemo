package serializer_test

import (
	"fmt"
	"testing"

	"github.com/its-dastan/grpcDemo/pb"
	"github.com/its-dastan/grpcDemo/sample"
	"github.com/its-dastan/grpcDemo/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.ReadProtobufFromJsonFile(laptop2, jsonFile)
	fmt.Println("laptop3")
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))
}
