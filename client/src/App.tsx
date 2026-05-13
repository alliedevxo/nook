import { Stack } from "@mantine/core";
import { Notebooks } from "./Notebooks";

export function App() {
  return (
    <Stack maw={480} mx="auto" mt="xl" px="md" gap="md">
      <Notebooks />
    </Stack>
  );
}
