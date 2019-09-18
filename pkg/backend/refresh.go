package backend

import (
	"context"

	"github.com/hashicorp/vault/sdk/logical"
	"golang.org/x/oauth2"
)

func getToken(ctx context.Context, storage logical.Storage, key string) (*oauth2.Token, error) {
	entry, err := storage.Get(ctx, key)
	if err != nil {
		return nil, err
	} else if entry == nil {
		return nil, nil
	}

	tok := &oauth2.Token{}
	if err := entry.DecodeJSON(tok); err != nil {
		return nil, err
	}

	return tok, nil
}

func (b *backend) refreshToken(ctx context.Context, storage logical.Storage, key string) (*oauth2.Token, error) {
	b.credMut.Lock()
	defer b.credMut.Unlock()

	// In case someone else refreshed this token from under us, we'll re-request
	// it here with the lock acquired.
	tok, err := getToken(ctx, storage, key)
	if err != nil {
		return nil, err
	} else if tok == nil {
		return nil, nil
	} else if tok.Valid() || tok.RefreshToken == "" {
		return tok, nil
	}

	c, err := getConfig(ctx, storage)
	if err != nil {
		return nil, err
	} else if c == nil {
		return nil, ErrNotConfigured
	}

	p, err := c.provider(b.providerRegistry)
	if err != nil {
		return nil, err
	}

	// Refresh.
	src := p.NewExchangeConfigBuilder(c.ClientID, c.ClientSecret).
		Build().
		TokenSource(ctx, tok)

	tok, err = src.Token()
	if err != nil {
		return nil, err
	}

	// Store the new token.
	entry, err := logical.StorageEntryJSON(key, tok)
	if err != nil {
		return nil, err
	}

	if err := storage.Put(ctx, entry); err != nil {
		return nil, err
	}

	return tok, nil
}

func (b *backend) getRefreshToken(ctx context.Context, storage logical.Storage, key string) (*oauth2.Token, error) {
	tok, err := getToken(ctx, storage, key)
	if err != nil {
		return nil, err
	} else if tok == nil {
		return nil, nil
	}

	if !tok.Valid() {
		return b.refreshToken(ctx, storage, key)
	}

	return tok, nil
}

func (b *backend) refreshPeriodic(ctx context.Context, req *logical.Request) error {
	view := logical.NewStorageView(req.Storage, credsPathPrefix)
	logical.ScanView(ctx, view, func(path string) {
		if _, err := b.getRefreshToken(ctx, req.Storage, view.ExpandKey(path)); err != nil {
			b.logger.Error("unable to refresh token: %+v", err)
		}
	})

	return nil
}
