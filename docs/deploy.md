# Deploy

- 이 application은 학습 용도로 개발되었습니다. 
- 이 문서에서는 local deployment에 대해서만 다룹니다.

## Local Deployment

```shell
docker-compose up -d
```

- change api
```shell
docker-compose up -d --build api
```