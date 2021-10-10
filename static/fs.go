package static

import "embed"

//go:embed *

// FS defines a static filesystem that can be served by the server.
var FS embed.FS
