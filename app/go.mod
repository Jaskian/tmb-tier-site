module github.com/jaskian/tmb-tier-site

go 1.20

require (
	github.com/jaskian/tmb-tier-site/tmb-json-parse v0.0.0-20230618151336-769456695be9
	github.com/jaskian/tmb-tier-site/web-builder v0.0.0-20230618151336-769456695be9
)

require github.com/jaskian/tmb-tier-site/shared v0.0.0-20230618151336-769456695be9 // indirect

replace github.com/jaskian/tmb-tier-site/tmb-json-parse => ../tmb-json-parse

replace github.com/jaskian/tmb-tier-site/shared => ../shared

replace github.com/jaskian/tmb-tier-site/web-builder => ../web-builder
