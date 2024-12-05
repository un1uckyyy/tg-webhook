# tg-webhook

```bash
docker build -t tg-bot:dev ./bot

docker save tg-bot:dev > tg-bot.tar
```

with microk8s
```bash
microk8s enable ingress

microk8s ctr image import tg-bot.tar
```

you may need to create a secret with certificates for your domain because telegram api require https,
so we use nginx-ingress to ensure a secure connection and reverse-proxy webhook queries by plain http
to kubernetes service and after that to deployment with bot
```bash
microk8s kubectl create secret tls tls-secret \
  --cert=path/to/cert/file --key=path/to/key/file
```

```bash
microk8s kubectl apply -f ./k8s -R
```

also you need to register your webhook endpoint in telegram api by performing http request
```bash
curl https://api.telegram.org/bot<YOUR_TOKEN>/setWebhook?url=<YOUR_WEBHOOK_URL>
```
you should receive response like that:

```json
{"ok":true,"result":true,"description":"Webhook was set"}
```