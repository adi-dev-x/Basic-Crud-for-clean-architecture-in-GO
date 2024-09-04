<!--
Hey, thanks for using the awesome-readme-template template.  
If you have any enhancements, then fork this project and create a pull request 
or just open an issue with the label "enhancement".

Don't forget to give this project a star for additional support ;)
Maybe you can mention me or this repo in the acknowledgements too
-->
<div align="center">

  <img src="assets/logo.png" alt="logo" width="200" height="auto" />
  <h1>CRUD-Clean Architecture in Go</h1>
  
  <p>
    A simple implementation of clean architecture using Go, the Echo framework, and GORM with a PostgreSQL database. Provides basic user management APIs for registration and login.
  </p>
  
  
<!-- Badges -->
<p>
  <a href="https://github.com/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO/graphs/contributors">
    <img src="https://img.shields.io/github/contributors/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO" alt="contributors" />
  </a>
  <a href="https://github.com/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO/commits/main">
    <img src="https://img.shields.io/github/last-commit/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO" alt="last update" />
  </a>
  <a href="https://github.com/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO/network/members">
    <img src="https://img.shields.io/github/forks/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO" alt="forks" />
  </a>
  <a href="https://github.com/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO/stargazers">
    <img src="https://img.shields.io/github/stars/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO" alt="stars" />
  </a>
  <a href="https://github.com/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO/issues/">
    <img src="https://img.shields.io/github/issues/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO" alt="open issues" />
  </a>
  <a href="https://github.com/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO.svg" alt="license" />
  </a>
</p>
   
<h4>
    <a href="https://github.com/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO">View Demo</a>
  <span> · </span>
    <a href="https://github.com/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO">Documentation</a>
  <span> · </span>
    <a href="https://github.com/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO/issues/">Report Bug</a>
  <span> · </span>
    <a href="https://github.com/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO/issues/">Request Feature</a>
  </h4>
</div>

<br />

<!-- Table of Contents -->
# :notebook_with_decorative_cover: Table of Contents

- [About the Project](#star2-about-the-project)
  * [Screenshots](#camera-screenshots)
  * [Tech Stack](#space_invader-tech-stack)
  * [Features](#dart-features)
  * [Environment Variables](#key-environment-variables)
- [Getting Started](#toolbox-getting-started)
  * [Prerequisites](#bangbang-prerequisites)
  * [Installation](#gear-installation)
  * [Run Locally](#running-run-locally)
  * [Deployment](#triangular_flag_on_post-deployment)
- [Usage](#eyes-usage)
- [Roadmap](#compass-roadmap)
- [Contributing](#wave-contributing)
  * [Code of Conduct](#scroll-code-of-conduct)
- [FAQ](#grey_question-faq)
- [License](#warning-license)
- [Contact](#handshake-contact)
- [Acknowledgements](#gem-acknowledgements)

<!-- About the Project -->
## :star2: About the Project

This project demonstrates a clean architecture implementation using Go, the Echo framework, and GORM with a PostgreSQL database. It provides two essential APIs for user management: registration and login.

### :camera: Screenshots

<div align="center"> 
  <img src="https://placehold.co/600x400?text=Your+Screenshot+here" alt="screenshot" />
</div>

### :space_invader: Tech Stack

<details>
  <summary>Server</summary>
  <ul>
    <li><a href="https://go.dev/">Golang</a></li>
    <li><a href="https://echo.labstack.com/">Echo</a></li>
    <li><a href="https://gorm.io/">GORM</a></li>
    <li><a href="https://www.postgresql.org/">PostgreSQL</a></li>
  </ul>
</details>

<!-- Features -->
### :dart: Features

- User registration API
- User login API

<!-- Environment Variables -->
### :key: Environment Variables

To run this project, you will need to add the following environment variables to your `.env` file:

- `DATABASE_URL`: URL for the PostgreSQL database connection.
- `JWT_SECRET`: Secret key for JWT token generation.

<!-- Getting Started -->
## :toolbox: Getting Started

### :bangbang: Prerequisites

- Go (1.18 or later)
- PostgreSQL
- GORM

### :gear: Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/adi-dev-x/Basic-Crud-for-clean-architecture-in-GO
