import { Box, Flex, Paper } from "@mantine/core";
import { Notebooks } from "./Notebooks";

export function App() {
  return (
    <Flex h="100vh" gap="md">
      <Paper w={320} withBorder bdrs={0} style={{ overflowY: "auto" }}>
        <Notebooks />
      </Paper>
      <Box flex={1}>
        <Paper h="100%" p="md" style={{ overflowY: "auto" }}>
          Future content here
        </Paper>
      </Box>
    </Flex>
  );
}
