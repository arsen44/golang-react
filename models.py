from django.db import models


class AccountEmailaddress(models.Model):
    email = models.CharField(unique=True, max_length=254)
    verified = models.BooleanField()
    primary = models.BooleanField()
    user = models.ForeignKey('AuthUser', models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'account_emailaddress'
        unique_together = (('user', 'email'),)


class AccountEmailconfirmation(models.Model):
    created = models.DateTimeField()
    sent = models.DateTimeField(blank=True, null=True)
    key = models.CharField(unique=True, max_length=64)
    email_address = models.ForeignKey(AccountEmailaddress, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'account_emailconfirmation'


class AuthGroup(models.Model):
    name = models.CharField(unique=True, max_length=150)

    class Meta:
        managed = False
        db_table = 'auth_group'


class AuthGroupPermissions(models.Model):
    id = models.BigAutoField(primary_key=True)
    group = models.ForeignKey(AuthGroup, models.DO_NOTHING)
    permission = models.ForeignKey('AuthPermission', models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'auth_group_permissions'
        unique_together = (('group', 'permission'),)


class AuthPermission(models.Model):
    name = models.CharField(max_length=255)
    content_type = models.ForeignKey('DjangoContentType', models.DO_NOTHING)
    codename = models.CharField(max_length=100)

    class Meta:
        managed = False
        db_table = 'auth_permission'
        unique_together = (('content_type', 'codename'),)


class AuthUser(models.Model):
    password = models.CharField(max_length=128)
    last_login = models.DateTimeField(blank=True, null=True)
    is_superuser = models.BooleanField()
    username = models.CharField(unique=True, max_length=150)
    first_name = models.CharField(max_length=150)
    last_name = models.CharField(max_length=150)
    email = models.CharField(max_length=254)
    is_staff = models.BooleanField()
    is_active = models.BooleanField()
    date_joined = models.DateTimeField()

    class Meta:
        managed = False
        db_table = 'auth_user'


class AuthUserGroups(models.Model):
    id = models.BigAutoField(primary_key=True)
    user = models.ForeignKey(AuthUser, models.DO_NOTHING)
    group = models.ForeignKey(AuthGroup, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'auth_user_groups'
        unique_together = (('user', 'group'),)


class AuthUserUserPermissions(models.Model):
    id = models.BigAutoField(primary_key=True)
    user = models.ForeignKey(AuthUser, models.DO_NOTHING)
    permission = models.ForeignKey(AuthPermission, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'auth_user_user_permissions'
        unique_together = (('user', 'permission'),)


class AuthtokenToken(models.Model):
    key = models.CharField(primary_key=True, max_length=40)
    created = models.DateTimeField()
    user = models.OneToOneField(AuthUser, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'authtoken_token'


class DeliveryBalance(models.Model):
    id = models.BigAutoField(primary_key=True)
    balance = models.DecimalField(max_digits=10, decimal_places=2)
    currency = models.CharField(max_length=3)
    last_updated = models.DateTimeField()
    is_active = models.BooleanField()
    company = models.ForeignKey('DeliveryCompany', models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_balance'


class DeliveryCar(models.Model):
    id = models.BigAutoField(primary_key=True)
    car_brand = models.CharField(max_length=50)
    number_plate = models.CharField(max_length=20)
    seat_number = models.CharField(max_length=20)
    photo_document = models.CharField(max_length=100)
    user = models.ForeignKey(AuthUser, models.DO_NOTHING, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'delivery_car'


class DeliveryClient(models.Model):
    id = models.BigAutoField(primary_key=True)
    phone_number = models.CharField(max_length=100)
    address = models.TextField()
    verification_code = models.CharField(max_length=6, blank=True, null=True)
    chat_id = models.CharField(max_length=100, blank=True, null=True)
    payment_method = models.CharField(max_length=50, blank=True, null=True)
    user = models.OneToOneField(AuthUser, models.DO_NOTHING)
    company = models.ForeignKey('DeliveryCompany', models.DO_NOTHING, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'delivery_client'


class DeliveryComment(models.Model):
    id = models.BigAutoField(primary_key=True)
    text = models.TextField()
    timestamp = models.DateTimeField()
    client = models.ForeignKey(AuthUser, models.DO_NOTHING)
    courier = models.ForeignKey(AuthUser, models.DO_NOTHING, related_name='deliverycomment_courier_set')
    trip = models.ForeignKey('DeliveryOrderdelivery', models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_comment'


class DeliveryCompany(models.Model):
    id = models.BigAutoField(primary_key=True)
    name = models.CharField(max_length=255)
    contact_email = models.CharField(max_length=254, blank=True, null=True)
    phone_number = models.CharField(max_length=20, blank=True, null=True)
    address = models.TextField(blank=True, null=True)
    created_at = models.DateTimeField()
    updated_at = models.DateTimeField()
    role = models.CharField(max_length=20)
    commission = models.DecimalField(max_digits=5, decimal_places=2, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'delivery_company'


class DeliveryCompanyearnings(models.Model):
    id = models.BigAutoField(primary_key=True)
    amount = models.DecimalField(max_digits=10, decimal_places=2)
    date = models.DateField()
    paid = models.BooleanField()
    company = models.ForeignKey(DeliveryCompany, models.DO_NOTHING)
    order_delivery = models.ForeignKey('DeliveryOrderdelivery', models.DO_NOTHING, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'delivery_companyearnings'


class DeliveryContactsubmission(models.Model):
    id = models.BigAutoField(primary_key=True)
    country = models.CharField(max_length=100)
    city = models.CharField(max_length=100)
    park_name = models.CharField(max_length=200)
    phone = models.CharField(max_length=20)
    created_at = models.DateTimeField()

    class Meta:
        managed = False
        db_table = 'delivery_contactsubmission'


class DeliveryCourier(models.Model):
    id = models.BigAutoField(primary_key=True)
    phone_number = models.CharField(max_length=100)
    avatar = models.CharField(max_length=100)
    partner_type = models.CharField(max_length=20)
    status = models.CharField(max_length=20)
    rating = models.FloatField()
    verification_code = models.CharField(max_length=6, blank=True, null=True)
    chat_id = models.CharField(max_length=100, blank=True, null=True)
    documents_provided = models.BooleanField()
    car = models.ForeignKey(DeliveryCar, models.DO_NOTHING, blank=True, null=True)
    partner = models.ForeignKey(DeliveryCompany, models.DO_NOTHING, blank=True, null=True)
    user = models.OneToOneField(AuthUser, models.DO_NOTHING)
    courier_variant = models.ForeignKey('DeliveryDelivery', models.DO_NOTHING, blank=True, null=True)
    services = models.CharField(max_length=50)

    class Meta:
        managed = False
        db_table = 'delivery_courier'


class DeliveryCourierActiveDeliveryVariations(models.Model):
    id = models.BigAutoField(primary_key=True)
    courier = models.ForeignKey(DeliveryCourier, models.DO_NOTHING)
    variation = models.ForeignKey('DeliveryVariation', models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_courier_active_delivery_variations'
        unique_together = (('courier', 'variation'),)


class DeliveryCourierActiveTaxiTariffs(models.Model):
    id = models.BigAutoField(primary_key=True)
    courier = models.ForeignKey(DeliveryCourier, models.DO_NOTHING)
    trippriceplan = models.ForeignKey('DeliveryTrippriceplan', models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_courier_active_taxi_tariffs'
        unique_together = (('courier', 'trippriceplan'),)


class DeliveryCourierActiveTripVariations(models.Model):
    id = models.BigAutoField(primary_key=True)
    courier = models.ForeignKey(DeliveryCourier, models.DO_NOTHING)
    tripvariation = models.ForeignKey('DeliveryTripvariation', models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_courier_active_trip_variations'
        unique_together = (('courier', 'tripvariation'),)


class DeliveryCourierAllowedDeliveryVariations(models.Model):
    id = models.BigAutoField(primary_key=True)
    courier = models.ForeignKey(DeliveryCourier, models.DO_NOTHING)
    variation = models.ForeignKey('DeliveryVariation', models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_courier_allowed_delivery_variations'
        unique_together = (('courier', 'variation'),)


class DeliveryCourierAllowedTaxiTariffs(models.Model):
    id = models.BigAutoField(primary_key=True)
    courier = models.ForeignKey(DeliveryCourier, models.DO_NOTHING)
    trippriceplan = models.ForeignKey('DeliveryTrippriceplan', models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_courier_allowed_taxi_tariffs'
        unique_together = (('courier', 'trippriceplan'),)


class DeliveryCourierAllowedTripVariations(models.Model):
    id = models.BigAutoField(primary_key=True)
    courier = models.ForeignKey(DeliveryCourier, models.DO_NOTHING)
    tripvariation = models.ForeignKey('DeliveryTripvariation', models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_courier_allowed_trip_variations'
        unique_together = (('courier', 'tripvariation'),)


class DeliveryCourierearnings(models.Model):
    id = models.BigAutoField(primary_key=True)
    amount = models.DecimalField(max_digits=10, decimal_places=2)
    date = models.DateField()
    paid = models.BooleanField()
    courier = models.ForeignKey(DeliveryCourier, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_courierearnings'


class DeliveryDelivery(models.Model):
    id = models.BigAutoField(primary_key=True)
    name = models.CharField(max_length=10)
    description = models.TextField()

    class Meta:
        managed = False
        db_table = 'delivery_delivery'


class DeliveryDocumentcourier(models.Model):
    id = models.BigAutoField(primary_key=True)
    document_type = models.CharField(max_length=100)
    document_image = models.CharField(max_length=100)
    upload_date = models.DateTimeField()
    status = models.CharField(max_length=20)
    user = models.ForeignKey(AuthUser, models.DO_NOTHING, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'delivery_documentcourier'


class DeliveryOrderdelivery(models.Model):
    id = models.UUIDField(primary_key=True)
    status = models.CharField(max_length=20)
    created = models.DateTimeField()
    updated = models.DateTimeField()
    pick_up_address = models.CharField(max_length=255)
    drop_off_address = models.CharField(max_length=255)
    amount = models.DecimalField(max_digits=10, decimal_places=2)
    paid = models.BooleanField()
    client = models.ForeignKey(AuthUser, models.DO_NOTHING, blank=True, null=True)
    courier = models.ForeignKey(AuthUser, models.DO_NOTHING, related_name='deliveryorderdelivery_courier_set', blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'delivery_orderdelivery'


class DeliveryOrderdeliveryVariations(models.Model):
    id = models.BigAutoField(primary_key=True)
    orderdelivery = models.ForeignKey(DeliveryOrderdelivery, models.DO_NOTHING)
    variation = models.ForeignKey('DeliveryVariation', models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_orderdelivery_variations'
        unique_together = (('orderdelivery', 'variation'),)


class DeliveryPayment(models.Model):
    id = models.BigAutoField(primary_key=True)
    amount = models.DecimalField(max_digits=10, decimal_places=2)
    date = models.DateField()
    status = models.CharField(max_length=20)
    courier = models.ForeignKey(DeliveryCourier, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_payment'


class DeliveryPaymentcard(models.Model):
    id = models.BigAutoField(primary_key=True)
    payment_id = models.CharField(max_length=50)
    first6 = models.CharField(max_length=6)
    last4 = models.CharField(max_length=4)
    expiry_year = models.CharField(max_length=4)
    expiry_month = models.CharField(max_length=2)
    card_type = models.CharField(max_length=50)
    user = models.ForeignKey(AuthUser, models.DO_NOTHING, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'delivery_paymentcard'


class DeliveryPriceplan(models.Model):
    id = models.BigAutoField(primary_key=True)
    name = models.CharField(max_length=255)
    base_price = models.DecimalField(max_digits=10, decimal_places=2)
    distance_rate = models.DecimalField(max_digits=10, decimal_places=2)
    time_rate = models.DecimalField(max_digits=10, decimal_places=2)
    surge_multiplier = models.DecimalField(max_digits=5, decimal_places=2)
    peak_hours_start = models.TimeField()
    peak_hours_end = models.TimeField()
    mediacontent = models.CharField(max_length=100)
    delivery = models.ForeignKey(DeliveryDelivery, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_priceplan'


class DeliveryRating(models.Model):
    id = models.BigAutoField(primary_key=True)
    rating_value = models.IntegerField(blank=True, null=True)
    created_at = models.DateTimeField()
    courier = models.ForeignKey(AuthUser, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_rating'


class DeliveryServicecommission(models.Model):
    id = models.BigAutoField(primary_key=True)
    commission = models.DecimalField(max_digits=5, decimal_places=2, blank=True, null=True)
    commission_date = models.DateTimeField()

    class Meta:
        managed = False
        db_table = 'delivery_servicecommission'


class DeliveryTrippriceplan(models.Model):
    id = models.BigAutoField(primary_key=True)
    name = models.CharField(max_length=255)
    base_price = models.DecimalField(max_digits=10, decimal_places=2)
    distance_rate = models.DecimalField(max_digits=10, decimal_places=2)
    time_rate = models.DecimalField(max_digits=10, decimal_places=2)
    surge_multiplier = models.DecimalField(max_digits=5, decimal_places=2)
    peak_hours_start = models.TimeField()
    peak_hours_end = models.TimeField()
    mediacontent = models.CharField(max_length=100)
    delivery = models.ForeignKey(DeliveryDelivery, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_trippriceplan'


class DeliveryTripvariation(models.Model):
    id = models.BigAutoField(primary_key=True)
    title = models.CharField(max_length=100)
    variation_name = models.CharField(max_length=100)
    description = models.TextField()
    price_value = models.DecimalField(max_digits=10, decimal_places=2)
    mediacontent = models.CharField(max_length=100)
    trip_variant_price = models.ForeignKey(DeliveryTrippriceplan, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_tripvariation'
        unique_together = (('trip_variant_price', 'variation_name'),)


class DeliveryVariation(models.Model):
    id = models.BigAutoField(primary_key=True)
    title = models.CharField(max_length=100)
    variation_name = models.CharField(max_length=100)
    description = models.TextField()
    price_value = models.DecimalField(max_digits=10, decimal_places=2)
    mediacontent = models.CharField(max_length=100)
    delivery_variant_price = models.ForeignKey(DeliveryPriceplan, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'delivery_variation'
        unique_together = (('delivery_variant_price', 'variation_name'),)


class DjangoAdminLog(models.Model):
    action_time = models.DateTimeField()
    object_id = models.TextField(blank=True, null=True)
    object_repr = models.CharField(max_length=200)
    action_flag = models.SmallIntegerField()
    change_message = models.TextField()
    content_type = models.ForeignKey('DjangoContentType', models.DO_NOTHING, blank=True, null=True)
    user = models.ForeignKey(AuthUser, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'django_admin_log'


class DjangoContentType(models.Model):
    app_label = models.CharField(max_length=100)
    model = models.CharField(max_length=100)

    class Meta:
        managed = False
        db_table = 'django_content_type'
        unique_together = (('app_label', 'model'),)


class DjangoMigrations(models.Model):
    id = models.BigAutoField(primary_key=True)
    app = models.CharField(max_length=255)
    name = models.CharField(max_length=255)
    applied = models.DateTimeField()

    class Meta:
        managed = False
        db_table = 'django_migrations'


class DjangoSession(models.Model):
    session_key = models.CharField(primary_key=True, max_length=40)
    session_data = models.TextField()
    expire_date = models.DateTimeField()

    class Meta:
        managed = False
        db_table = 'django_session'


class DjangoSite(models.Model):
    domain = models.CharField(unique=True, max_length=100)
    name = models.CharField(max_length=50)

    class Meta:
        managed = False
        db_table = 'django_site'


class SocialaccountSocialaccount(models.Model):
    provider = models.CharField(max_length=200)
    uid = models.CharField(max_length=191)
    last_login = models.DateTimeField()
    date_joined = models.DateTimeField()
    extra_data = models.JSONField()
    user = models.ForeignKey(AuthUser, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'socialaccount_socialaccount'
        unique_together = (('provider', 'uid'),)


class SocialaccountSocialapp(models.Model):
    provider = models.CharField(max_length=30)
    name = models.CharField(max_length=40)
    client_id = models.CharField(max_length=191)
    secret = models.CharField(max_length=191)
    key = models.CharField(max_length=191)
    provider_id = models.CharField(max_length=200)
    settings = models.JSONField()

    class Meta:
        managed = False
        db_table = 'socialaccount_socialapp'


class SocialaccountSocialappSites(models.Model):
    id = models.BigAutoField(primary_key=True)
    socialapp = models.ForeignKey(SocialaccountSocialapp, models.DO_NOTHING)
    site = models.ForeignKey(DjangoSite, models.DO_NOTHING)

    class Meta:
        managed = False
        db_table = 'socialaccount_socialapp_sites'
        unique_together = (('socialapp', 'site'),)


class SocialaccountSocialtoken(models.Model):
    token = models.TextField()
    token_secret = models.TextField()
    expires_at = models.DateTimeField(blank=True, null=True)
    account = models.ForeignKey(SocialaccountSocialaccount, models.DO_NOTHING)
    app = models.ForeignKey(SocialaccountSocialapp, models.DO_NOTHING, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'socialaccount_socialtoken'
        unique_together = (('app', 'account'),)


class SpatialRefSys(models.Model):
    srid = models.IntegerField(primary_key=True)
    auth_name = models.CharField(max_length=256, blank=True, null=True)
    auth_srid = models.IntegerField(blank=True, null=True)
    srtext = models.CharField(max_length=2048, blank=True, null=True)
    proj4text = models.CharField(max_length=2048, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'spatial_ref_sys'
