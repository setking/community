package middlewares

// 判断是否有admin权限
//func IsAdminAuth() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		claims, _ := c.Get("claims")
//		currentUser := claims.(*models.CustomClaims)
//		if currentUser.AuthorityId != 1 {
//			c.JSON(http.StatusForbidden, gin.H{
//				"msg": "没有权限",
//			})
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}
