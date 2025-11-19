import requests

url = "https://example.com"
try:
    r = requests.get(url, timeout=5)
    if r.status_code == 200:
        print("Site accessible")
    else:
        print("Site inaccessible", r.status_code)
except:
    print("Erreur de connexion")
