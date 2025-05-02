FROM node:20-alpine
WORKDIR /app

# 環境変数ファイルを最初にコピーしてビルドキャッシュを効かせる
COPY frontend/.env.production .env

COPY frontend/package*.json ./
RUN npm ci --omit=dev

COPY frontend .
RUN npm run build

EXPOSE 3000
CMD ["npm","run","start"]