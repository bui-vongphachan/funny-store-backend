package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {

	router := gin.Default()

	CURRENT_PROXY := os.Getenv("CURRENT_PROXY")

	router.SetTrustedProxies([]string{CURRENT_PROXY})

	// create_user_routes.CreateUser(db, router)
	// company_routes.CreateCompany(db, router)
	// job_routes.CreateJob(db, router)
	// job_routes.CreateJobCategory(db, router)
	// job_routes.CreateCareer(db, router)
	// job_routes.GetJobDetail(db, router)
	// job_routes.GetJobDetailV2(db, router)
	// job_routes.UpdateJob(db, router)

	return router
}
