package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/bridge/core/internal/data/mem"
	"gitlab.com/tokend/bridge/core/internal/proxy"
	"gitlab.com/tokend/bridge/core/internal/service/handlers"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	proxyRepo, err := proxy.NewProxyRepo(s.cfg.Chains())
	if err != nil {
		panic(err)
	}

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxSigner(s.cfg.Signer()),
			handlers.CtxChainsQ(mem.NewChainsQ(s.cfg.Chains())),
			handlers.CtxTokensQ(mem.NewTokenQ(s.cfg.Tokens())),
			handlers.CtxTokenChainsQ(mem.NewTokenChainsQ(s.cfg.TokenChains())),
			handlers.CtxProxyRepo(proxyRepo),
		),
	)
	r.Route("/v1", func(r chi.Router) {
		r.Get("/chains", handlers.Chains)
		r.Route("/tokens", func(r chi.Router) {
			r.Get("/", handlers.Tokens)
			r.Get("/{token_id}/nfts/{nft_id}", handlers.Nft)
		})
		r.Route("/transfers", func(r chi.Router) {
			r.Post("/withdraw", handlers.Withdraw)
			r.Post("/deposit", handlers.Deposit)
			r.Post("/approve", handlers.Approve)
		})
	})

	return r
}

func strPtr(value string) *string {
	return &value
}
