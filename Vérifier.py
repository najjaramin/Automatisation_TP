import requests

url = "https://example.com"
response = requests.get(url)
if response.status_code == 200:
    print("Site accessible")
else:
    print("Site inaccessible")
