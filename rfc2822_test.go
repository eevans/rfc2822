
package rfc2822

import (
    "testing"
    "os"
)

const (
    DATA = "testdata.txt"
)

var testHeaders = map[string] string {
    "Header0": "Value0",
    "Header1": "Value1 Value1",
    "Header2": "Value2\n Value2\n Value2",
    "Header3": "Value3 Value3\n Value3\n Value3",
}

func TestParse(t *testing.T) {
    var (
        msg *Message
        err os.Error
    )

    if msg, err = ReadFile(DATA); err != nil {
       t.Error(err) 
    }

    for key, testValue := range testHeaders {
        if value, err := msg.GetHeader(key); err != nil {
            t.Error(err)
        } else {
            if value != testValue {
                t.Errorf("%s returned %s, expected %s", key, value, testValue)
            }
        }
    }
}

// vi: ai sw=4 ts=4 tw=0 et
