# calls/serializers.py
from rest_framework import serializers
from .models import Staff, Complaint, Category

class StaffSerializer(serializers.ModelSerializer):
    class Meta:
        model = Staff
        fields = '__all__' 

class ComplaintSerializer(serializers.ModelSerializer):
    class Meta:
        model = Complaint
        fields = '__all__'  

class CategorySerializer(serializers.ModelSerializer):
    class Meta:
        model = Category
        fields = '__all__'  
class loginSerializeers(serializers.ModelSerializer):
    class Meta:
        model=Staff
        fields=['email','password']