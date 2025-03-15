package models

import (
	"time"

	"github.com/google/uuid"
)

// AccountEmailaddress represents the account_emailaddress table
type AccountEmailaddress struct {
	Email    string `gorm:"uniqueIndex;size:254"`
	Verified bool
	Primary  bool
	UserID   uint     `gorm:"column:user_id"`
	User     AuthUser `gorm:"foreignKey:UserID"`

	// Add gorm.Model if you want ID, CreatedAt, UpdatedAt, DeletedAt fields
}

func (AccountEmailaddress) TableName() string {
	return "account_emailaddress"
}

// AccountEmailconfirmation represents the account_emailconfirmation table
type AccountEmailconfirmation struct {
	ID             uint `gorm:"primaryKey"`
	Created        time.Time
	Sent           *time.Time
	Key            string              `gorm:"uniqueIndex;size:64"`
	EmailAddressID uint                `gorm:"column:email_address_id"`
	EmailAddress   AccountEmailaddress `gorm:"foreignKey:EmailAddressID"`
}

func (AccountEmailconfirmation) TableName() string {
	return "account_emailconfirmation"
}

// AuthGroup represents the auth_group table
type AuthGroup struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"uniqueIndex;size:150"`
}

func (AuthGroup) TableName() string {
	return "auth_group"
}

// AuthGroupPermissions represents the auth_group_permissions table
type AuthGroupPermissions struct {
	ID           uint           `gorm:"primaryKey"`
	GroupID      uint           `gorm:"column:group_id;uniqueIndex:idx_group_permission"`
	PermissionID uint           `gorm:"column:permission_id;uniqueIndex:idx_group_permission"`
	Group        AuthGroup      `gorm:"foreignKey:GroupID"`
	Permission   AuthPermission `gorm:"foreignKey:PermissionID"`
}

func (AuthGroupPermissions) TableName() string {
	return "auth_group_permissions"
}

// AuthPermission represents the auth_permission table
type AuthPermission struct {
	ID            uint              `gorm:"primaryKey"`
	Name          string            `gorm:"size:255"`
	ContentTypeID uint              `gorm:"column:content_type_id;uniqueIndex:idx_content_type_codename"`
	Codename      string            `gorm:"size:100;uniqueIndex:idx_content_type_codename"`
	ContentType   DjangoContentType `gorm:"foreignKey:ContentTypeID"`
}

func (AuthPermission) TableName() string {
	return "auth_permission"
}

// AuthUser represents the auth_user table
type AuthUser struct {
	ID          uint   `gorm:"primaryKey"`
	Password    string `gorm:"size:128"`
	LastLogin   *time.Time
	IsSuperuser bool
	Username    string `gorm:"uniqueIndex;size:150"`
	FirstName   string `gorm:"size:150"`
	LastName    string `gorm:"size:150"`
	Email       string `gorm:"size:254"`
	IsStaff     bool
	IsActive    bool
	DateJoined  time.Time
}

func (AuthUser) TableName() string {
	return "auth_user"
}

// AuthUserGroups represents the auth_user_groups table
type AuthUserGroups struct {
	ID      uint      `gorm:"primaryKey"`
	UserID  uint      `gorm:"column:user_id;uniqueIndex:idx_user_group"`
	GroupID uint      `gorm:"column:group_id;uniqueIndex:idx_user_group"`
	User    AuthUser  `gorm:"foreignKey:UserID"`
	Group   AuthGroup `gorm:"foreignKey:GroupID"`
}

func (AuthUserGroups) TableName() string {
	return "auth_user_groups"
}

// AuthUserUserPermissions represents the auth_user_user_permissions table
type AuthUserUserPermissions struct {
	ID           uint           `gorm:"primaryKey"`
	UserID       uint           `gorm:"column:user_id;uniqueIndex:idx_user_permission"`
	PermissionID uint           `gorm:"column:permission_id;uniqueIndex:idx_user_permission"`
	User         AuthUser       `gorm:"foreignKey:UserID"`
	Permission   AuthPermission `gorm:"foreignKey:PermissionID"`
}

func (AuthUserUserPermissions) TableName() string {
	return "auth_user_user_permissions"
}

// AuthtokenToken represents the authtoken_token table
type AuthtokenToken struct {
	Key     string `gorm:"primaryKey;size:40"`
	Created time.Time
	UserID  uint     `gorm:"column:user_id;uniqueIndex"`
	User    AuthUser `gorm:"foreignKey:UserID"`
}

func (AuthtokenToken) TableName() string {
	return "authtoken_token"
}

// DeliveryBalance represents the delivery_balance table
type DeliveryBalance struct {
	ID          uint    `gorm:"primaryKey"`
	Balance     float64 `gorm:"type:decimal(10,2)"`
	Currency    string  `gorm:"size:3"`
	LastUpdated time.Time
	IsActive    bool
	CompanyID   uint            `gorm:"column:company_id"`
	Company     DeliveryCompany `gorm:"foreignKey:CompanyID"`
}

func (DeliveryBalance) TableName() string {
	return "delivery_balance"
}

// DeliveryCar represents the delivery_car table
type DeliveryCar struct {
	ID            uint      `gorm:"primaryKey"`
	CarBrand      string    `gorm:"size:50"`
	NumberPlate   string    `gorm:"size:20"`
	SeatNumber    string    `gorm:"size:20"`
	PhotoDocument string    `gorm:"size:100"`
	UserID        *uint     `gorm:"column:user_id"`
	User          *AuthUser `gorm:"foreignKey:UserID"`
}

func (DeliveryCar) TableName() string {
	return "delivery_car"
}

// DeliveryClient represents the delivery_client table
type DeliveryClient struct {
	ID               uint             `gorm:"primaryKey"`
	PhoneNumber      string           `gorm:"size:100"`
	Address          string           `gorm:"type:text"`
	VerificationCode *string          `gorm:"size:6"`
	ChatID           *string          `gorm:"size:100"`
	PaymentMethod    *string          `gorm:"size:50"`
	UserID           uint             `gorm:"column:user_id;uniqueIndex"`
	User             AuthUser         `gorm:"foreignKey:UserID"`
	CompanyID        *uint            `gorm:"column:company_id"`
	Company          *DeliveryCompany `gorm:"foreignKey:CompanyID"`
}

func (DeliveryClient) TableName() string {
	return "delivery_client"
}

// DeliveryCompany represents the delivery_company table
type DeliveryCompany struct {
	ID           uint    `gorm:"primaryKey"`
	Name         string  `gorm:"size:255"`
	ContactEmail *string `gorm:"size:254"`
	PhoneNumber  *string `gorm:"size:20"`
	Address      *string `gorm:"type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Role         string   `gorm:"size:20"`
	Commission   *float64 `gorm:"type:decimal(5,2)"`
}

