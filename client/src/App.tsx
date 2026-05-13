import { useState, KeyboardEvent } from "react";
import { Greet } from "../wailsjs/go/app/App";
import { Button, Stack, TextInput, Title, Text } from "@mantine/core";

export function App() {
  const [name, setName] = useState("");
  const [result, setResult] = useState("");
  const [loading, setLoading] = useState(false);

  async function handleGreet() {
    setLoading(true);

    try {
      const message = await Greet(name.trim() || "world");
      setResult(message);
    } finally {
      setLoading(false);
    }
  }

  function handleKeydown(e: KeyboardEvent<HTMLInputElement>) {
    if (e.key === "Enter") {
      void handleGreet();
    }
  }

  return (
    <Stack maw={480} mx="auto" mt="xl" px="md" gap="md">
      <Title order={1}>Nook</Title>
      <TextInput
        label="Your name"
        value={name}
        onKeyDown={handleKeydown}
        onChange={(e) => setName(e.currentTarget.value)}
      />
      <Button onClick={() => void handleGreet()} loading={loading}>
        Greet
      </Button>
      {result ? (
        <Text fw={600} size="lg">
          {result}
        </Text>
      ) : null}
    </Stack>
  );
}
