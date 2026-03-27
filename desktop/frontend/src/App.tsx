import { useState } from "react";
import logo from "./assets/images/logo-universal.png";

function App() {
  const [msg, setMsg] = useState("");

  return (
    <div id="App">
      <img src={logo} id="logo" alt="logo" />
      <div id="result" className="result">
        resultText
      </div>
    </div>
  );
}

export default App;
