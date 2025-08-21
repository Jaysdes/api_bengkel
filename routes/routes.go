package routes

import (
	"api_bengkel/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Golang Bengkel Berjalan",
		})
	})
	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	api := router.Group("/api")
	{

		// Jenis Kendaraan
		api.GET("/jenis_kendaraan", controllers.GetJenisKendaraans)
		api.GET("/jenis_kendaraan/:id", controllers.GetJenisKendaraanByID)
		api.POST("/jenis_kendaraan", controllers.CreateJenisKendaraan)
		api.PUT("/jenis_kendaraan/:id", controllers.UpdateJenisKendaraan)
		api.DELETE("/jenis_kendaraan/:id", controllers.DeleteJenisKendaraan)

		// Customer routes
		api.GET("/customers", controllers.GetCustomers)
		api.GET("/customers/:id", controllers.GetCustomerByID)
		api.POST("/customers", controllers.CreateCustomer)
		api.PUT("/customers/:id", controllers.UpdateCustomer)
		api.DELETE("/customers/:id", controllers.DeleteCustomer)

		// Mekanik routes
		api.GET("/mekanik", controllers.GetMekaniks)
		api.GET("/mekanik/:id", controllers.GetMekanikByID)
		api.POST("/mekanik", controllers.CreateMekanik)
		api.PUT("/mekanik/:id", controllers.UpdateMekanik)
		api.DELETE("/mekanik/:id", controllers.DeleteMekanik)

		// Transaksi routes
		api.GET("/transaksi", controllers.GetTransaksis)
		api.GET("/transaksi/:id", controllers.GetTransaksiByID)
		api.POST("/transaksi", controllers.CreateTransaksi)
		api.PUT("/transaksi/:id", controllers.UpdateTransaksi)
		api.DELETE("/transaksi/:id", controllers.DeleteTransaksi)
		api.PUT("/transaksi/:id/bayar", controllers.BayarTransaksi)
		api.GET("/transaksi/:id/nota", controllers.CetakNota)

		//Jenis jasa
		api.GET("/jenis_jasa", controllers.GetJenisJasas)
		api.GET("/jenis_jasa/:id", controllers.GetJenisJasaByID)
		api.POST("/jenis_jasa", controllers.CreateJenisJasa)
		api.PUT("/jenis_jasa/:id", controllers.UpdateJenisJasa)
		api.DELETE("/jenis_jasa/:id", controllers.DeleteJenisJasa)

		//Jenis Service
		api.GET("/jenis_service", controllers.GetJenisServices)
		api.GET("/jenis_service/:id", controllers.GetJenisServiceByID)
		api.POST("/jenis_service", controllers.CreateJenisService)
		api.PUT("/jenis_service/:id", controllers.UpdateJenisService)
		api.DELETE("/jenis_service/:id", controllers.DeleteJenisService)

		//sparepart

		api.GET("/sparepart/:id", controllers.GetSparepartByID)
		api.GET("/sparepart", controllers.GetSpareparts)
		api.POST("/sparepart", controllers.CreateSparepart)
		api.PUT("/sparepart/:id", controllers.UpdateSparepart)
		api.DELETE("/sparepart/:id", controllers.DeleteSparepart)
		api.POST("/sparepart/kurangi_stok", controllers.KurangiStokSparepart)

		//spk
		api.GET("/spk", controllers.GetSPK)
		api.GET("/spk/:id", controllers.GetSPKByID)
		api.POST("/spk", controllers.CreateSPK)
		api.PUT("/spk/:id", controllers.UpdateSPK)
		api.DELETE("/spk/:id", controllers.DeleteSPK)

		//detail transaksi
		api.GET("/detail_transaksi", controllers.GetDetailTransaksi)
		api.GET("/detail_transaksi/:id", controllers.GetDetailTransaksiByID)
		api.POST("/detail_transaksi", controllers.CreateDetailTransaksi)
		api.PUT("/detail_transaksi/:id", controllers.UpdateDetailTransaksi)
		api.DELETE("/detail_transaksi/:id", controllers.DeleteDetailTransaksi)
		api.PUT("/detail_transaksi/:id/bayar", controllers.BayarDetailTransaksi)
		api.PUT("/detail_transaksi/:id/batal", controllers.BatalDetailTransaksi)

		// Crud Users
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUserByID)
		api.POST("/users", controllers.CreateUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)

		//Proses
		api.GET("/proses", controllers.GetAllProses)
		api.GET("/proses/:id", controllers.GetProsesByID)
		api.PUT("/proses/:id", controllers.UpdateProses)
		api.DELETE("/proses/:id", controllers.DeleteProses)

		//laporan
		api.GET("/laporan", controllers.GetAllLaporan)
		api.POST("/laporan", controllers.CreateLaporan)
		api.PUT("/laporan/:id", controllers.UpdateLaporan)
		api.DELETE("/laporan/:id", controllers.DeleteLaporan)

	}
}
