from rest_framework.response import Response
from rest_framework import status
from calls.models import Staff  # Ensure you import your Staff model
from calls.serializers import loginSerializeers  # Import your serializer

from calls import models,serializers
from rest_framework.decorators import api_view
from rest_framework.authtoken.models import Token


@api_view(['POST'])
def login(request):
    serializer = loginSerializeers(data=request.data)
    
    if serializer.is_valid():
        email = serializer.validated_data['email']
        password = request.data.get("password")
        
        try:
            user = Staff.objects.get(email=email)
          
            if user.password == password:  
                token, created = Token.objects.get_or_create(user=user)
                return Response({
                    'id': user.id,
                    'token': token.key,
                })
            else:
                return Response({"error": "Invalid credentials"}, status=status.HTTP_400_BAD_REQUEST)
        except Staff.DoesNotExist:
            return Response({"error": "User does not exist"}, status=status.HTTP_404_NOT_FOUND)
    
    return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)
