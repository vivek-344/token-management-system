# Token Management System
I have used PostgreSQL to keep the data persistent. Whenever the app is ran, it checks for the old usage count and reset its count. If there are no tokens in the database, the system loads it automatically. when a user inputs, a token with least usage is chosen and increased its count by one. And display after performing all the operations.

---

#### To Set Up the Project

1. Clone the repository (of course)

2. add app.env with following variables

```

POSTGRES_USER=username

POSTGRES_PASSWORD=password

POSTGRES_DB=db_name

POSTGRES_HOST=localhost

DB_SOURCE=postgresql://username:password@postgres:5432/db_name?sslmode=disable

```

3. run the following command in app directory

```

docker compose run app

```

That's it!!! :)

---
#### Future Improvements
1. Optimization for speed
2. Progress Indication
3. Adding tests

---

#### Difficulties Faced
1. Unclear problem statement (it could have been more descriptive)
2. Using database/sql package for the first time (I used to use sqlc for my projects)