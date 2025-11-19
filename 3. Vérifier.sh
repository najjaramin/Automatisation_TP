URL="https://example.com"
STATUS=$(curl -o /dev/null -s -w "%{http_code}" $URL)
if [ "$STATUS" -eq 200 ]; then
  echo "Site accessible"
else
  echo "Site inaccessible $STATUS"
fi
