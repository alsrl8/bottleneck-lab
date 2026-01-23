# Bottleneck Lab

병목을 직접 만들고, 관찰하고, 해결하는 백엔드 인프라 실습 프로젝트.

## 목표

- DB connection pool exhaustion 경험 및 튜닝
- 메모리 누수 / OOM 유발 및 스트리밍 처리로 해결
- 네트워크 지연 시뮬레이션 및 circuit breaker 구현

## 과제 내용 및 결과 확인

👉 [Bottleneck Lab 과제 보기](https://alsrl8.github.io/bottleneck-lab/)

## 아키텍처

![architecture](docs/architecture(2).png)

## 기술 스택

- Go, MySQL, Docker Compose
- k6 / hey (부하 테스트)
- Toxiproxy (네트워크 chaos)
- Grafana (선택)
