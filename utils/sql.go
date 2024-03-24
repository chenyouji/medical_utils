package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type MySQL struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

func InitModel(sql *MySQL) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?loc=Local&parseTime=True&charset=utf8mb4", sql.User, sql.Password, sql.Host, sql.Port, sql.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&MedicalUser{}, &MedicalAddress{}, &MedicalConsult{}, &MedicalDoctor{}, &MedicalRegistration{}, &MedicalDrugs{}, &MedicalDoctorDepartment{}, &MedicalHospital{}, &MedicalIllness{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type MedicalUser struct {
	gorm.Model
	Username string `gorm:"varchar(20);not null;comment:用户名"`
	Password string `gorm:"varchar(50);not null;comment:密码"`
	Name     string `gorm:"varchar(15);not null;comment:姓名"`
	IdCard   string `gorm:"char(18);not null;comment:身份证号码"`
	Mobile   string `gorm:"char(11);not null;comment:手机号"`
	Age      int32  `gorm:"tinyint(3);not null;comment:年龄"`
	Sex      int32  `gorm:"tinyint(1);not null;comment:性别（1为男性，0为女性）"`
	Image    string `gorm:"varchar(200);comment:头像"`
}

func (MedicalUser) TableName() string {
	return "medical_user"
}

type MedicalAddress struct {
	gorm.Model
	UserID    int32  `gorm:"int;not null;comment:用户id"`
	Address   string `gorm:"varchar(50);comment:地址"`
	Name      string `gorm:"varchar(10);comment:收件人"`
	IsDefault int32  `gorm:"tinyint(1);not null;comment:是否默认选中1选中0不选中"`
}

func (MedicalAddress) TableName() string {
	return "medical_address"
}

type MedicalRegistration struct {
	gorm.Model
	UserID         int32   `gorm:"int;not null;comment:用户id"`
	DoctorID       int32   `gorm:"int;not null;comment:医生id"`
	Hospital       string  `gorm:"varchar(20);not null;comment:医院名称"`
	Amount         float32 `gorm:"decimal(10,2);not null;comment:挂号费"`
	PayType        int32   `gorm:"tinyint(1);not null;comment:支付方式"`
	RegistrationSn string  `gorm:"varchar(50);not null;comment:挂号号码"`
}

func (MedicalRegistration) TableName() string {
	return "medical_registration"
}

type MedicalDoctor struct {
	gorm.Model
	Name         string `gorm:"varchar(10);not null;comment:姓名"`
	DepartmentId int32  `gorm:"tinyint(3);not null;comment:科室id"`
	HospitalId   int32  `gorm:"tinyint(3);not null;comment:所属医院id"`
	Detail       string `gorm:"varchar(50);not null;comment:描述"`
	Title        string `gorm:"varchar(10);not null;comment:职称"`
	Mobile       string `gorm:"char(11);not null;comment:手机号"`
	Image        string `gorm:"varchar(200);comment:头像"`
	Adept        string `gorm:"varchar(100);not null;comment:擅长"`
}

func (MedicalDoctor) TableName() string {
	return "medical_doctor"
}

type MedicalDoctorDepartment struct {
	gorm.Model
	Name string `gorm:"varchar(10);not null;comment:科室"`
}

func (MedicalDoctorDepartment) TableName() string {
	return "medical_doctor_department"
}

type MedicalHospital struct {
	gorm.Model
	Name string `gorm:"varchar(12);not null;comment:医院名称"`
}

func (MedicalHospital) TableName() string {
	return "medical_hospital"
}

type MedicalDrugs struct {
	gorm.Model
	Name           string  `gorm:"varchar(20);not null;comment:药名"`
	Detail         string  `gorm:"varchar(150);not null;comment:药品描述"`
	DrugType       int32   `gorm:"tinyint(1);not null;comment:药品类型(1，西药，2，中药)"`
	IsPrescription int32   `gorm:"tinyint(1);not null;comment:是否是处方药(1为RX,2为OTC)"`
	InsDrugs       int32   `gorm:"tinyint(1);not null;comment:是否是医保药(1为是,0为否)"`
	Dosage         string  `gorm:"varchar(50);not null;comment:用药指导"`
	Taboo          string  `gorm:"varchar(50);not null;comment:饮食禁忌"`
	Price          float32 `gorm:"decimal(10,2);not null;comment:药品价格"`
	Image          string  `gorm:"varchar(200);comment:药品图片"`
}

func (MedicalDrugs) TableName() string {
	return "medical_drugs"
}

type MedicalConsult struct {
	gorm.Model
	UserID   int32  `gorm:"int;not null;comment:用户id"`
	DoctorID int32  `gorm:"int;not null;comment:医生id"`
	SendType int32  `gorm:"tinyint(1);not null;comment:消息发送类型(1,医生发送给用户，2,用户发送给医生)"`
	Message  string `gorm:"varchar(50);not null;comment:消息内容"`
}

func (MedicalConsult) TableName() string {
	return "medical_consult"
}

type MedicalIllness struct {
	gorm.Model
	Name               string `gorm:"varchar(50);not null;comment:病种名称"`
	SymptomDetail      string `gorm:"varchar(300);not null;comment:症状描述"`
	TreatmentRecommend string `gorm:"varchar(300);not null;comment:治疗建议"`
}

func (MedicalIllness) TableName() string {
	return "medical_illness"
}
