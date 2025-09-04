# 构建阶段
FROM node:16-alpine as build-stage
WORKDIR /app
# 复制 package.json 和 package-lock.json（必须确保两者存在且匹配）
COPY package.json package-lock.json ./
# 升级 npm 到最新版，解决旧 lock 文件兼容问题
RUN npm install -g npm@latest
# 使用 npm ci 代替 npm install（更严格、更稳定，专为 CI/CD 设计）
RUN npm ci 
# 复制所有前端代码
COPY . .
# 构建生产环境代码
RUN npm run build  

# 运行阶段（Nginx 部署静态文件）
FROM nginx:alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf  # 若有 Nginx 配置需同步复制
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
