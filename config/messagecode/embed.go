package messagecode

import _ "embed"

//go:embed catalog/message_codes.json
var messageCodesJSON []byte
