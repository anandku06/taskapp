# 📝 Taskly

A modern, full-stack Todo application built with **Go (Fiber)** for the backend and **React (Vite + Chakra UI)** for the frontend.  
Easily manage your daily tasks with a beautiful, responsive interface and real-time updates!

---

## 🚀 Features

- ⚡️ **Fast API**: Go Fiber backend with MongoDB for persistent storage
- 🎨 **Modern UI**: React + Chakra UI for a sleek, responsive design
- 🌗 **Dark/Light Mode**: Seamless color mode switching
- 🔄 **Live Reload**: Instant backend reloads with [Air](https://github.com/cosmtrek/air)
- 🔍 **React Query**: Efficient data fetching and caching
- 📝 **CRUD Operations**: Create, read, update, and delete your todos
- 🛡️ **CORS Support**: Secure API access for your frontend

---

## 📦 Tech Stack

- **Backend:** Go, Fiber, MongoDB
- **Frontend:** React, Vite, Chakra UI, React Query
- **Dev Tools:** Air, dotenv, VS Code

---

## 🛠️ Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/anandku06/taskly.git
cd taskly
```

### 2. Setup Environment Variables

Create a `.env` file in the root directory:

```env
MONGODB_URI=your_mongodb_connection_string
ENV=development
PORT=4000
```

### 3. Run the Backend

Install dependencies and start the Go server with live reload:

```bash
go mod tidy
air
```

### 4. Run the Frontend

```bash
cd client
npm install
npm run dev
```

---

## 📂 Project Structure

```
taskly/
│
├── main.go                # Go Fiber backend
├── air.toml               # Air config for live reload
├── .env                   # Environment variables
├── client/                # React frontend
│   ├── src/
│   │   ├── App.tsx
│   │   ├── components/
│   │   │   ├── TodoList.tsx
│   │   │   └── TodoItem.tsx
│   │   └── chakra/
│   │       └── theme.ts
│   └── ...
└── ...
```

---

## 🤝 Contributing

Contributions are welcome!  
Feel free to open issues or submit pull requests.

---

## 📄 License

MIT License

---

## 🙏 Acknowledgements

- [Fiber](https://gofiber.io/)
- [Chakra UI](https://chakra-ui.com/)
- [React Query](https://tanstack.com/query/latest)
- [Air](https://github.com/cosmtrek/air)

---

Made with ❤️ by Anand