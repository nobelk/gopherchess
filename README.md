# Gopher chess
## Package
github.com/nobelk/gopherchess

# Generating mocks
## Notice
Package name should not contain special chars, otherwise mockgen fails to output to a file.
`mockgen -source=../engine/engine.go -destination=mockengine.go -package uci`
