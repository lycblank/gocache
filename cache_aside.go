package gocache

import "context"

type cacheAside struct {
    cache       Cache
    persistence Persistence
}

func NewCacheAside(cache Cache, persistence Persistence) (CacheService, error) {
    if cache == nil && persistence == nil {
        return nil, CacheAndPersistenceNull
    }
    ca := &cacheAside{
        cache:       cache,
        persistence: persistence,
    }
    return ca, nil
}

func (ca *cacheAside) Get(ctx context.Context, entity Entity) (err error) {
    if ca.cache != nil {
        if err = ca.cache.Get(ctx, entity); err == nil {
            return nil
        }
    }

    if ca.persistence != nil {
        err = ca.persistence.Get(ctx, entity)
        if err != nil {
            return err
        }
        // insert entity to cache
        if terr := ca.cache.Create(ctx, entity); terr != nil {
            // todo record log
        }
    }

    return err
}

func (ca *cacheAside) Delete(ctx context.Context, entity Entity) (err error) {
    if ca.cache != nil {
        if err = ca.cache.Delete(ctx, entity); err != nil {
            return err
        }
    }

    if ca.persistence != nil {
        if err = ca.persistence.Delete(ctx, entity); err != nil {
            return err
        }
    }

    return err
}
func (ca *cacheAside) Update(ctx context.Context, entity Entity) (err error) {
    if ca.persistence != nil {
        if err = ca.persistence.Update(ctx, entity); err != nil {
            return err
        }
        // delete cache after update persistence
        if ca.cache != nil {
            if terr := ca.cache.Delete(ctx, entity); terr != nil {
                // todo record error
            }
        }
    } else {
        // direct update cache wher persistence is null
        if ca.cache != nil {
            if err = ca.cache.Update(ctx, entity); err != nil {
                return err
            }
        }
    }
    return nil
}
func (ca *cacheAside) Create(ctx context.Context, entity Entity) (err error) {
    if ca.persistence != nil {
        if err = ca.persistence.Create(ctx, entity); err != nil {
            return err
        }
    } else {
        if err = ca.cache.Create(ctx, entity); err != nil {
            return err
        }
    }
    return nil
}
