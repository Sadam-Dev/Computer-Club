package controllers

import (
	"ComputerClub/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userIDCtx           = "userID"
	userRoleCtx         = "userRole"
)

func checkUserAuthentication(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "empty auth header"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
		return
	}

	accessToken := headerParts[1]
	claims, err := service.ParseToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Claims extracted from token: %+v\n", claims)

	c.Set(userIDCtx, claims.UserID)
	c.Set(userRoleCtx, claims.RoleCode)

	fmt.Printf("User ID: %d, Role Code: %s\n", claims.UserID, claims.RoleCode)

	c.Next()
}

const (
	roleUser       = "user"
	roleAdmin      = "admin"
	roleSuperAdmin = "superAdmin"
)

// Middleware для проверки, что пользователь является администратором
func adminRequired(c *gin.Context) {
	userRole, exists := c.Get(userRoleCtx)
	fmt.Printf("User Role: %v\n", userRole)
	if !exists || (userRole != roleAdmin && userRole != roleSuperAdmin) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You don't have access to this resource"})
		return
	}
	c.Next()
}

// Middleware для проверки, что пользователь является суперадминистратором
func superAdminRequired(c *gin.Context) {
	userRole, exists := c.Get(userRoleCtx)
	if !exists || userRole != roleSuperAdmin {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You don't have access to this resource"})
		return
	}
	c.Next()
}

func adminOrSuperAdminRequired(c *gin.Context) {
	userRole, exists := c.Get(userRoleCtx)
	if !exists || (userRole != roleAdmin && userRole != roleSuperAdmin) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You don't have access to this resource"})
		return
	}
	c.Next()
}

func userRequired(c *gin.Context) {
	userRole, exists := c.Get(userRoleCtx)
	if !exists || userRole != roleUser {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You don't have access to this resource"})
		return
	}
	c.Next()
}
