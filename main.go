package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	router := gin.Default()
	router.POST("/", func(ctx *gin.Context) {
		var body any
		if err := ctx.ShouldBind(&body); err != nil {
			log.Fatal().Err(err).Msg("Couldn't parse body")
		}
		log.Info().Interface("body", body).Send()
	})
	router.Run()
}
