import { Accordion, Text, Flex } from "@mantine/core";
import { useEffect, useState } from "react";
import { GetNotebooks } from "../wailsjs/go/app/App";
import { db } from "../wailsjs/go/models";

export function Notebooks() {
  const [notebooks, setNotebooks] = useState<db.Notebook[]>([]);
  const [_loading, setLoading] = useState(false);
  const [_errors, setErrors] = useState(false);

  function transformNotebooks(notebooks: db.Notebook[]) {
    return notebooks.map((notebook) => (
      <Accordion.Item key={notebook.id} value={notebook.Title}>
        <Accordion.Control>
          <AccordionLabel {...notebook} />
        </Accordion.Control>
        <Accordion.Panel>List notes here</Accordion.Panel>
      </Accordion.Item>
    ));
  }

  function AccordionLabel({ Title }: db.Notebook) {
    return (
      <Flex component="span" gap="md" align="center" wrap="nowrap">
        <div>
          <Text span>{Title}</Text>
        </div>
      </Flex>
    );
  }

  useEffect(() => {
    async function load() {
      try {
        setLoading(true);
        const data = await GetNotebooks();
        console.log("data", data);
        setNotebooks(data);
      } catch (e) {
        setErrors(true);
      } finally {
        setLoading(false);
      }
    }

    void load();
  }, []);

  return (
    <Accordion chevronPosition="right" variant="contained">
      {transformNotebooks(notebooks)}
    </Accordion>
  );
}