func (DeliveryCompany) TableName() string {
	return "delivery_company"
}

// DeliveryComment represents the delivery_comment table
type DeliveryComment struct {
	ID        uint   `gorm:"primaryKey"`
	Text      string `gorm:"type:text"`
	Timestamp time.Time
	ClientID  uint                  `gorm:"column:client_id"`
	CourierID uint                  `gorm:"column:courier_id"`
	TripID    uuid.UUID             `gorm:"column:trip_id;type:uuid"`
	Client    AuthUser              `gorm:"foreignKey:ClientID"`
	Courier   AuthUser              `gorm:"foreignKey:CourierID"`
	Trip      DeliveryOrderdelivery `gorm:"foreignKey:TripID"`
}

func (DeliveryComment) TableName() string {
	return "delivery_comment"
}

// DeliveryCompanyearnings represents the delivery_companyearnings table
type DeliveryCompanyearnings struct {
	ID              uint      `gorm:"primaryKey"`
	Amount          float64   `gorm:"type:decimal(10,2)"`
	Date            time.Time `gorm:"type:date"`
	Paid            bool
	CompanyID       uint                   `gorm:"column:company_id"`
	OrderDeliveryID *uuid.UUID             `gorm:"column:order_delivery_id;type:uuid"`
	Company         DeliveryCompany        `gorm:"foreignKey:CompanyID"`
	OrderDelivery   *DeliveryOrderdelivery `gorm:"foreignKey:OrderDeliveryID"`
}

