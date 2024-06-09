package parser

import (
	"video-tools/internal/types"
)

type Parser interface {
	Parse() (*types.ParseResult, error)
}
