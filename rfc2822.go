// This package implements an RFC2822 parser.  It does not (yet )support
// the entire standard.
package rfc2822

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

const (
    CR   = "\r"
    LF   = "\n"
    CRLF = "\r\n"

    NOTFOUND = 1
)

// An Error can represent any printable error condition.
type Error struct {
    Type int
    args []string
}

func (err Error) String() string {
    switch err.Type {
    case NOTFOUND:
        return fmt.Sprintf("%s not found", strings.Join(err.args, " "))
    }
    return "Invalid Error Type"
}

// ParseError represents an error encountered parsing a message.
type ParseError struct {
    lineNo int
    reason string
}

func (err ParseError) String() string {
    return fmt.Sprintf("Error parsing input at line %d: %s",
        err.lineNo,
        err.reason)
}

type Header struct {
    key, value string
}

// A Message type encapsulates one RFC 2822 message.
type Message struct {
    headers map[string][]Header
    body    []string
}

// GetHeader retrieves an unstructured header value by its name, or an error
// if the requested header does not exist.  If more than one header of a 
// given name exists, the last one is returned.
func (msg *Message) GetHeader(header string) (val string, err os.Error) {
    if val, good := msg.headers[strings.ToLower(header)]; good {
        return val[len(val)-1].value, nil
    }
    return val, Error{NOTFOUND, []string{"header", fmt.Sprintf("'%s'", header)}}
}

// GetHeaders returns one or more unstructured headers for a name, or an error
// if the requested header does not exist.  When more than one header exists
// for a name, they are returned in the order they appear in the message.
func (msg *Message) GetHeaders(header string) (headers []Header, err os.Error) {
    var good bool
    if headers, good = msg.headers[strings.ToLower(header)]; good {
        return
    }
    err = Error{NOTFOUND, []string{"header", fmt.Sprintf("'%s'", header)}}
    return
}

// GetBody retrieves a message body if it exists, or an error if not.
func (msg *Message) GetBody() (val string, err os.Error) {
    if len(msg.body) < 1 {
        return val, Error{NOTFOUND, []string{"message body"}}
    }
    return strings.Join(msg.body, " "), nil
}

// ReadFile parses an RFC 2822 formatted input and returns a Message type.
func Read(reader io.Reader) (msg *Message, err os.Error) {
    buff := bufio.NewReader(reader)
    headers := make(map[string][]Header)

    var (
        key, val, lowerKey string
        lineNo             int
        inContent          bool = false
        body               []string
    )

    for {
        line, ioerr := buff.ReadString('\n')
        lineNo++

        if ioerr != nil {
            if ioerr != os.EOF {
                return nil, ioerr
            }
            if len(line) == 0 {
                break
            }
        }

        switch {
        case inContent:
            body = append(body, strings.TrimSpace(line))
        case strings.HasPrefix(line, LF) || strings.HasPrefix(line, CRLF):
            inContent = true
            continue
        case strings.HasPrefix(line, " "): // a field-body continuation?
            if len(key) == 0 {
                return nil, ParseError{lineNo, "No match for continuation"}
            }

            // Concatenate onto previous value
            val = fmt.Sprintf("%s\n %s", val, strings.TrimSpace(line))
            field := Header{key, val}
            headers[lowerKey][len(headers[lowerKey])-1] = field
        default:
            if i := strings.Index(line, ":"); i > 0 {
                key = strings.TrimSpace(line[0:i])
                lowerKey = strings.ToLower(key)
                val = strings.TrimSpace(line[i+1:])
                field := Header{key, val}

                if _, has := headers[lowerKey]; has {
                    headers[lowerKey] = append(headers[lowerKey], field)
                } else {
                    headers[lowerKey] = []Header{field}
                }
            } else {
                return nil, ParseError{lineNo, "Cannot parse field"}
            }
        }
    }

    msg = &Message{headers, body}
    return msg, nil
}

// ReadFile parses an RFC 2822 formatted file and returns a Message type.
func ReadFile(fname string) (msg *Message, err os.Error) {
    var file *os.File

    if file, err = os.Open(fname); err != nil {
        return msg, err
    }
    defer file.Close()

    return Read(file)
}

// vi: ai sw=4 ts=4 tw=0 et
