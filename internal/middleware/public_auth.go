package middleware

//* Public auth
// func (mw *MiddlewareManager) PublicAuth() fiber.Handler {
// 	return func(c *fiber.Ctx) error {

// 		requestTime := c.Get(cg.APIHeaderRequestTime)
// 		publicToken := c.Get(cg.APIHeaderApiKey)
// 		exc := exception.NewException(languages.ParseLanguage(c.Locals(cg.CtxLanguage)), mw.logger)

// 		newPublicToken := encryption.GenerateSHA256(fmt.Sprintf("%s:%v", mw.cfg.Authorization.Public.SecretKey, requestTime))
// 		if publicToken != newPublicToken {
// 			return exc.WriteErrorResponse(http.StatusUnauthorized, exc.Message.PermissionError("api"), errors.New(exc.Message.GeneralInvalidError("token").String()))
// 		}
// 		return c.Next()
// 	}
// }
