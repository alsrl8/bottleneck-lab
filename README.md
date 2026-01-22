# Bottleneck Lab

병목을 직접 만들고, 관찰하고, 해결하는 백엔드 인프라 실습 프로젝트.

## 목표

- DB connection pool exhaustion 경험 및 튜닝
- 메모리 누수 / OOM 유발 및 스트리밍 처리로 해결
- 네트워크 지연 시뮬레이션 및 circuit breaker 구현

## 아키텍처

![architecture](docs/architecture(2).png)

## 기술 스택

- Go, MySQL, Docker Compose
- k6 / hey (부하 테스트)
- Toxiproxy (네트워크 chaos)
- Grafana (선택)

## 상세 과제

👉 [Bottleneck Lab 과제 보기](https://YOUR_GITHUB_USERNAME.github.io/bottleneck-lab/)

> `docs/index.html`을 GitHub Pages로 호스팅하거나 로컬에서 열어서 사용