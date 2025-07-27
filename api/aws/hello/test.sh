aws lambda invoke \
  --function-name go-hello \
  --payload '{"name":"Allen3"}' \
  --log-type Tail \
  --region us-west-2 \
  --profile my-dev-profile \
  --cli-binary-format raw-in-base64-out \
  output.json

# ./test.sh