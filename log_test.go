package log

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestUnMarshalJSONFormat(t *testing.T) {
	file, err := os.Open("json_format.log")
	require.NoError(t, err)

	defer file.Close()
	data, err := UnMarshalJSONFormat(file)
	require.NoError(t, err)
	require.Equal(t, 66, len(data))
}
