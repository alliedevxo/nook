import { Accordion, Text } from "@mantine/core";
import { useEffect, useState } from "react";
import { GetNotebooks } from "../wailsjs/go/app/App";
import { db } from "../wailsjs/go/models";

import "@mantine/core/styles.css";

export function Notebooks() {
  const [notebooks, setNotebooks] = useState<db.Notebook[]>([]);

  function transformNotebooks(notebooks: db.Notebook[]) {
    return notebooks.map((notebook) => (
      <Accordion.Item
        key={notebook.id}
        value={notebook.id + " " + notebook.Title}
      >
        <Accordion.Control>
          <AccordionLabel {...notebook} />
        </Accordion.Control>
        <Accordion.Panel>List notes here</Accordion.Panel>
      </Accordion.Item>
    ));
  }

  useEffect(() => {
    async function load() {
        const data = await GetNotebooks();
        setNotebooks(data);
    }

    void load();
  }, []);

  return (
    <Accordion chevronPosition="right">
      {transformNotebooks(notebooks)}
    </Accordion>
  );
}

function AccordionLabel({ Title }: db.Notebook) {
  return <Text span>{Title}</Text>;
}
