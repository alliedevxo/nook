declare module "../wailsjs/go/app/App" {
  export function Greet(name: string): Promise<string>;
  export function GetNotebooks(): Promise<Notebook[]>;
}
