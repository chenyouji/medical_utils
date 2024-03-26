package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Db *gorm.DB

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
	//err = db.AutoMigrate(&MedicalUser{}, &MedicalAddress{}, &MedicalConsult{}, &MedicalDoctor{}, &MedicalRegistration{}, &MedicalDrugs{}, &MedicalDoctorDepartment{}, &MedicalHospital{}, &MedicalIllness{})
	err = db.AutoMigrate(&MedicalEncyclopedia{}, &MedicalFollow{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

type MedicalUser struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;comment:用户名"`
	Password string `gorm:"type:varchar(50);not null;comment:密码"`
	Name     string `gorm:"type:varchar(15);not null;comment:姓名"`
	IdCard   string `gorm:"type:char(18);not null;comment:身份证号码"`
	Mobile   string `gorm:"type:char(11);not null;comment:手机号"`
	Age      int32  `gorm:"type:tinyint(3);not null;comment:年龄"`
	Sex      int32  `gorm:"type:tinyint(1);not null;comment:性别（1为男性，0为女性）"`
	Image    string `gorm:"type:varchar(200);comment:头像"`
}

func (MedicalUser) TableName() string {
	return "medical_user"
}

type MedicalAddress struct {
	gorm.Model
	UserID    int32  `gorm:"type:int;not null;comment:用户id"`
	Address   string `gorm:"type:varchar(50);comment:地址"`
	Name      string `gorm:"type:varchar(10);comment:收件人"`
	IsDefault int32  `gorm:"type:tinyint(1);not null;comment:是否默认选中1选中0不选中"`
}

func (MedicalAddress) TableName() string {
	return "medical_address"
}

type MedicalRegistration struct {
	gorm.Model
	UserID         int32   `gorm:"type:int;not null;comment:用户id"`
	DoctorID       int32   `gorm:"type:int;not null;comment:医生id"`
	Hospital       string  `gorm:"type:varchar(20);not null;comment:医院名称"`
	Amount         float32 `gorm:"type:decimal(10,2);not null;comment:挂号费"`
	PayType        int32   `gorm:"type:tinyint(1);not null;comment:支付方式"`
	RegistrationSn string  `gorm:"type:varchar(50);not null;comment:挂号号码"`
}

func (MedicalRegistration) TableName() string {
	return "medical_registration"
}

type MedicalDoctor struct {
	gorm.Model
	Name         string `gorm:"type:varchar(10);not null;comment:姓名"`
	DepartmentId int32  `gorm:"type:tinyint(3);not null;comment:科室id"`
	HospitalId   int32  `gorm:"type:tinyint(3);not null;comment:所属医院id"`
	Detail       string `gorm:"type:varchar(50);not null;comment:描述"`
	Title        string `gorm:"type:varchar(10);not null;comment:职称"`
	Mobile       string `gorm:"type:char(11);not null;comment:手机号"`
	Image        string `gorm:"type:varchar(200);comment:头像"`
	Adept        string `gorm:"type:varchar(100);not null;comment:擅长"`
}

func (MedicalDoctor) TableName() string {
	return "medical_doctor"
}

type MedicalDoctorDepartment struct {
	gorm.Model
	Name string `gorm:"type:varchar(10);not null;comment:科室"`
}

func (MedicalDoctorDepartment) TableName() string {
	return "medical_doctor_department"
}

type MedicalHospital struct {
	gorm.Model
	Name string `gorm:"type:varchar(12);not null;comment:医院名称"`
}

func (MedicalHospital) TableName() string {
	return "medical_hospital"
}

type MedicalDrugs struct {
	gorm.Model
	Name           string  `gorm:"type:varchar(20);not null;comment:药名"`
	Detail         string  `gorm:"type:varchar(150);not null;comment:药品描述"`
	DrugType       int32   `gorm:"type:tinyint(1);not null;comment:药品类型(1，西药，2，中药)"`
	IsPrescription int32   `gorm:"type:tinyint(1);not null;comment:是否是处方药(1为RX,2为OTC)"`
	InsDrugs       int32   `gorm:"type:tinyint(1);not null;comment:是否是医保药(1为是,2为否)"`
	Dosage         string  `gorm:"type:varchar(50);not null;comment:用药指导"`
	Taboo          string  `gorm:"type:varchar(50);not null;comment:饮食禁忌"`
	Price          float32 `gorm:"type:decimal(10,2);not null;comment:药品价格"`
	Image          string  `gorm:"type:varchar(200);comment:药品图片"`
}

func (MedicalDrugs) TableName() string {
	return "medical_drugs"
}

type MedicalConsult struct {
	gorm.Model
	UserID   int32  `gorm:"type:int;not null;comment:用户id"`
	DoctorID int32  `gorm:"type:int;not null;comment:医生id"`
	SendType int32  `gorm:"type:tinyint(1);not null;comment:消息发送类型(1,医生发送给用户，2,用户发送给医生)"`
	Message  string `gorm:"type:varchar(50);not null;comment:消息内容"`
}

func (MedicalConsult) TableName() string {
	return "medical_consult"
}

type MedicalIllness struct {
	gorm.Model
	Name               string `gorm:"type:varchar(50);not null;comment:病种名称"`
	SymptomDetail      string `gorm:"type:varchar(300);not null;comment:症状描述"`
	TreatmentRecommend string `gorm:"type:varchar(300);not null;comment:治疗建议"`
}

func (MedicalIllness) TableName() string {
	return "medical_illness"
}

type MedicalEncyclopedia struct {
	gorm.Model
	DoctorId     int32  `gorm:"type:int;not null;comment:医生id"`
	DepartmentId int32  `gorm:"type:tinyint(3);not null;comment:科室id"`
	CrowdId      int32  `gorm:"type:tinyint(1);comment:1儿童2男性3上班族4老年人5女性"`
	Name         string `gorm:"type:varchar(30);not null;comment:疾病名称"`
	Overview     string `gorm:"type:text;not null;comment:概述"`
	Symptom      string `gorm:"type:varchar(200);not null;comment:症状"`
	Etiology     string `gorm:"type:varchar(200);comment:病因"`
	FindDoctor   string `gorm:"type:varchar(200);comment:就医"`
	Treatment    string `gorm:"type:varchar(200);comment:治疗"`
	Daily        string `gorm:"type:varchar(200);comment:日常"`
	Prevent      string `gorm:"type:varchar(200);comment:预防"`
}

func (MedicalEncyclopedia) TableName() string {
	return "medical_encyclopedia"
}

type MedicalFollow struct {
	gorm.Model
	Uid      int32 `gorm:"type:int;not null;comment:用户id"`
	DoctorId int32 `gorm:"type:int;not null;comment:医生id"`
	Status   int32 `gorm:"type:tinyint(1);not null;comment:1已关注2未关注"`
}

func (MedicalFollow) TableName() string {
	return "medical_follow"
}

type MedicalCrowd struct {
	gorm.Model
	Name string `gorm:"type:varchar(5);comment:适应人群"`
}

func (MedicalCrowd) TableName() string {
	return "medical_crowd"
}
