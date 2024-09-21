from django.urls import path, include
from rest_framework.routers import DefaultRouter
from .views import StaffViewSet, ComplaintViewSet

router = DefaultRouter()
router.register(r'staff', StaffViewSet)
router.register(r'complaints', ComplaintViewSet)

urlpatterns = [
    path('', include(router.urls)),
]
