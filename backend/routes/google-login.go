package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	gf "github.com/shareed2k/goth_fiber"
)

const (
	key = "523793815073-33qjc07d518jpoh5hcrv4ar12aa6atdc.apps.googleusercontent.com"
	sec = "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"
)

func Create(app *fiber.App) {

	app.Get("/api/oauth/auth/:provider/callback", func(ctx *fiber.Ctx) error {
		fmt.Println("google auth callback called")
		user, err := gf.CompleteUserAuth(ctx)
		if err != nil {
			return err
		}
		ctx.JSON(user)
		fmt.Println("google auth callback called")
		return nil
	})

	app.Get("/api/oauth/logout/:provider", func(ctx *fiber.Ctx) error {
		gf.Logout(ctx)
		ctx.Redirect("/")
		return nil
	})

	app.Get("/oauth/auth/:provider", func(ctx *fiber.Ctx) error {
		if gothUser, err := gf.CompleteUserAuth(ctx); err == nil {
			ctx.JSON(gothUser)
			fmt.Println("google auth called")
		} else {
			gf.BeginAuthHandler(ctx)
			fmt.Println("google auth called")
		}
		return nil
	})
	app.Get("/api/oauth", func(ctx *fiber.Ctx) error {
		ctx.Format("<p><a href='/api/oauth/auth/google'>google</a></p>")
		return nil
	})
	goth.UseProviders(
		google.New(key, sec, "http://localhost:8000/api/oauth/auth/google/callback"),
	)

}