func (DeliveryCompanyearnings) TableName() string {
	return "delivery_companyearnings"
}

// DeliveryContactsubmission represents the delivery_contactsubmission table
type DeliveryContactsubmission struct {
	ID        uint   `gorm:"primaryKey"`
	Country   string `gorm:"size:100"`
	City      string `gorm:"size:100"`
	ParkName  string `gorm:"size:200"`
	Phone     string `gorm:"size:20"`
	CreatedAt time.Time
}

func (DeliveryContactsubmission) TableName() string {
	return "delivery_contactsubmission"
}

// DeliveryCourier represents the delivery_courier table
type DeliveryCourier struct {
	ID                uint   `gorm:"primaryKey"`
	PhoneNumber       string `gorm:"size:100"`
	Avatar            string `gorm:"size:100"`
	PartnerType       string `gorm:"size:20"`
	Status            string `gorm:"size:20"`
	Rating            float32
	VerificationCode  *string `gorm:"size:6"`
	ChatID            *string `gorm:"size:100"`
	DocumentsProvided bool
	CarID             *uint             `gorm:"column:car_id"`
	PartnerID         *uint             `gorm:"column:partner_id"`
	UserID            uint              `gorm:"column:user_id;uniqueIndex"`
	CourierVariantID  *uint             `gorm:"column:courier_variant_id"`
	Services          string            `gorm:"size:50"`
	Car               *DeliveryCar      `gorm:"foreignKey:CarID"`
	Partner           *DeliveryCompany  `gorm:"foreignKey:PartnerID"`
	User              AuthUser          `gorm:"foreignKey:UserID"`
	CourierVariant    *DeliveryDelivery `gorm:"foreignKey:CourierVariantID"`
}

func (DeliveryCourier) TableName() string {
	return "delivery_courier"
}

// DeliveryDelivery represents the delivery_delivery table
type DeliveryDelivery struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:10"`
	Description string `gorm:"type:text"`
}

func (DeliveryDelivery) TableName() string {
	return "delivery_delivery"
}

// DeliveryOrderdelivery represents the delivery_orderdelivery table
type DeliveryOrderdelivery struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	Status         string    `gorm:"size:20"`
	Created        time.Time
	Updated        time.Time
	PickUpAddress  string  `gorm:"size:255"`
	DropOffAddress string  `gorm:"size:255"`
	Amount         float64 `gorm:"type:decimal(10,2)"`
	Paid           bool
	ClientID       *uint     `gorm:"column:client_id"`
	CourierID      *uint     `gorm:"column:courier_id"`
	Client         *AuthUser `gorm:"foreignKey:ClientID"`
	Courier        *AuthUser `gorm:"foreignKey:CourierID"`
}

func (DeliveryOrderdelivery) TableName() string {
	return "delivery_orderdelivery"
}

// DeliveryPriceplan represents the delivery_priceplan table
type DeliveryPriceplan struct {
	ID              uint             `gorm:"primaryKey"`
	Name            string           `gorm:"size:255"`
	BasePrice       float64          `gorm:"type:decimal(10,2)"`
	DistanceRate    float64          `gorm:"type:decimal(10,2)"`
	TimeRate        float64          `gorm:"type:decimal(10,2)"`
	SurgeMultiplier float64          `gorm:"type:decimal(5,2)"`
	PeakHoursStart  time.Time        `gorm:"type:time"`
	PeakHoursEnd    time.Time        `gorm:"type:time"`
	Mediacontent    string           `gorm:"size:100"`
	DeliveryID      uint             `gorm:"column:delivery_id"`
	Delivery        DeliveryDelivery `gorm:"foreignKey:DeliveryID"`
}

