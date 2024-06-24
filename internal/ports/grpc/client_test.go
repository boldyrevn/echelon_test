package grpc

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestParseVideoID(t *testing.T) {
    tests := []struct {
        url        string
        expectedID string
        ok         bool
    }{
        {
            url:        "https://www.youtube.com/watch?v=4wO7dI6_kk",
            expectedID: "4wO7dI6_kk",
            ok:         true,
        },
        {
            url:        "www.youtube.com/watch?v=4wO7dI6_kk",
            expectedID: "4wO7dI6_kk",
            ok:         true,
        },
        {
            url:        "whatahell",
            expectedID: "",
            ok:         false,
        },
        {
            url:        "www.youtube.com/watch",
            expectedID: "",
            ok:         false,
        },
    }

    for _, test := range tests {
        id, err := ParseVideoID(test.url)
        assert.Equal(t, test.expectedID, id)
        if test.ok {
            assert.NoError(t, err)
        } else {
            assert.Error(t, err)
        }
    }
}
