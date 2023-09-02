# Advanced Transaction Application

## Introduction

This is an advanced bank-like transaction application, a robust and highly concurrent monolithic architecture application designed for production use. It empowers users to perform secure and efficient transactions.

## File Structure

The application consists of two primary layers:

- **Database Layer**: Manages various essential tables,
  - `db`: Contains advanced database table definitions and configurations, including those for accounts, users, transaction entries, transfer records, token sessions, email verification records, some advanced database techniques such as foreign keys, indexing are used.
- **Web API Layer**: Exposes endpoints, including HTTP and gRPC, along with several critical components:
  - `api`: Hosts HTTP APIs.
  - `gapi`: Provides gRPC APIs for specific use cases.
  - `mail`: Includes Gmail SMTP configuration.
  - `worker`: Handles the logic for sending emails, featuring Redis-based asynchronous processing.
  - `token`: Houses the PASETO token model.
  - `dockerfile`, `eks`: Configuration files for Docker and AWS EKS (Elastic Kubernetes Service).
  - `util`: Contains password hashing logic.
  - Unit tests are included for key modules.

## Tech Stack

The application leverages the following technologies:

- **Web Framework**: Utilizes Gin, gRPC, and gRPC Gateway for web services, supporting both RPC and HTTP requests.
- **Asynchronous Processing**: Redis is employed for asynchronous processing, enhancing scalability.
- **Version Control**: Git is used to track application milestones.
- **Containerization**: Docker, Docker Compose, Kubernetes are implemented to facilitate scaling.
- **Production Infrastructure**: AWS services, including AWS ECR, AWS EKS (Elastic Kubernetes Service), and AWS RDS (Relational Database Service), are employed for robust production capabilities. AWS IAM and AWS Secrets Manager are used for enhanced security and configuration management.
- **CI/CD**: GitHub Actions are implemented to establish a comprehensive CI/CD pipeline.

## Workflow

1. **User Registration**: Users sign up for accounts, and their information, including securely hashed passwords using bcrypt, is stored in PostgreSQL.
2. **User Authentication**: When users attempt to log in, a PASETO token and refresh token are sent to the user's registered email address for strict and secure authentication.
3. **Transaction Operations**: After confirming the token, users can perform essential CRUD operations, including adding funds, reviewing personal transaction histories, conducting transactions with others, and deleting their accounts.

## Key Highlights

One of the significant challenges in this application is optimizing performance, which is achieved through Test-Driven Development (TDD). To handle potential deadlocks during concurrent transactions (e.g., A to B and B to A simultaneously), a logic layer that compares account IDs is added to the database layer. This preprocessing logic ensures that transactions between accounts always occur from the smaller account ID to the larger account ID, mitigating deadlock scenarios.

## Getting Started

To launch the application:

1. Check the CI/CD workflow for deployment details.
2. Send HTTP requests using your preferred method, such as curl or Postman; RPC requests using Evans CLI or Postman.
3. Explore additional command-line options in the Makefile for advanced usage.
