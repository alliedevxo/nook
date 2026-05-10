import { Greet } from "../wailsjs/go/app/App";

const nameInput = document.getElementById("name");
const goBtn = document.getElementById("go");
const result = document.getElementById("result");

goBtn.addEventListener("click", async () => {
  const message = await Greet(nameInput.value || "world");
  result.textContent = message;
});
