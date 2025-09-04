# 构建阶段：使用与最新npm兼容的Node.js版本（20.x）
FROM node:20-alpine as build-stage
WORKDIR /app

# 复制依赖配置文件
COPY package.json package-lock.json ./

# 安装依赖（不强制升级npm，避免版本冲突）
RUN npm install

# 复制所有前端代码
COPY . .

# 构建生产环境代码
RUN npm run build

# 运行阶段：使用Nginx部署
FROM nginx:alpine as production-stage
# 复制构建产物到Nginx
COPY --from=build-stage /app/dist /usr/share/nginx/html
# 复制Nginx配置（如果项目中有自定义配置）
COPY nginx.conf /etc/nginx/conf.d/default.conf 2>/dev/null || true
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
