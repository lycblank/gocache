package gocache

import (
    "context"
)

type Cache interface {
    Get(ctx context.Context, entity Entity) error
    Delete(ctx context.Context, entity Entity) error
    Update(ctx context.Context, entity Entity) error
    Create(ctx context.Context, entity Entity) error
}

type Entity interface {
    GetPrimaryKey(ctx context.Context) string
    Decoder
    Encoder
}

type Decoder interface {
    Decode(ctx context.Context, datas []byte) error
}

type Encoder interface {
    Encoder(ctx context.Context) ([]byte, error)
}