package main

import (
	"net/http/httputil"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	router := gin.Default()
	router.POST("/", func(ctx *gin.Context) {
		log.Info().
			Str("host", ctx.Request.Host).
			Str("remoteAddr", ctx.Request.RemoteAddr).
			Str("uri", ctx.Request.RequestURI)

		requestDump, err := httputil.DumpRequest(ctx.Request, true)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to dump request")
		}
		log.Info().Str("req", string(requestDump))

		var body any
		if err := ctx.ShouldBind(&body); err != nil {
			log.Fatal().Err(err).Msg("Couldn't parse body")
		}
		log.Info().Interface("body", body).Send()
	})
	router.Run()
}
