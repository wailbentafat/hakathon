import configparser
[DEFAULT]
username = user123
password = pass123

[database]
host = localhost
port = 5432
database_name = mydb

[api]
url = http://localhost:2020
timeout = 30

# Create a ConfigParser object
config = configparser.ConfigParser()

# Read the configuration file
config.read('config.ini')

# Accessing values
username = config['DEFAULT']['username']
password = config['DEFAULT']['password']
db_host = config['database']['host']
db_port = config.getint('database', 'port')  # Use getint for integers
api_url = config['api']['url']
api_timeout = config.getint('api', 'timeout')

# Print the values
print(f"Username: {username}")
print(f"Password: {password}")
print(f"Database Host: {db_host}")
print(f"Database Port: {db_port}")
print(f"API URL: {api_url}")
print(f"API Timeout: {api_timeout}")
