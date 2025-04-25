# TicketShop Movie Service

A microservice-based movie ticket booking system built with Go, MongoDB, Redis, and Kafka.

## Features

- Movie Management
  - View available movies
  - View coming soon movies
  - Movie showtimes
  - Seat availability

- Ticket Booking
  - Real-time seat reservation
  - Payment processing
  - Booking confirmation

## Tech Stack

- **Backend:** Go
- **Databases:** 
  - MongoDB (Primary database)
  - Redis (Caching)
- **Message Broker:** Apache Kafka
- **Containerization:** Docker
- **Container Orchestration:** Kubernetes

## Services Architecture

- **Movie Service:** Handles movie-related operations
- **Customer Service:** Manages customer information and authentication
- **Payment Service:** Processes payments and transactions with Omise
- **Booking Service:** Manages ticket reservations

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- MongoDB
- Redis
- Apache Kafka

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/yourusername/TicketShop-Movie.git
