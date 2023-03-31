# 2023 Solution Challenge - GSTEP

## 🌎 GSTEP is a service for 
In today's world, it's hard to feel the impact of environmental changes. We hear news about climate change, endangered species, and warnings from scientists, yet many people seem indifferent. Even those who do care often find it difficult to see the results of their actions, as the earth is so vast that the actions of a single person seem insignificant. This is the problem we want to solve: to provide feedback to those who take action towards creating a better world, since in the real world, feedback is often lacking.

## ️❗ Sustainable Development Goals
<img src="https://user-images.githubusercontent.com/113160789/229126005-9464e46a-ed95-4553-9b6c-56743e631c7b.png" width="500">

## 📹 Demo Video Link


## 🛠 Project Architecure
![Group 43 (1)](https://user-images.githubusercontent.com/113160789/229115333-7f399fd0-4959-4a94-90d2-a785212f3b74.png)

## 📱 Execution Method
The backend server repository for GSTEP is primarily intended for advanced users who wish to set up their own server environment. Although most users may not require or prefer to replicate this server, the following steps can guide those who want to do so.

1. [Download and install Go](https://go.dev/doc/install)
2. Configure PostgreSQL server
3. Set the environment variable with following command
```shell
$ export GSTEP_DSN="your_postgresql_server_DSN"
```
4. Run the server with following command in root directory
```shell
$ go run main.go
```

## 👥 Contributors
|[백승연](https://github.com/yeonannab)|[윤영재](https://github.com/yun0jae)|[이혜민](https://github.com/haaem)|[장율](https://github.com/yuljang)|
|---|---|---|---|
|<img src="https://user-images.githubusercontent.com/113160789/229120952-c360c9e1-f11f-425b-90ae-0049f13d2249.jpeg" width="200">|<img src="https://user-images.githubusercontent.com/113160789/229120233-f007bf1f-c202-418a-946d-761c13d3f686.jpeg" width="200">|<img src="https://user-images.githubusercontent.com/113160789/229121468-dffb5a12-8483-432c-9c9b-544369784f93.jpeg" width="200">|<img src="https://user-images.githubusercontent.com/113160789/229120613-3b83e377-3ec7-45c6-a542-c32566bfc1c1.jpeg" width="200">|