
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

func TestManyHeaders(t *testing.T) {
    var (
        msg *Message
        err os.Error
        headers []Header
    )

    if msg, err = ReadFile(DATA); err != nil {
        t.Error(err)
    }

    if headers, err = msg.GetHeaders("header0"); err == nil {
        if headers[0].value != "Unexpected" {
            t.Errorf("expected \"Unexpected\", got \"%s\"", headers[0].value)
        }
        if headers[1].value != "Value0" {
            t.Errorf("expected \"Value0\", got \"%s\"", headers[0].value)
        }
    } else {
        t.Error(err)
    }
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
