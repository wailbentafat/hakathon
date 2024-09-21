from django.db import models
 
 


class Staff(models.Model):
    password = models.CharField(max_length=128)
    email = models.EmailField(unique=True)
    is_admin = models.BooleanField(default=False)
    first_name = models.CharField(max_length=30)
    last_name = models.CharField(max_length=30)

    def __str__(self):
        return f"{self.first_name} {self.last_name}"


class Complaint(models.Model):
    bank_card = models.CharField(max_length=20, blank=True, null=True,default="")
    name = models.CharField(max_length=100, blank=True,null=True,default="")
    location = models.CharField(max_length=255, blank=True, null=True)
    phone_number = models.CharField(max_length=15, blank=True, null=True)
    description = models.TextField(blank=True, null=True)
    staff_id = models.ForeignKey(Staff, on_delete=models.SET_NULL, null=True, blank=True)
    created_at = models.DateTimeField(auto_now_add=True)
    email_address = models.EmailField(blank=True, null=True)
    bank_name = models.CharField(max_length=100, blank=True, null=True)
    website_url = models.URLField(blank=True, null=True)
    national_id_number = models.CharField(max_length=20, blank=True, null=True)
    card_type = models.CharField(max_length=50, blank=True, null=True)
    incident_date = models.DateField(blank=True, null=True)
    transaction_amount = models.DecimalField(max_digits=10, decimal_places=2, blank=True, null=True)
    transaction_date = models.DateField(blank=True, null=True)
    merchant_name = models.CharField(max_length=100, blank=True, null=True)
    merchant_registration = models.CharField(max_length=100, blank=True, null=True)

    def __str__(self):
        return f"Complaint by {self.name} "


class Category(models.Model):
    name = models.CharField(max_length=100)

    def __str__(self):
        return self.name