func (DeliveryPriceplan) TableName() string {
	return "delivery_priceplan"
}

// DeliveryVariation represents the delivery_variation table
type DeliveryVariation struct {
	ID                     uint              `gorm:"primaryKey"`
	Title                  string            `gorm:"size:100"`
	VariationName          string            `gorm:"size:100;uniqueIndex:idx_delivery_variation_name"`
	Description            string            `gorm:"type:text"`
	PriceValue             float64           `gorm:"type:decimal(10,2)"`
	Mediacontent           string            `gorm:"size:100"`
	DeliveryVariantPriceID uint              `gorm:"column:delivery_variant_price_id;uniqueIndex:idx_delivery_variation_name"`
	DeliveryVariantPrice   DeliveryPriceplan `gorm:"foreignKey:DeliveryVariantPriceID"`
}

func (DeliveryVariation) TableName() string {
	return "delivery_variation"
}

// DeliveryOrderdeliveryVariations represents the delivery_orderdelivery_variations table
type DeliveryOrderdeliveryVariations struct {
	ID              uint                  `gorm:"primaryKey"`
	OrderdeliveryID uuid.UUID             `gorm:"column:orderdelivery_id;type:uuid;uniqueIndex:idx_orderdelivery_variation"`
	VariationID     uint                  `gorm:"column:variation_id;uniqueIndex:idx_orderdelivery_variation"`
	Orderdelivery   DeliveryOrderdelivery `gorm:"foreignKey:OrderdeliveryID"`
	Variation       DeliveryVariation     `gorm:"foreignKey:VariationID"`
}

func (DeliveryOrderdeliveryVariations) TableName() string {
	return "delivery_orderdelivery_variations"
}

// DeliveryTrippriceplan represents the delivery_trippriceplan table
type DeliveryTrippriceplan struct {
	ID              uint             `gorm:"primaryKey"`
	Name            string           `gorm:"size:255"`
	BasePrice       float64          `gorm:"type:decimal(10,2)"`
	DistanceRate    float64          `gorm:"type:decimal(10,2)"`
	TimeRate        float64          `gorm:"type:decimal(10,2)"`
	SurgeMultiplier float64          `gorm:"type:decimal(5,2)"`
	PeakHoursStart  time.Time        `gorm:"type:time"`
	PeakHoursEnd    time.Time        `gorm:"type:time"`
	Mediacontent    string           `gorm:"size:100"`
	DeliveryID      uint             `gorm:"column:delivery_id"`
	Delivery        DeliveryDelivery `gorm:"foreignKey:DeliveryID"`
}

func (DeliveryTrippriceplan) TableName() string {
	return "delivery_trippriceplan"
}

// DeliveryTripvariation represents the delivery_tripvariation table
type DeliveryTripvariation struct {
	ID                 uint                  `gorm:"primaryKey"`
	Title              string                `gorm:"size:100"`
	VariationName      string                `gorm:"size:100;uniqueIndex:idx_trip_variation_name"`
	Description        string                `gorm:"type:text"`
	PriceValue         float64               `gorm:"type:decimal(10,2)"`
	Mediacontent       string                `gorm:"size:100"`
	TripVariantPriceID uint                  `gorm:"column:trip_variant_price_id;uniqueIndex:idx_trip_variation_name"`
	TripVariantPrice   DeliveryTrippriceplan `gorm:"foreignKey:TripVariantPriceID"`
}

func (DeliveryTripvariation) TableName() string {
	return "delivery_tripvariation"
}

// DjangoContentType represents the django_content_type table
type DjangoContentType struct {
	ID       uint   `gorm:"primaryKey"`
	AppLabel string `gorm:"size:100;uniqueIndex:idx_app_model"`
	Model    string `gorm:"size:100;uniqueIndex:idx_app_model"`
}

func (DjangoContentType) TableName() string {
	return "django_content_type"
}

// Additional models could be added here to complete the migration
