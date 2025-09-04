# 构建阶段
FROM node:16-alpine as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install  # 安装依赖
COPY . .
RUN npm run build  # 构建静态文件

# 运行阶段（用Nginx部署）
FROM nginx:alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html  # 复制构建产物
COPY nginx.conf /etc/nginx/conf.d/default.conf  # 复制Nginx配置（如果有）
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
