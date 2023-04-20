package src

import "embed"

//go:embed banks.json
//go:embed banknames.json
//go:embed IFSC.json
//go:embed custom-sublets.json
//go:embed sublet.json
var EmbeddedFileStorage embed.FS
