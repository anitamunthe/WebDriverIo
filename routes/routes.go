// api/routes/routes.go

package routes

import (
	"Rest-Api/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	kelas := r.Group("/kelas")
	{
		kelas.GET("/", handlers.GetKelasList)
		kelas.GET("/:id", handlers.GetKelasByID)
		kelas.POST("/", handlers.CreateKelas)
		kelas.PUT("/:id", handlers.UpdateKelas)
		kelas.DELETE("/:id", handlers.DeleteKelas)
	}

	r.POST("/login", handlers.Login)

	aspek := r.Group("/aspek")
	{
		aspek.GET("/", handlers.GetAspekList)
		aspek.GET("/:id", handlers.GetAspekByID)
		aspek.POST("/", handlers.CreateAspek)
		aspek.PUT("/:id", handlers.UpdateAspek)
		aspek.DELETE("/:id", handlers.DeleteAspek)
	}

	kegiatan := r.Group("/kegiatan")
	{
		kegiatan.GET("/", handlers.GetAllKegiatan)
		kegiatan.POST("/", handlers.CreateKegiatan)
		kegiatan.PUT("/:id", handlers.UpdateKegiatan)
	}

    poinAspek := r.Group("/poin-aspek")
    {
        poinAspek.GET("/", handlers.GetAllPoinAspek)
        poinAspek.GET("/:id", handlers.GetPoinAspekByID)
        poinAspek.POST("/", handlers.CreatePoinAspek)
        poinAspek.PUT("/:id", handlers.UpdatePoinAspek)
        poinAspek.DELETE("/:id", handlers.DeletePoinAspek)
    }


	return r
}
