import { Accordion, Text } from "@mantine/core";
import { useEffect, useState } from "react";
import { GetNotebooks } from "../wailsjs/go/app/App";
import { db } from "../wailsjs/go/models";

import "@mantine/core/styles.css";

export function Notebooks() {
  const [notebooks, setNotebooks] = useState<db.Notebook[]>([]);

  useEffect(() => {
    async function load() {
      const data = await GetNotebooks();
      setNotebooks(data);
    }

    void load();
  }, []);

  return (
    <Accordion chevronPosition="right">
      {notebooks.map((notebook) => (
        <Accordion.Item
          key={notebook.id}
          value={notebook.id + " " + notebook.title}
        >
          <Accordion.Control>
            <AccordionLabel {...notebook} />
          </Accordion.Control>
          <Accordion.Panel>List notes here</Accordion.Panel>
        </Accordion.Item>
      ))}
    </Accordion>
  );
}

function AccordionLabel({ title }: db.Notebook) {
  return <Text span>{title}</Text>;
}
