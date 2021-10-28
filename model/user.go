package model

import (
	"database/sql"

	"gorm.io/gorm"
)

//gorm.Model --> es una Struct embebida, lo unico que hace es agregar los campos ID, CreatedAt, UpdatedAt DeletedAt
//Es importante tener en cuenta que GORM va a considerar al campo ID como primary key
//En casos en los que querramos modificar ese comportamiento, podemos usar tags de Gorm para indicar que campo va ser la primary key
//Ejemplo si en vez de ID quisieramos UserID deberiamos agregar el tag `gorm:"primaryKey"` luego de type.
// UserID uint `gorm:"primaryKey"`

type User struct {
	gorm.Model
	Name  string
	Email string
	Phone sql.NullString
}

// equals
// type User struct {
// 	ID        uint `gorm:"primaryKey"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// 	Name      string
// 	Email      string
// }
// TableName retorna el nombre de la tabla usado por Gorm
func (User) TableName() string {
	return "public.users"
}
