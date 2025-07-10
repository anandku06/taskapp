import { Stack, Container } from "@chakra-ui/react";
import Navbar  from "./components/Navbar";
import TodoForm from "./components/TodoForm";
import TodoList from "./components/TodoList";

export const BASE_URL = "http://localhost:4000/api"

const App = () => {
  return (
    <Stack h="100vh">
      <Navbar />
      <Container>
        <TodoForm />
        <TodoList />
      </Container>
    </Stack>
  );
};

export default App;
