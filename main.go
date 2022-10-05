package main

import (
	"fmt"
	"go-mod/src/config"
	"go-mod/src/entities"
	"go-mod/src/models"
)

func main() {
	Demo5()

	// 	c := swapi.DefaultClient

	// 	if atst, err := c.Planet(1); err == nil {

	// 		fmt.Println("Clima:", atst.Climate)
	// 		fmt.Println("Creado:", atst.Created)
	// 		fmt.Println("Diametro:", atst.Diameter)
	// 		fmt.Println("Editado:", atst.Edited)
	// 		fmt.Println("Pelicula:", atst.FilmURLs)

	// 		fmt.Println("Gravedad:", atst.Gravity)
	// 		fmt.Println("Nombre:", atst.Name)
	// 		fmt.Println("Periodo_orbital:", atst.OrbitalPeriod)
	// 		fmt.Println("Poblacion:", atst.Population)
	// 		fmt.Println("Residentes:", atst.ResidentURLs)

	// 		fmt.Println("Periodo_de_rotacion:", atst.RotationPeriod)
	// 		fmt.Println("Agua_superficial:", atst.SurfaceWater)
	// 		fmt.Println("Terreno:", atst.Terrain)
	// 		fmt.Println("url:", atst.URL)

	// 	}
}

func Demo1() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		planetModel := models.PlanetModel{
			Db: db,
		}
		planets, err2 := planetModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, planet := range planets {
				fmt.Println(planet.ToString())
				fmt.Println("---------------")
			}
		}
	}
}

//busqueda por id
func Demo2() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		planetModel := models.PlanetModel{
			Db: db,
		}
		planet, err2 := planetModel.Find(2)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(planet.ToString())
		}
	}
}

//Guardar datos
func Demo3() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		planetModel := models.PlanetModel{
			Db: db,
		}
		planet := entities.Planet{
			Clima:               "Cálido",
			Fecha_creacion:      "2014-12-09T13:50:49.641000Z",
			Diametro:            "10465",
			Fecha_edicion:       "2014-12-20T20:58:18.411000Z",
			Pelicula:            "https://swapi.dev/api/films/1/",
			Gravedad:            "1 standard",
			Nombre:              "Jhunior",
			Periodo_orbital:     "304",
			Poblacion:           "200000",
			Residentes:          "https://swapi.dev/api/people/1/",
			Periodo_de_rotacion: "23",
			Agua_superficial:    "1",
			Terreno:             "desert",
			Url:                 "https://swapi.dev/api/planets/1/",
		}
		err2 := planetModel.Create(&planet)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("Planet Info")
			fmt.Println(planet.ToString())
		}
	}
}

//Actualizar datos
func Demo4() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		planetModel := models.PlanetModel{
			Db: db,
		}
		planet := entities.Planet{
			Id: 2,

			Clima:               "Alex",
			Fecha_creacion:      "2014-12-09T13:50:49.641000Z",
			Diametro:            "10465",
			Fecha_edicion:       "2014-12-20T20:58:18.411000Z",
			Pelicula:            "https://swapi.dev/api/films/1/",
			Gravedad:            "1 standard",
			Nombre:              "Tatooine",
			Periodo_orbital:     "304",
			Poblacion:           "200000",
			Residentes:          "https://swapi.dev/api/people/1/",
			Periodo_de_rotacion: "23",
			Agua_superficial:    "1",
			Terreno:             "desert",
			Url:                 "https://swapi.dev/api/planets/1/",
		}
		rowsAffected, err2 := planetModel.Update(&planet)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("rowsAffected: ", rowsAffected)
		}
	}
}

//Eliminar datos
func Demo5() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		planetModel := models.PlanetModel{
			Db: db,
		}
		rowsAffected, err2 := planetModel.Delete(2)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("rowsAffected: ", rowsAffected)
		}
	}
}
