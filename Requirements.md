### Requirements Document for Web Forum Project

---

### Project Overview:

The objective of this project is to create a **web forum** that allows users to register, login, post content, comment on posts, like/dislike posts, and filter posts based on categories, created date, and liked status. The forum will be powered by a **SQLite database**, and the web application will need to handle user authentication, session management, and basic communication features between users.

---

### Functional Requirements:

1. **User Registration & Login:**
   - **Registration:**
     - **Email**: Must ask for the user's email.
     - **Username**: Must ask for the user's desired username.
     - **Password**: Must ask for the user's password.
     - The password should be stored **encrypted** (Bonus task).
     - When registering, check if the email is already taken. If yes, return an error response.
     - Successful registration should create a new user record in the database.

   - **Login:**
     - **Login Form**: Users should be able to input their email and password.
     - If the credentials are correct, the user will be logged in.
     - If the credentials are incorrect, return an error response.
     - **Session Management**: Use **cookies** for session management, allowing only one active session per user.
     - The cookie should have an expiration date. You can decide on how long the session will last.
     - Optionally, use **UUID** for generating session identifiers (Bonus task).

   - **Authentication**:
     - Users must be able to authenticate with their email and password (credentials should be verified against the database).
     - The system must return an error response if the password is incorrect.

2. **Post Creation & Categorization:**
   - Only **registered users** can create posts and comments.
   - Users should be able to associate one or more **categories** with a post (the categories are to be defined by the system or user).
   - Posts must be stored in the database with their associated categories.
   - Posts should be **visible to all users**, including **non-registered** users.

3. **Comments:**
   - Registered users can comment on posts.
   - **Non-registered** users can only view posts and comments but cannot add comments.

4. **Likes and Dislikes:**
   - **Likes/Dislikes**: Only registered users can like or dislike posts and comments.
   - The system should keep track of the number of **likes** and **dislikes** for each post and comment.
   - The number of likes and dislikes should be visible to **all users**.

5. **Filtering Posts:**
   - Registered users should be able to filter posts by:
     - **Categories** (subforums or topic-specific).
     - **Created Posts** (filter by creation date).
     - **Liked Posts** (only those posts that the user has liked).
   - Non-registered users can only view posts but will not be able to use the filtering options that require user authentication.

---

### Database Requirements (SQLite):

- You must use SQLite to store the following data:
  - **Users** (id, email, username, password hash, session information)
  - **Posts** (id, user_id, content, timestamp, categories)
  - **Comments** (id, user_id, post_id, content, timestamp)
  - **Likes/Dislikes** (id, user_id, post_id/comment_id, type: like/dislike)

- Use the following SQL queries at minimum:
  - **SELECT** (to retrieve data from tables)
  - **CREATE** (to create tables in the database)
  - **INSERT** (to insert new records)

---

### Non-Functional Requirements:

1. **Error Handling**:
   - The application should handle website errors properly and return appropriate **HTTP status codes**.
   - Technical errors, such as database failures or invalid inputs, must be handled gracefully.

2. **Code Quality**:
   - The code should adhere to good coding practices, be **modular**, and use **separation of concerns**.
   - You must ensure **unit testing** is implemented for the major components (like registration, login, post creation).

---

### Technologies:

1. **Backend**:
   - **Go** (Golang) will be used for backend development.
   - **SQLite3** for the database.
   - **bcrypt** (optional, for password encryption).
   - **UUID** (optional, for session management).

2. **Frontend**:
   - Use **plain HTML** (no frontend libraries like React, Angular, Vue, etc.).

3. **Docker**:
   - You must use Docker to containerize the application.
   - Create a **Docker image** and ensure that the application works inside a containerized environment.
   - Handle **compatibility** and **dependency management** effectively.

---

### Estimated Skills and Effort:

1. **Backend Development (Go)**:
   - Strong knowledge of **Go** and its standard library.
   - Experience working with **SQLite3** (basic queries, database design).
   - Understanding of **HTTP** and session management.
   - Ability to implement user authentication (login, registration).
   - Experience working with **cookies** for session management.
   - Optional: Knowledge of **bcrypt** for password encryption.
   - Optional: Familiarity with **UUID** generation for session management.

2. **Frontend Development**:
   - Basic HTML and **form handling** (registration, login).
   - Basic understanding of how to make a web page interactive without frontend frameworks.

3. **Docker**:
   - Understanding of **Docker containers**.
   - Ability to **containerize** the application and create Docker images.
   - Managing **dependencies** and ensuring the application runs in a Docker container.

4. **Database Design and SQL**:
   - Ability to design a proper **database schema** (e.g., users, posts, comments, likes/dislikes).
   - Knowledge of **SQL** queries (SELECT, INSERT, CREATE).
   - Ability to handle **data integrity** (e.g., unique email, password matching, etc.).

---

### Required Tasks:

1. **Backend Development:**
   - Set up the backend server using Go.
   - Implement user registration and login endpoints.
   - Implement post creation and comment functionality.
   - Implement like/dislike functionality.
   - Implement post filtering based on categories, created date, and likes.
   - Implement session management using cookies.

2. **Database Management:**
   - Design the SQLite database schema.
   - Implement the necessary queries (SELECT, INSERT, CREATE).
   - Implement a mechanism to handle likes/dislikes and categorize posts.

3. **Frontend Development:**
   - Create HTML forms for registration, login, post creation, and commenting.
   - Display posts, comments, and likes/dislikes.

4. **Testing:**
   - Write unit tests for key components (e.g., login, registration, session handling).

5. **Dockerization:**
   - Dockerize the application and create a working Docker image.
   - Ensure the application works in a containerized environment.

---

### Delivery:

1. **Source Code**: All source code files, including Go files, HTML files, and database schema.
2. **Unit Tests**: Test files for unit testing.
3. **Docker Image**: A functional Docker image for deployment.

This project will give you hands-on experience with backend development, authentication, session management, database design, and containerization using Docker. 

