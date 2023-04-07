package models

type GoUsers struct {
	Id        int64  `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name" gorm:" type: varchar(255) COLLATE utf8mb4_unicode_ci not null"`
	LastName  string `json:"last_name" gorm:" type: varchar(255) COLLATE utf8mb4_unicode_ci not null"`
	Email     string `json:"email" gorm:" type: varchar(255) COLLATE utf8mb4_unicode_ci not null"`
	Password  string `json:"password" gorm:" type: varchar(255) COLLATE utf8mb4_unicode_ci not null"`
	Age       int    `json:"age"`
}
