import { Button, Stack, Container } from "@chakra-ui/react";
import Navbar  from "./components/Navbar";

const App = () => {
  return (
    <Stack h="100vh">
      <Navbar />
      <Container>
        {/* <TodoForm />
        <TodoList /> */}
      </Container>
    </Stack>
  );
};

export default App;
