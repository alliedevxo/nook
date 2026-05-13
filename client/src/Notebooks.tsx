import { Accordion, Text } from "@mantine/core";
import { useEffect, useState } from "react";
import { GetNotebooks, GetNotes } from "../wailsjs/go/app/App";
import { db } from "../wailsjs/go/models";

import "@mantine/core/styles.css";

export function Notebooks() {
  const [notebooks, setNotebooks] = useState<db.Notebook[]>([]);
  const [notes, setNotes] = useState<db.Note[]>([]);
  const [opened, setOpened] = useState<string | null>(null);

  useEffect(() => {
    async function loadNotebooks() {
      const data = await GetNotebooks();
      setNotebooks(data);
    }

    void loadNotebooks();
  }, []);

  useEffect(() => {
    async function loadNotes() {
      const data = await GetNotes(Number(opened));
      setNotes(data);
    }

    void loadNotes();
  }, [opened]);

  return (
    <Accordion chevronPosition="right" onChange={setOpened} value={opened}>
      {notebooks.map((notebook) => (
        <Accordion.Item key={notebook.id} value={String(notebook.id)}>
          <Accordion.Control>
            <AccordionLabel {...notebook} />
          </Accordion.Control>
          <AccordionPanel notes={notes ?? []} />
        </Accordion.Item>
      ))}
    </Accordion>
  );
}

function AccordionLabel({ title }: db.Notebook) {
  return <Text span>{title}</Text>;
}

function AccordionPanel({ notes }: { notes: db.Note[] }) {
  return (
    <Accordion.Panel pl="md">
      {notes.map((note) => (
        <Text>{note.title}</Text>
      ))}
    </Accordion.Panel>
  );
}
